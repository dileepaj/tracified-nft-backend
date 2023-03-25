package apiHandler

import (
	"encoding/json"
	"net/http"
	"time"

	customizedNFTFacade "github.com/dileepaj/tracified-nft-backend/businessFacade/customizedNFTFacade"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
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
	err := decoder.Decode(&requestGenOTP)
	if err != nil {
		errors.BadRequest(W, "Invalid data")
		return
	}
	// Checks if the API has the ncessary params filled
	if requestGenOTP.ProductID != "" || requestGenOTP.Email != "" {
		logs.InfoLogger.Println(requestGenOTP.ProductID, requestGenOTP.Email)
		// Generate OTP
		otp, err := customizedNFTFacade.GenerateOTP(requestGenOTP.Email)
		// Triggered if there is an error generating the OTP
		if err != nil || otp == "" {
			errors.BadRequest(W, "Failed to generate OTP")
			// If the OTP is generated ProductID will be sent to get the batchID
		} else {
			//GET batch id when RURI product ID is given
			batchData, err := customizedNFTFacade.GetBatchIDDatabyItemID(requestGenOTP.ProductID)
			// If the API is unable to retreive the batch ID response with err msg will be sent
			if err != nil || otp == "" {
				errors.BadRequest(W, "Failed to retrive BatchID data")
				return
			} //If batch ID is retreived user email,otp and batch ID will be sent to be saved in the ruriOtp DB
			_, error := SaveUserOTPMapping(requestGenOTP.Email, otp, batchData.BatchID)
			if error != nil {
				errors.BadRequest(W, error.Error())
				return
			} else {

				// err1 := customizedNFTFacade.SendEmail(otp, requestGenOTP.Email)
				// if err1 != nil {
				// 	errors.BadRequest(W, "Incorrect email address")
				// } else {
				commonResponse.SuccessStatus[string](W, "Email Has been Sent to "+requestGenOTP.Email)
				//}
			}
		}
	} else { //* If necssary params are not there Error Message will be sent as a response
		errors.BadRequest(W, "Product ID or email is missing")
	}
}
func SaveUserOTPMapping(email string, otp string, batchID string) (string, error) {
	var userAuth models.UserAuth
	userAuth.Email = email
	userAuth.Otp = otp
	userAuth.BatchID = batchID //Default number of attempts
	userAuth.Validated = "False"
	userAuth.ExpireDate = primitive.NewDateTimeFromTime(GenerateOTPExpireDate())
	logs.InfoLogger.Println("NEW exp date created: ", userAuth.ExpireDate)
	result, error := customizedNFTFacade.SaveOTP(userAuth)
	if error != nil {
		return result, error
	} else {
		return result, nil
	}
}
func GenerateOTPExpireDate() time.Time {
	currentDate := time.Now()
	logs.InfoLogger.Println("OTP Generated on : ", currentDate)
	duration := time.Hour * 24 * 30
	expireDate := currentDate.Add(duration)
	logs.InfoLogger.Println("Expiration Date : ", expireDate)
	return expireDate
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
	//Checks if the API has the ncessary params filled
	if requestValidateOTP.OTPCode != "" || requestValidateOTP.Email != "" {
		status, errs := customizedNFTFacade.GetNFTStatus(requestValidateOTP.Email, requestValidateOTP.OTPCode)
		if errs != nil {
			errors.BadRequest(W, "Failed to NFT Status")
			return
		} else {
			logs.InfoLogger.Println("status: ", status)
			if status == "Minted" {
				errors.BadRequest(W, "NFT already Minted")
				return
			}
			rst, err := customizedNFTFacade.ValidateOTP(requestValidateOTP.Email, requestValidateOTP.OTPCode)
			logs.InfoLogger.Println("rst: ", rst)
			if err != nil {
				errors.BadRequest(W, "Failed to validate OTP")
				return
			}
			if rst == "Invalid OTP" {
				errors.BadRequest(W, rst)
				return
			}
			var tempBatchID = rst
			rst1, err1 := customizedNFTFacade.GetSVGbyEmailandBatchID(requestValidateOTP.Email, tempBatchID)
			if err != nil {
				errors.BadRequest(W, err1.Error())
				return
			}
			commonResponse.SuccessStatus[string](W, rst1.SVG)
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
	err := decoder.Decode(&requestResendOTP)
	if err != nil {
		errors.BadRequest(W, "Invalid data")
		return
	}
	if requestResendOTP.ProductID != "" || requestResendOTP.Email != "" {
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
			rst, err := UpdateCustomerIndormation(requestResendOTP.Email, batchData.BatchID, otp)
			if err != nil {
				errors.BadRequest(W, "failed to send email!")
				return
			}
			err1 := customizedNFTFacade.SendEmail(otp, requestResendOTP.Email)
			if err1 != nil {
				errors.BadRequest(W, "incorrect email address")
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
	rst, err := customizedNFTFacade.ResendOTP(userAuth)
	if err != nil {
		return rst, err
	}
	return rst, nil
}
