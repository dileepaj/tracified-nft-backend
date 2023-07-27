package apiHandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/marketplaceBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/gorilla/mux"
)

func SaveWalletNFTStates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var nft models.NFTWalletState
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&nft)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateInsertNftState(nft)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err := marketplaceBusinessFacade.StoreNFTState(nft)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}

func SaveWalletNFTTXNs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var txn models.NFTWalletStateTXN
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&txn)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}

	err = validations.ValidateInsertNftStateTXN(txn)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err := marketplaceBusinessFacade.StoreNFTStateTXN(txn)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}

func UpdateWalletNFTState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var updateObj requestDtos.UpdateNFTState
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateObj)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		ErrorMessage := err.Error()
		errors.BadRequest(w, ErrorMessage)
	} else {
		_, err1 := marketplaceBusinessFacade.UpdateNFTState(updateObj)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			message := "NFT State updated successfully."
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}

func DeleteWalletNFTByIssuer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var deleteNftStateObj requestDtos.DeleteNFTState
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&deleteNftStateObj)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	} else {
		deletedCount, err1 := marketplaceBusinessFacade.DeleteNFTStateByIssuer(deleteNftStateObj)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			if deletedCount == 0 {
				w.WriteHeader(http.StatusNoContent)
				if err != nil {
					logs.ErrorLogger.Println(err)
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			message := "NFT State have been deleted"
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}

func GetWalletTxnsByIssuer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	results, err1 := marketplaceBusinessFacade.GetWalletNFTTxnsByIssuer(vars["issuerpublickey"])
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

func GetWalletNFTByStateAndCurrentOwner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	var pagination requestDtos.WalletNFTsForMatrixView
	pagination.Blockchain = r.URL.Query().Get("blockchain")
	var StateToSearch, stateErr = strconv.Atoi(r.URL.Query().Get("nftstatus"))
	if stateErr != nil {
		errors.BadRequest(w, "Invalid NFT state")
		return
	}
	var pubKey = r.URL.Query().Get("currentowner")
	pgsize, err1 := strconv.Atoi(r.URL.Query().Get("pagesize"))
	if err1 != nil {
		errors.BadRequest(w, "Requested invalid page size.")
		return
	}
	pagination.PageSize = int32(pgsize)
	requestedPage, err2 := strconv.Atoi(r.URL.Query().Get("requestedPage"))
	if err2 != nil {
		errors.BadRequest(w, "Requested page error")
		return
	}
	pagination.RequestedPage = int32(requestedPage)
	pagination.SortbyFeild = "blockchain"
	logs.InfoLogger.Println("Received pagination requested: ", pagination)
	results, err := marketplaceBusinessFacade.GetWalletNFTByState(pagination, StateToSearch, pubKey)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if pagination.PageSize <= 0 {
			errors.BadRequest(w, "Page size should be greater than zero")
			return
		}
		if pagination.RequestedPage < 0 {
			errors.BadRequest(w, "Requested page size should be greater than zero")
			return
		}
		if results.PaginationInfo.TotalPages < pagination.RequestedPage {
			errors.BadRequest(w, "requested page does not exist")
			return
		}
		if results.Content == nil {
			errors.BadRequest(w, "Collection does not have any NFTs")
			return
		}
		commonResponse.SuccessStatus[models.PaginateWalletNFTResponse](w, results)
	}
}

func GetWalletNFTByStatusAndRequested(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	var pagination requestDtos.WalletNFTsForMatrixView
	pagination.Blockchain = r.URL.Query().Get("blockchain")
	CollectionToSearch, stateErr := strconv.Atoi(r.URL.Query().Get("nftstatus"))
	if stateErr != nil {
		errors.BadRequest(w, "Invalid State.")
		return
	}
	var pubKey = r.URL.Query().Get("nftrequested")
	pgsize, err1 := strconv.Atoi(r.URL.Query().Get("pagesize"))
	if err1 != nil {
		errors.BadRequest(w, "Requested invalid page size.")
		return
	}
	pagination.PageSize = int32(pgsize)
	requestedPage, err2 := strconv.Atoi(r.URL.Query().Get("requestedPage"))
	if err2 != nil {
		errors.BadRequest(w, "Requested page error")
		return
	}
	pagination.RequestedPage = int32(requestedPage)
	pagination.SortbyFeild = "blockchain"
	logs.InfoLogger.Println("Received pagination requested: ", pagination)
	results, err := marketplaceBusinessFacade.GetWalletNFTByStateForRequested(pagination, CollectionToSearch, pubKey)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if pagination.PageSize <= 0 {
			errors.BadRequest(w, "Page size should be greater than zero")
			return
		}
		if pagination.RequestedPage < 0 {
			errors.BadRequest(w, "Requested page size should be greater than zero")
			return
		}
		if results.PaginationInfo.TotalPages < pagination.RequestedPage {
			errors.BadRequest(w, "requested page does not exist")
			return
		}
		if results.Content == nil {
			errors.BadRequest(w, "Collection does not have any NFTs")
			return
		}
		commonResponse.SuccessStatus[models.PaginateWalletNFTResponse](w, results)
	}
}
