package apiHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/marketplaceBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/gorilla/mux"

	//	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
)

func CreateFavourites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var favObject models.Favourite
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&favObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	fmt.Println(favObject)
	err = validations.ValidateInsertFavourites(favObject)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		_, err1 := marketplaceBusinessFacade.CreateFavourites(favObject)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {

			w.WriteHeader(http.StatusOK)
			message := "New Favourite Added"
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}

func GetFavouritesByUserPK(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	fmt.Println(vars["userid"])
	results, err1 := marketplaceBusinessFacade.GetFavouritesByUserPK(vars["userid"])
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

func GetAllFavourites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	log.Println("calling func Get All Favourites....")
	results, err1 := marketplaceBusinessFacade.GetAllFavourites()

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
