package apiHandler

import (
	"encoding/json"
	"net/http"

	customizedNFTFacade "github.com/dileepaj/tracified-nft-backend/businessFacade/customizedNFTFacade"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonMethods"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**
 **Description:function is used to generate and savee a new OTP and send a email to the customer with the otp. The function will also retreive item data such
 **as itemID,batchID etc by using the provided productID parameter
 **Returns:Object ID of the new OTP created
 */
func InitNFT(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var requestGenOTP requestDtos.GenOTP
	decoder := json.NewDecoder(r.Body)

	tenantName, paramerr := r.URL.Query()["tenant"]
	if !paramerr {
		errors.BadRequest(W, "invalid Query param")
		return
	}
	tenantRst, getTenantErr := customizedNFTFacade.ValidateWalletTenant(tenantName[0])
	if getTenantErr != nil || tenantRst.Name == "" {
		errors.BadRequest(W, "Invalid tenant")
		return
	}

	err := decoder.Decode(&requestGenOTP)
	if err != nil {
		errors.BadRequest(W, "Invalid data")
		return
	}
	// Checks if the API has the ncessary params filled
	if requestGenOTP.ProductID != "" || requestGenOTP.Email != "" {
		if !validations.ValidateEmailFormat(requestGenOTP.Email) {
			errors.BadRequest(W, "Invalid Email address")
			return
		}
		logs.InfoLogger.Println(requestGenOTP.ProductID, requestGenOTP.Email)
		// Generate OTP
		otp, err := customizedNFTFacade.GenerateOTP(requestGenOTP.Email)
		// Triggered if there is an error generating the OTP
		if err != nil || otp == "" {
			errors.BadRequest(W, "Failed to generate OTP")
			// If the OTP is generated ProductID will be sent to get the batchID
		} else {
			// GET batch id when RURI product ID is given
			batchData, err := customizedNFTFacade.GetBatchIDDatabyItemID(requestGenOTP.ProductID)
			// If the API is unable to retreive the batch ID response with err msg will be sent
			if err != nil || otp == "" {
				errors.BadRequest(W, "Failed to retrive BatchID data")
				return
			} // If batch ID is retreived user email,otp and batch ID will be sent to be saved in the ruriOtp DB
			encodedOTP := commonMethods.StringToSHA256(otp)
			_, error := SaveUserOTPMapping(requestGenOTP.Email, encodedOTP, batchData.BatchID, requestGenOTP.ProductID)
			if error != nil {
				errors.BadRequest(W, error.Error())
				return
			} else {
				err1 := customizedNFTFacade.SendEmail(otp, requestGenOTP.Email)
				if err1 != nil {
					errors.BadRequest(W, "Incorrect email address")
				} else {
					commonResponse.SuccessStatus[string](W, "Email Has been Sent to "+requestGenOTP.Email)
				}
			}
		}
	} else { //* If necssary params are not there Error Message will be sent as a response
		errors.BadRequest(W, "Product ID or email is missing")
	}
}

func SaveUserOTPMapping(email string, otp string, batchID string, shopId string) (string, error) {
	var userAuth models.UserAuth
	userAuth.Email = email
	userAuth.Otp = otp
	userAuth.BatchID = batchID // Default number of attempts
	userAuth.ShopID = shopId
	userAuth.Validated = "False"
	userAuth.ExpireDate = primitive.NewDateTimeFromTime(customizedNFTFacade.GenerateOTPExpireDate())
	logs.InfoLogger.Println("NEW exp date created: ", userAuth.ExpireDate)
	result, error := customizedNFTFacade.SaveOTP(userAuth)
	if error != nil {
		return result, error
	} else {
		return result, nil
	}
}

/**
 **Description:This function is used to validate a OTP provided by the user. The email and the otp will be sent where the api.Checks if the otp and the email recived
 **have a matching record in the DB
 **Returns:If There is a matching email and a OTP in the DB the generated SVG will be reutrned as a response. If not error msg is sent as response
 */
func ValidateOTP(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var requestValidateOTP requestDtos.ValidateOTP
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestValidateOTP)
	if err != nil {
		errors.BadRequest(W, "Invalid data")
		return
	}
	// Checks if the API has the ncessary params filled
	if requestValidateOTP.OTPCode != "" || requestValidateOTP.Email != "" {
		encodedOTP := commonMethods.StringToSHA256(requestValidateOTP.OTPCode)
		if !validations.ValidateEmailFormat(requestValidateOTP.Email) {
			errors.BadRequest(W, "Invalid email address")
			return
		}
		status, errs := customizedNFTFacade.GetNFTStatus(requestValidateOTP.Email, encodedOTP)
		if errs != nil {
			errors.BadRequest(W, "Failed to get NFT status")
			return
		} else {
			if status == "Minted" {
				errors.BadRequest(W, "NFT already minted")
				return
			}
			rst, shopID, err := customizedNFTFacade.ValidateOTP(requestValidateOTP.Email, encodedOTP)
			if err != nil {
				errors.BadRequest(W, "Failed to validate OTP")
				return
			}
			if rst == "Invalid OTP" {
				errors.BadRequest(W, rst)
				return
			}
			rstshopid, shopIderr := customizedNFTFacade.GetNFTStatusbyShopID(shopID)
			if shopIderr != nil || rstshopid == "Minted" {
				errors.BadRequest(W, "NFT already claimed")
				return
			}
			tempBatchID := rst
			rst1, err1 := customizedNFTFacade.GetSVGbyEmailandBatchID(requestValidateOTP.Email, tempBatchID)
			if err != nil {
				errors.BadRequest(W, err1.Error())
				return
			}
			Response := models.Response{
				ShopID: shopID,
				SVGID:  rst1.SvgID,
				Status: "Valid",
			}
			commonResponse.SuccessStatus[models.Response](W, Response)
		}
	} else {
		errors.BadRequest(W, "Email or OTP missing")
		return
	}
}

/**
 **Description : Is used to generate a new OTP code and send to the provided email address
 **Returns : If the OTP is succesfully generated the batch ID is returned.
 */
func ResentOTP(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var requestResendOTP requestDtos.GenOTP
	decoder := json.NewDecoder(r.Body)

	tenantName, paramerr := r.URL.Query()["tenant"]
	if !paramerr {
		errors.BadRequest(W, "invalid Query param")
		return
	}
	tenantRst, getTenantErr := customizedNFTFacade.ValidateWalletTenant(tenantName[0])
	if getTenantErr != nil || tenantRst.Name == "" {
		errors.BadRequest(W, "Invalid tenant")
		return
	}

	err := decoder.Decode(&requestResendOTP)
	if err != nil {
		errors.BadRequest(W, "Invalid data")
		return
	}
	if requestResendOTP.ProductID != "" || requestResendOTP.Email != "" {
		if !validations.ValidateEmailFormat(requestResendOTP.Email) {
			errors.BadRequest(W, "Invalid Email address")
			return
		}
		otp, err := customizedNFTFacade.GenerateOTP(requestResendOTP.Email)
		if err != nil {
			errors.BadRequest(W, "Failed to retrive BatchID data")
			return
		} else {
			batchData, err := customizedNFTFacade.GetBatchIDDatabyItemID(requestResendOTP.ProductID)
			if err != nil {
				errors.BadRequest(W, "Failed to retrive BatchID data")
				return
			}
			encodedOTP := commonMethods.StringToSHA256(otp)
			rst, err := UpdateCustomerIndormation(requestResendOTP.Email, batchData.BatchID, encodedOTP)
			if err != nil {
				errors.BadRequest(W, "Failed oepration : "+err.Error())
				return
			}
			err1 := customizedNFTFacade.SendEmail(otp, requestResendOTP.Email)
			if err1 != nil {
				errors.BadRequest(W, "Failed to send email")
				return
			} else {
				commonResponse.SuccessStatus[string](W, rst)
			}

		}

	}
}

func UpdateCustomerIndormation(email string, batchID string, otp string) (string, error) {
	var userAuth models.UserAuth
	userAuth.Email = email
	userAuth.BatchID = batchID
	userAuth.Otp = otp
	userAuth.ExpireDate = primitive.NewDateTimeFromTime(customizedNFTFacade.GenerateOTPExpireDate())
	rst, err := customizedNFTFacade.ResendOTP(userAuth)
	if err != nil {
		return rst, err
	}
	return rst, nil
}
