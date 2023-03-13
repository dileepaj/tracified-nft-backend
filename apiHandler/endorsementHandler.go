package apiHandler

import (
	"encoding/json"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/marketplaceBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/gorilla/mux"
)

func CreateEndorsement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var endorse models.Endorse
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&endorse)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}

	err = validations.ValidateInsertEndorsement(endorse)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		user, err := marketplaceBusinessFacade.GetEndorsedStatus(endorse.PublicKey)
		if err != nil {
			logs.ErrorLogger.Println("Fialed to Check User :", err.Error())
			errors.BadRequest(w, "User Registration failed")
		}
		if user.PublicKey == "" {
			result, err := marketplaceBusinessFacade.StoreEndorse(endorse)
			if err != nil {
				errors.BadRequest(w, err.Error())
				return
			} else {
				commonResponse.SuccessStatus[string](w, result)
			}
			return
		}
		_, updateErr := marketplaceBusinessFacade.UpdateExsistingUserStatus(endorse)
		if updateErr != nil {
			errors.BadRequest(w, "Error occured")
			return
		}
		commonResponse.SuccessStatus[string](w, "Re-endorsment request Saved")

	}
}

func GetEndorsementbyStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["status"] != "" {
		results, err := marketplaceBusinessFacade.GetEndorsementByStatus(vars["status"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[[]models.Endorse](w, results)
		}
	} else {
		errors.BadRequest(w, "")
	}
}

func GetEndorsedStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	results, err1 := marketplaceBusinessFacade.GetEndorsedStatus(vars["publickey"])
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}

}

func UpdateEndorsedStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var updateObj requestDtos.UpdateEndorsementByPublicKey
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateObj)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	} else {
		_, err1 := marketplaceBusinessFacade.UpdateEndorsement(updateObj)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			//If endrosment status is updated successfully another DB call is made to retive all the enrosment detials by user
			endorsmentrst, err1 := marketplaceBusinessFacade.GetEndorsmentByUserPK(updateObj.PublicKey)
			if err1 != nil {
				logs.ErrorLogger.Println("Failed to get endorsment data : ", err1.Error())
			}
			emailErr := marketplaceBusinessFacade.SendEndorsmentEmail(endorsmentrst)
			if emailErr != nil {
				errors.InternalError(w, emailErr.Error())
				return
			}
			w.WriteHeader(http.StatusOK)
			message := "Endorsement updated successfully."
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}

func UpdateEndorsement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var updateObj requestDtos.UpdateEndorsement
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateObj)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	} else {
		_, err1 := marketplaceBusinessFacade.UpdateSetEndorsement(updateObj)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			message := "Endorsement updated successfully."
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}
