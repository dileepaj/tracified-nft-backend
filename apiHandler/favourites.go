package apiHandler

import (
	"encoding/json"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/marketplaceBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/middleware"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/gorilla/mux"
)

func CreateFavourites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var favObject models.Favourite
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&favObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
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
	results, err1 := marketplaceBusinessFacade.GetFavouritesByUserPK(vars["user"])
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


func GetFavouritesByBlockchainAndIdentifier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")

	vars := mux.Vars(r)
	if vars["blockchain"] != "" && vars["nftidentifier"] != "" {

		result, id, err := marketplaceBusinessFacade.GetFavouritesByBlockchainAndIdentifier(vars["blockchain"], vars["nftidentifier"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			if len(result) > 5 {
				var hotpicks models.Hotpicks
				hotpicks = models.Hotpicks{
					NFTIdentifier: id,
					HotPicks:      true,
				}
				result, err := marketplaceBusinessFacade.UpdateHotPicks(hotpicks)
				if err != nil {
					errors.BadRequest(w, err.Error())
				} else {
					commonResponse.SuccessStatus[models.NFT](w, result)
				}
			}
			commonResponse.SuccessStatus[[]models.Favourite](w, result)
		}
	} else {
		errors.BadRequest(w, "")
	}

}
