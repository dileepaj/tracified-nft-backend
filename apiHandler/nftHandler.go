package apiHandler

import (
	"encoding/json"
	"fmt"
	"log"
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

func CreateNFT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var test models.NFT

	//var requestNFTObject requestDtos.CreateNFTRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&test)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}

	log.Println("---------------------test-------------", test)
	log.Println("--------------------------------------------------------------inside nft store---------------------", test)
	err = validations.ValidateRequestNFTObject(test)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {

		result, err := marketplaceBusinessFacade.StoreNFT(test)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}

func CreateOwner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var test2 models.Ownership
	log.Println("-------------------------------------------testing 1 ------------------------------------")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&test2)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	log.Println("-------------------------------------------testing 2 ------------------------------------")
	fmt.Println(test2)
	log.Println("------------------------------------testing 6 ---------------------------------------------------")
	_, err1 := marketplaceBusinessFacade.StoreOwner(test2)
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {

		w.WriteHeader(http.StatusOK)
		message := "New owner Added"
		err = json.NewEncoder(w).Encode(message)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}

}

func MakeSale(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
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

func GetAllONSaleNFT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
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

func GetBlockchainSpecificNFT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
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

func GetNFTbyTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
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

func GetWatchListNFT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
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

func GetNFTByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
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

func GetNFTByTenentName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
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

func CreateTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	log.Println("-------------------------------------------testing 1 ------------------------------------")
	var tags models.Tags
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tags)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	log.Println("-------------------------------------------testing 2 ------------------------------------")
	fmt.Println(tags)
	_, err1 := marketplaceBusinessFacade.CreateTags(tags)
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {

		w.WriteHeader(http.StatusOK)
		message := "New Tags Added"
		err = json.NewEncoder(w).Encode(message)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}

}

func GetAllTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	log.Println("calling func Get All Tags....")
	results, err1 := marketplaceBusinessFacade.GetAllTags()

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

func GetTagsByNFTIdentifier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	log.Println("---------------------------------------test 1------------------------------")
	fmt.Println(vars["nftidentifier"])
	results, err1 := marketplaceBusinessFacade.GetTagsByNFTIdentifier(vars["nftidentifier"])
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
