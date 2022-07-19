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

/*
	All functions here a triggered by api Calls and invokes respective marketplaceBusinessFace Class Methods
*/
//Retrevies data from the Json Body and decodes it into a model class (Faq).Thenthe CreateFaq() method is invoked from marketplaceBusinessFacade
func CreatePartner(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var requestCreatePartner models.Partner
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestCreatePartner)
	if err != nil {
		logs.ErrorLogger.Println("Error occured while decoding JSON in CreateFaq:faqHandler: ", err.Error())
	}
	err = validations.ValidatePartner(requestCreatePartner)
	if err != nil {
		errors.BadRequest(W, err.Error())
	} else {
		result, err1 := marketplaceBusinessFacade.CreatePartner(requestCreatePartner)
		if err1 != nil {
			errors.BadRequest(W, err.Error())
		} else {
			commonResponse.SuccessStatus[string](W, result)
		}
	}
}

//Trigger the GetAllFaq() method that will return all the FAQs
func GetAllPartner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	results, err := marketplaceBusinessFacade.GetAllPartner()
	if err != nil {
		ErrorMessage := err.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while encoding JSON in GetAllFaq(FaqHandler): ", err.Error())
		}
		return
	}
}

//Trigger the GetFaqByID() method that will return The specific FAQ with the ID passed via the API
func GetPartnerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := marketplaceBusinessFacade.GetPartnerByID(vars["_id"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		commonResponse.SuccessStatus[models.Partner](w, result)
	}

}

// Retreives and decodees the Object ID and update from the API and invokes the UpdateQuestion() To update contents in a Specific FAQ
func UpdatePartnerbyID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var updatePartner requestDtos.UpdatePartner
	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode((&updatePartner))
	if err != nil {
		logs.ErrorLogger.Println("Error occured while decoding JSON in UpdateFaqbyID(FaqHandler):", err.Error())
	} else {
		_, err = marketplaceBusinessFacade.UpdatePartnerbyID(updatePartner)
		if err != nil {
			Errormsg := err.Error()
			errors.BadRequest(w, Errormsg)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			message := "FAQ Status has been Updated"
			err := json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println("Error occured while encoding JSON in UpdateFaqbyID(FaqHandler):", err.Error())
			}
			return
		}
	}

}
