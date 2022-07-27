package apiHandler

import (
	"encoding/json"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/ruriBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/gorilla/mux"
)

/**
 **Description:function is used to generate and savee a new OTP and send a email to the customer with the otp. The function will also retreive item data such
 **as itemID,batchID etc by using the provided productID parameter
 **Returns:Object ID of the new OTP created
 */
func RuriNewOTP(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	if vars["productID"] != "" || vars["email"] != "" { // Checks if the API has the ncessary params filled
		logs.InfoLogger.Println(vars["productID"], vars["email"])
		otp, err := ruriBusinessFacade.GenerateOTP(vars["email"]) // Generate OTP
		if err != nil || otp == "" {                              // Triggered if there is an error generating the OTP
			errors.BadRequest(W, "Failed to generate OTP")
		} else { // If the OTP is generated ProductID will be sent to get the batchID
			batchData, err := ruriBusinessFacade.GetBatchIDDatabyItemID(vars["productID"]) //GET batch id when RURI product ID is given
			// If the API is unable to retreive the batch ID response with err msg will be sent
			if err != nil || otp == "" {
				errors.BadRequest(W, "Failed to retrive BatchID data")
				return
			} //If batch ID is retreived user email,otp and batch ID will be sent to be saved in the ruriOtp DB
			var userAuth models.UserAuth
			userAuth.Email = vars["email"]
			userAuth.Otp = otp
			userAuth.BatchID = batchData.BatchID
			result, error := ruriBusinessFacade.SaveOTP(userAuth)
			if error != nil {
				errors.BadRequest(W, "")
			} else {
				commonResponse.SuccessStatus[string](W, result)
			}
		}
	} else { //* If necssary params are not there Error Message will be sent as a response
		errors.BadRequest(W, "Product ID or email is missing")
	}
}

/**
 **Description:This function is used to validate a OTP provided by the user. The email and the otp will be sent where the api.Checks if the otp and the email recived
 *				*have a matching record in the DB
 **Returns:If There is a matching email and a OTP in the DB the generated SVG will be reutrned as a response. If not error msg is sent as response
 */
func ValidateOTP(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	if vars["email"] != "" || vars["otp"] != "" { //Checks if the API has the ncessary params filled
		rst, err := ruriBusinessFacade.ValidateOTP(vars["email"], vars["otp"])
		logs.InfoLogger.Println("rst: ", rst)
		if err != nil {
			errors.BadRequest(W, "Failed to validate OTP")
			return
		}
		if rst == "Invalid OTP" {
			errors.BadRequest(W, rst)
			return
		}
		var tempBatchID = "RURI_VSAPPH_013" //? Templary hardcoded
		rst1, err1 := ruriBusinessFacade.GenerateandSaveSVG(tempBatchID, vars["email"])
		if err != nil {
			errors.BadRequest(W, err1.Error())
			return
		}
		commonResponse.SuccessStatus[responseDtos.SVGforNFTResponse](W, rst1)
	} else {
		errors.BadRequest(W, "email or OTP missing")
		return
	}
}

/**
 **Description : Is used to generate a new OTP code and send to the provided email address
 **Returns : If the OTP is succesfully generated the batch ID is returned.
 */
func ResentOTP(W http.ResponseWriter, r *http.Request) {
	var userAuth models.UserAuth
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	if vars["productID"] != "" || vars["email"] != "" {
		otp, err := ruriBusinessFacade.GenerateOTP(vars["email"])
		if err != nil {
			errors.BadRequest(W, "Failed to retrive BatchID data")
			return
		} else {
			batchData, err := ruriBusinessFacade.GetBatchIDDatabyItemID(vars["productID"])
			if err != nil {
				errors.BadRequest(W, "Failed to retrive BatchID data")
				return
			}
			userAuth.Email = vars["email"]
			userAuth.BatchID = batchData.BatchID
			userAuth.Otp = otp
			rst, err := ruriBusinessFacade.ResendOTP(userAuth)
			if err != nil {
				errors.BadRequest(W, "Failed to retrive BatchID data")
				return
			}
			commonResponse.SuccessStatus[string](W, rst)
		}

	}
}

/**
 ** Description : Updaes the DB with the image hash provided. Seareches by mongo DB object ID
 ** Returns : returns the svg if successfully updated.
 */
func UpdateSVGUserMappingbySha256(W http.ResponseWriter, r *http.Request) {
	var UpdateSVG models.UserNFTMapping
	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode((&UpdateSVG))
	if err != nil {
		logs.ErrorLogger.Println("Error Decoding request body data : ", err.Error())
		errors.BadRequest(W, err.Error())
		return
	} else {
		rst, err := ruriBusinessFacade.UpdateUserMappingbySha256(UpdateSVG)
		if err != nil {
			logs.ErrorLogger.Println("Error While updating hash : ", err.Error())
			errors.BadRequest(W, err.Error())
			return
		}
		commonResponse.SuccessStatus[responseDtos.SVGforNFTResponse](W, rst)

	}
}

/**
 ** Description : Will be used to get the the svg by the hash
 ** Reutrns : will reutrn the  SVG is a valide hash is provided
 */
func GetSVGbySha256(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	if vars["hash"] != "" {
		rst, err := ruriBusinessFacade.GetSVGbySha256(vars["hash"])
		if err != nil {
			errors.BadRequest(W, err.Error())
		} else {
			commonResponse.SuccessStatus[string](W, rst)
		}
	}
}

//! TEsting methods remove after full impl
func SaveTDPDataByBatchID(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	logs.InfoLogger.Println("URL Param:", vars["batchID"])
	ruriBusinessFacade.GetTDPDataByBatchID(vars["batchID"])
}

//! TEsting methods remove after full impl
func GenerateSVG(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	logs.InfoLogger.Println("URL Param:", vars["batchID"])
	rst, _ := ruriBusinessFacade.GenerateandSaveSVG(vars["batchID"], vars["email"])
	commonResponse.SuccessStatus[string](W, rst.SVG)
}
