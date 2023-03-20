package apiHandler

import (
	"encoding/json"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/marketplaceBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
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
func CreateFaq(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var requestCreateFaq models.Faq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestCreateFaq)
	if err != nil {
		logs.ErrorLogger.Println("Error occured while decoding JSON in CreateFaq:faqHandler: ", err.Error())
	}
	err = validations.ValidateFaq(requestCreateFaq)
	if err != nil {
		errors.BadRequest(W, err.Error())
	} else {
		result, err1 := marketplaceBusinessFacade.CreateFaq(requestCreateFaq)
		if err1 != nil {
			errors.BadRequest(W, err.Error())
		} else {
			commonResponse.SuccessStatus[string](W, result)
		}
	}
}

// Trigger the GetAllFaq() method that will return all the FAQs
func GetAllFaq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	results, err := marketplaceBusinessFacade.GetAllFaq()
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

// Trigger the GetFaqByID() method that will return The specific FAQ with the ID passed via the API
func GetFaqByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	result, err := marketplaceBusinessFacade.GetFaqByID(vars["_id"])
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		commonResponse.SuccessStatus[models.Faq](w, result)
	}

}

// Retreives and decodees the Object ID and update from the API and invokes the UpdateQuestion() To update contents in a Specific FAQ
func UpdateFaqbyID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var updateFaq requestDtos.UpdateFaq
	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode((&updateFaq))
	if err != nil {
		logs.ErrorLogger.Println("Error occured while decoding JSON in UpdateFaqbyID(FaqHandler):", err.Error())
	} else {
		_, err = marketplaceBusinessFacade.UpdateFaqbyID(updateFaq)
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

func UpdateUserFAQStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var updateObj requestDtos.UpdateUserFAQ
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateObj)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	} else {
		_, err1 := marketplaceBusinessFacade.UpdateUserFAQ(updateObj)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			//If userfaq status is updated successfully another DB call is made to retive all the faq detials by id
			faqdata, err1 := marketplaceBusinessFacade.GetUserFAQByID(updateObj.UserQuestionID)
			logs.InfoLogger.Println("UserFAQ data recived: ", faqdata)
			if err1 != nil {
				logs.ErrorLogger.Println("Failed to get endorsment data : ", err1.Error())
			}
			emailErr := marketplaceBusinessFacade.SendResponseToFAQ(faqdata)
			if emailErr != nil {
				errors.InternalError(w, emailErr.Error())
				return
			}
			w.WriteHeader(http.StatusOK)
			message := "UserFAQ updated successfully."
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}

func StoreFAQ(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var requestStoreFaq models.UserQuestions
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestStoreFaq)
	if err != nil {
		logs.ErrorLogger.Println("Error occured while decoding JSON in CreateFaq:faqHandler: ", err.Error())
	}
	err = validations.ValidateUserFaq(requestStoreFaq)
	if err != nil {
		errors.BadRequest(W, err.Error())
	} else {
		result, err1 := marketplaceBusinessFacade.StoreUserFAQ(requestStoreFaq)
		if err1 != nil {
			errors.BadRequest(W, err.Error())
		} else {
			commonResponse.SuccessStatus[string](W, result)
		}
	}
}

func GetUserFAQbyStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["status"] != "" {
		results, err := marketplaceBusinessFacade.GetUserFAQByStatus(vars["status"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[[]responseDtos.GetPendingUserFAQ](w, results)
		}
	} else {
		errors.BadRequest(w, "")
	}
}

func GetUserFAQAttachmentbyID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["qid"] != "" {
		rst, err := marketplaceBusinessFacade.GetFAQAttachmentbyID(vars["qid"])
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			errors.BadRequest(w, "failed to get image")
			return
		}
		if len(rst.Attachment) > 0 {
			commonResponse.SuccessStatus[string](w, rst.Attachment)
		} else {
			commonResponse.SuccessStatus[string](w, "No Image")
		}

	}
}
