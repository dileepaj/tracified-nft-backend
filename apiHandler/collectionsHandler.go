package apiHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/marketplaceBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"

	//	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/gorilla/mux"
)

func CreateCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	log.Println("-------------------------------------------testing 1 ------------------------------------")
	var createCollectionObject models.NFTCollection
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createCollectionObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	log.Println("-------------------------------------------testing 2 ------------------------------------")
	fmt.Println(createCollectionObject)
	err = validations.ValidateInsertCollection(createCollectionObject)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		log.Println("------------------------------------testing 6 ---------------------------------------------------")
		_, err1 := marketplaceBusinessFacade.CreateCollection(createCollectionObject)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {

			w.WriteHeader(http.StatusOK)
			message := "New Collection Added"
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}
func GetCollectionByUserPK(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	log.Println("---------------------------------------test 1------------------------------")
	fmt.Println(vars["userid"])
	results, err1 := marketplaceBusinessFacade.GetCollectionByUserPK(vars["userid"])
	fmt.Println("results-----------------------", results)
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

func GetCollectionById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	log.Println("---------------------------------------test 1------------------------------")
	fmt.Println(vars["_id"])
	results, err1 := marketplaceBusinessFacade.GetCollectionById(vars["_id"])
	fmt.Println("results-----------------------", results)
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

func GetAllCollections(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	log.Println("calling func Get All Collections....")
	results, err1 := marketplaceBusinessFacade.GetAllCollections()

	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		fmt.Println("results-----------------------", results)
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}
}
func UpdateCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var updateCollectionObj requestDtos.UpdateCollection
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateCollectionObj)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	} else {
		_, err1 := marketplaceBusinessFacade.UpdateCollection(updateCollectionObj)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			message := "Collection Name updated successfully."
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}

func DeleteCollectionById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	log.Println("-----------------------------------------test 1----------------------------------------")
	var deleteCollectionObj requestDtos.DeleteCollectionById
	decoder := json.NewDecoder(r.Body)
	log.Println("-----------------------------------------test 2----------------------------------------", deleteCollectionObj)
	err := decoder.Decode(&deleteCollectionObj)
	log.Println("-----------------------------------------test 3----------------------------------------", err)
	if err != nil {
		log.Println("------------------------------------------------------------error")
		logs.ErrorLogger.Println(err.Error())
	} else {
		log.Println("-----------------------------------------test 4----------------------------------------")
		err1 := marketplaceBusinessFacade.DeleteCollectionById(deleteCollectionObj)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			message := "Collection has been deleted"
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}

}

func DeleteCollectionByUserPK(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var deleteCollectionObj requestDtos.DeleteCollectionByUserPK
	log.Println("-----------------------------------------test 1----------------------------------------")
	decoder := json.NewDecoder(r.Body)
	log.Println("-----------------------------------------test 2----------------------------------------")
	err := decoder.Decode(&deleteCollectionObj)
	log.Println("-----------------------------------------test 2----------------------------------------", err)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	} else {
		log.Println("-----------------------------------------test 3----------------------------------------")
		err1 := marketplaceBusinessFacade.DeleteCollectionByUserPK(deleteCollectionObj)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			message := "Collections have been deleted"
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}
