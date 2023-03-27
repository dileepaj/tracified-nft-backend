package apiHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/marketplaceBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/gorilla/mux"
)

func CreateCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var createCollectionObject models.NFTCollection
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createCollectionObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateInsertCollection(createCollectionObject)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
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

func CreateSVG(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var createSVGObject models.SVG
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createSVGObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateInsertSVG(createSVGObject)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		rst, err1 := marketplaceBusinessFacade.CreateSVG(createSVGObject)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {

			w.WriteHeader(http.StatusOK)
			//message := "New SVG Added"
			err = json.NewEncoder(w).Encode(rst)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}

func UpdateSvgBlockChain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var updateSVGObject models.SVG
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateSVGObject)
	logs.InfoLogger.Println("Starting svg BC Update: ", updateSVGObject)
	if err != nil {
		logs.ErrorLogger.Println("Error while decoding into json in UpdateSvg:collectionHandler: " + err.Error())
	}
	err = validations.ValidateInsertSVG(updateSVGObject)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		_, err = marketplaceBusinessFacade.UpdateSVGBlockchain(updateSVGObject)
		if err != nil {
			Errormsg := err.Error()
			errors.BadRequest(w, Errormsg)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			message := "SVG Block chain has been Updated"
			err := json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println("Error occured while encoding JSON in UpdateSvgBlockChain(collectionHandlers):", err.Error())
			}
			return
		}
	}
}

func GetCollectionByUserPK(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	results, err1 := marketplaceBusinessFacade.GetCollectionByUserPK(vars["userid"])
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

func GetCollectionByPublicKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	results, err1 := marketplaceBusinessFacade.GetCollectionByPublicKey(vars["publickey"])
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

func DeleteCollectionByUserPK(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var deleteCollectionObj requestDtos.DeleteCollectionByUserPK
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&deleteCollectionObj)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	} else {
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

func GetCollectionByUserPKAndMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	fmt.Println("we hwre ", vars["userid"], vars["publickey"])
	results, err1 := marketplaceBusinessFacade.GetCollectionByUserPKByMail(vars["userid"], vars["publickey"])
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
