package apiHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/marketplaceBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func CreateWatchList(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var requestCreateWatchList models.WatchList
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestCreateWatchList)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateInsertWatchList(requestCreateWatchList)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err1 := marketplaceBusinessFacade.CreateWatchList(requestCreateWatchList)
		if err1 != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}

func GetWatchListByUserPK(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	results, err1 := marketplaceBusinessFacade.GetWatchListByUserPK(vars["user"])
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

func GetAllWatchLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	results, err1 := marketplaceBusinessFacade.GetAllWatchLists()

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

func FindWatchListsByBlockchainAndIdentifier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["blockchain"] != "" || vars["nftidentifier"] != "" {
		result, id, err := marketplaceBusinessFacade.FindWatchListsByBlockchainAndIdentifier(vars["blockchain"], vars["nftidentifier"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			if len(result) > 5 {
				var trend models.Trending
				trend = models.Trending{
					NFTIdentifier: id,
					Trending:      true,
				}
				result, err := marketplaceBusinessFacade.UpdateTrending(trend)
				if err != nil {
					errors.BadRequest(w, err.Error())
				} else {
					commonResponse.SuccessStatus[models.NFT](w, result)
				}
			}
			commonResponse.SuccessStatus[[]models.WatchList](w, result)

		}
	} else {
		errors.BadRequest(w, "")
	}
}

func GetWatchListsByBlockchainAndIdentifier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["blockchain"] != "" || vars["nftidentifier"] != "" {
		result, id, err := marketplaceBusinessFacade.FindWatchListsByBlockchainAndIdentifier(vars["blockchain"], vars["nftidentifier"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			log.Println("Id is: ", id)
			commonResponse.SuccessStatus[[]models.WatchList](w, result)

		}
	} else {
		errors.BadRequest(w, "")
	}
}
