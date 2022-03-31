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
	"github.com/dileepaj/tracified-nft-backend/utilities/middleware"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/gorilla/mux"
)

func CreateNFT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		var requestNFTObject requestDtos.CreateNFTRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&requestNFTObject)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
		}
		err = validations.ValidateRequestNFTObject(requestNFTObject)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			result, err := marketplaceBusinessFacade.StoreNFT(requestNFTObject)
			if err != nil {
				errors.BadRequest(w, err.Error())
			} else {
				commonResponse.SuccessStatus[string](w, result)
			}
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func MakeSale(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		var makeSaleRequestObject requestDtos.UpdateNFTSALERequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&makeSaleRequestObject)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
		}
		err = validations.ValidateMakeSale(makeSaleRequestObject)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			result, err := marketplaceBusinessFacade.MakeSaleNFT(makeSaleRequestObject)
			if err != nil {
				errors.BadRequest(w, err.Error())
			} else {
				commonResponse.SuccessStatus[responseDtos.ResponseNFTMakeSale](w, result)
			}
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func GetAllONSaleNFT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		vars := mux.Vars(r)
		if vars["status"] != "" || vars["userpk"] != "" {
			results, err := marketplaceBusinessFacade.GetAllONSaleNFT(vars["status"], vars["userpk"])
			if err != nil {
				errors.BadRequest(w, err.Error())
			} else {
				commonResponse.SuccessStatus[[]models.NFT](w, results)
			}
		} else {
			errors.BadRequest(w, "")
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func GetBlockchainSpecificNFT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		vars := mux.Vars(r)
		if vars["blockchain"] != "" {
			results, err := marketplaceBusinessFacade.GetBlockchainSpecificNFT(vars["blockchain"])
			if err != nil {
				errors.BadRequest(w, err.Error())
			} else {
				commonResponse.SuccessStatus[[]models.NFT](w, results)
			}
		} else {
			errors.BadRequest(w, "")
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func GetNFTbyTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		vars := mux.Vars(r)
		if vars["tags"] != "" {
			results, err := marketplaceBusinessFacade.GetNFTbyTagsName(vars["tags"])
			if err != nil {
				errors.BadRequest(w, err.Error())
			} else {
				commonResponse.SuccessStatus[[]models.NFT](w, results)
			}
		} else {
			errors.BadRequest(w, "")
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func GetWatchListNFT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		vars := mux.Vars(r)
		if vars["userId"] != "" {
			results, err := marketplaceBusinessFacade.GetWatchListNFT(vars["userId"])
			if err != nil {
				errors.BadRequest(w, err.Error())
			} else {
				commonResponse.SuccessStatus[[]models.NFT](w, results)
			}
		} else {
			errors.BadRequest(w, "")
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func GetNFTByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		vars := mux.Vars(r)
		if len(vars["userId"]) != 0 {
			result, err := marketplaceBusinessFacade.GetNFTbyAccount(vars["userId"])
			if err != nil {
				errors.BadRequest(w, err.Error())
			} else {
				commonResponse.SuccessStatus[[]models.NFT](w, result)
			}
		} else {
			errors.BadRequest(w, "")
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func GetNFTByTenentName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		vars := mux.Vars(r)
		if len(vars["tenentname"]) != 0 {
			result, err := marketplaceBusinessFacade.GetNFTbyTenentName((vars["tenentname"]))
			if err != nil {
				errors.BadRequest(w, err.Error())
			} else {
				commonResponse.SuccessStatus[[]models.NFT](w, result)
			}
		} else {
			errors.BadRequest(w, "")
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}
