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
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func VerifyFavouriteTogglebUserPK(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	rst, err := marketplaceBusinessFacade.VerifyFavouriteTogglebUserPK(vars["blockchain"], vars["user"], vars["nftidentifer"])
	if err != nil {
		logs.ErrorLogger.Print("Failed to perform favourites retrival for user : ", err.Error())
		errors.BadRequest(w, err.Error())
		return
	}
	if rst.User == "" {
		rst.User = "Add to favourite"
		w.WriteHeader(http.StatusOK)
		err1 := json.NewEncoder(w).Encode(rst)
		if err1 != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	err1 := json.NewEncoder(w).Encode(rst)
	if err1 != nil {
		logs.ErrorLogger.Println(err)
	}
}

func RemoveUserfromFavourite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	objectID, err := primitive.ObjectIDFromHex(vars["favouriteID"])
	if err != nil {
		logs.WarningLogger.Println("Invalid favouriteID: ", err.Error())
		errors.BadRequest(w, "Invalid favouriteID")
		return
	}
	rst, delerr := marketplaceBusinessFacade.RemoveUserFromFavourites(objectID)
	if delerr != nil {
		logs.WarningLogger.Println("Failed to remove user from favouriteID: ", err.Error())
		errors.BadRequest(w, "Failed to remove user from favourites")
		return
	}
	if rst > 0 {
		w.WriteHeader(http.StatusOK)
		err1 := json.NewEncoder(w).Encode("Removed from favourites")
		if err1 != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	} else {
		errors.BadRequest(w, "Failed to remove user from favourites")
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

func FavouritesByBlockchainAndIdentifier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["blockchain"] != "" && vars["nftidentifier"] != "" {

		result, id, err := marketplaceBusinessFacade.GetFavouritesByBlockchainAndIdentifier(vars["blockchain"], vars["nftidentifier"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			log.Println("Id is: ", id)
			commonResponse.SuccessStatus[[]models.Favourite](w, result)
		}
	} else {
		errors.BadRequest(w, "")
	}

}
