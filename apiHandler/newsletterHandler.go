package apiHandler

import (
	"encoding/json"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/marketplaceBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/gorilla/mux"
)

func CreateNewsLetter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var requestCreateNewsLetter models.NewsLetter
	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode(&requestCreateNewsLetter)
	if err != nil {
		logs.ErrorLogger.Println("Error While Decoding JSON in CreateNewsLetter:newsletterHandler : ", err.Error())
	}
	err = validations.ValidateInsertNewsLetter(requestCreateNewsLetter)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err1 := marketplaceBusinessFacade.CreateNewsLetter(requestCreateNewsLetter)
		if err1 != nil {
			errors.BadRequest(w, err1.Error())
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}

func GetAllNewsLetters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	results, err := marketplaceBusinessFacade.GetAllNewsLetters()
	if err != nil {
		ErrorMessage := err.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			logs.ErrorLogger.Println("Error while encoding JSON in GetAllNewsLetters:newsLetterHandler: ", err)
		}
		return
	}
}

func GetNewslettersByAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	results, err := marketplaceBusinessFacade.GetNewsLetterByAuthor(vars["name"])
	if err != nil {
		ErrorMessage := err.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			logs.ErrorLogger.Println("Error while encoding JSON in GetNewslettersByAuthor:newsLetterHandler: ", err)
		}
		return
	}

}

func GetNewsletterByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	results, err := marketplaceBusinessFacade.GetNewsletterByID(vars["_id"])
	if err != nil {
		ErrorMessage := err.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		commonResponse.SuccessStatus[models.NewsLetter](w, results)
		return
	}
}

func Subscribe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var updateObj models.Subscription
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateObj)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	} else {
		_, err1 := marketplaceBusinessFacade.AddSubscription(updateObj)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			emailErr := marketplaceBusinessFacade.SubscribeToNewsLetter(updateObj.UserMail)
			if emailErr != nil {
				errors.InternalError(w, emailErr.Error())
				return
			}
			w.WriteHeader(http.StatusOK)
			message := "Subscription added successfully."
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}

func CheckIfSubscribed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	rst, err := marketplaceBusinessFacade.CheckIfSubscribed(vars["email"])
	if err != nil {
		w.WriteHeader(http.StatusOK)
		message := "not subscribed"
		err = json.NewEncoder(w).Encode(message)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	message := rst
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
}
