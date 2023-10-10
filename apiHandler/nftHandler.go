package apiHandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/customizedNFTFacade"
	"github.com/dileepaj/tracified-nft-backend/businessFacade/marketplaceBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/commons"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonMethods"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/middleware"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/gorilla/mux"
)

func CreateNFT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var nft models.NFT
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&nft)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}

	err = validations.ValidateInsertNft(nft)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err := marketplaceBusinessFacade.StoreNFT(nft)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}

func GetAllNFTs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	results, err1 := marketplaceBusinessFacade.GetAllNFTs()

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

func GetImageBase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["imagebase64"] != "" {
		results, err := marketplaceBusinessFacade.GetImageBase(vars["imagebase64"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[models.NFT](w, results)
		}
	} else {
		errors.BadRequest(w, "")
	}
}

func SaveTXN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var txn models.TXN
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&txn)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}

	err = validations.ValidateRequestTXNObject(txn)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err := marketplaceBusinessFacade.StoreTXN(txn)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}

func CreateOwner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var owner models.Ownership
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&owner)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	_, err1 := marketplaceBusinessFacade.StoreOwner(owner)
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
			commonResponse.SuccessStatus[models.NFT](w, result)
		}
	}
}

func GetAllONSaleNFT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["sellingstatus"] != "" || vars["currentownerpk"] != "" {
		results, err := marketplaceBusinessFacade.GetAllONSaleNFT(vars["sellingstatus"], vars["currentownerpk"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[[]models.NFT](w, results)
		}
	} else {
		errors.BadRequest(w, "")
	}
}

func GetOneONSaleNFT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["sellingstatus"] != "" || vars["nftidentifer"] != "" || vars["blockchain"] != "" {
		results, err := marketplaceBusinessFacade.GetOneONSaleNFT(vars["sellingstatus"], vars["nftidentifier"], vars["blockchain"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[[]models.NFT](w, results)
			return
		}
	} else {
		errors.BadRequest(w, "")
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
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
	var pagination requestDtos.NFTsForMatrixView
	tagToSearch := vars["tag"]
	pgsize, err1 := strconv.Atoi(vars["pagesize"])
	if err1 != nil {
		errors.BadRequest(w, "Requested invalid page size.")
		return
	}
	pagination.PageSize = int32(pgsize)
	requestedPage, err2 := strconv.Atoi(vars["requestedPage"])
	if err2 != nil {
		errors.BadRequest(w, "Requested page error")
		return
	}
	pagination.RequestedPage = int32(requestedPage)
	pagination.SortbyFeild = "timestamp"
	logs.InfoLogger.Println("Received pagination requested: ", pagination)
	results, err := marketplaceBusinessFacade.GEtNFTbyTagsName(pagination, tagToSearch)
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
			errors.BadRequest(w, "Requested tag does not exist")
			return
		}
		commonResponse.SuccessStatus[models.Paginateresponse](w, results)
	}
}

func GetNFTbyStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["sellingstatus"] != "" {
		results, err := marketplaceBusinessFacade.GetNFTBySellingStatus(vars["sellingstatus"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			if results[0].ArtistName == "" {
				errors.BadRequest(w, "No Content")
				return
			}
			commonResponse.SuccessStatus[[]models.NFT](w, results)
			return
		}
	} else {
		errors.BadRequest(w, "invalid status")
	}
}

func GetNFTByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if len(vars["currentownerpk"]) != 0 {
		result, err := marketplaceBusinessFacade.GetNFTbyAccount(vars["currentownerpk"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[[]models.NFT](w, result)
		}
	} else {
		errors.BadRequest(w, "")
	}
}

func GetSVGBySHA256(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if len(vars["hash"]) != 0 {
		result, err := marketplaceBusinessFacade.GetSVGByHash(vars["hash"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[models.SVG](w, result)
		}
	} else {
		errors.BadRequest(w, "")
	}
}

func GetLastNFTByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if len(vars["creatoruserid"]) != 0 {
		result, err := marketplaceBusinessFacade.GetLastNFTbyUserId(vars["creatoruserid"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[[]models.NFT](w, result)
		}
	} else {
		errors.BadRequest(w, "")
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func GetNFTByTenentName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if len(vars["creatoruserid"]) != 0 {
		result, err := marketplaceBusinessFacade.GetNFTbyTenentName((vars["creatoruserid"]))
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[[]models.NFT](w, result)
		}
	} else {
		errors.BadRequest(w, "")
	}
}

func GetNFTByBlockchain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")

	vars := mux.Vars(r)
	if len(vars["blockchain"]) != 0 {
		result, err := marketplaceBusinessFacade.GetNFTbyBlockchain((vars["blockchain"]))
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[[]models.NFT](w, result)
		}
	} else {
		errors.BadRequest(w, "")
	}
}

func GetNFTByBlockchainAndUserPK(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["currentownerpk"] != "" || vars["blockchain"] != "" {
		results, err := marketplaceBusinessFacade.GetNFTByBlockchainAndUserPK(vars["currentownerpk"], vars["blockchain"])
		if err != nil {
			errors.BadRequest(w, err.Error())
			return
		} else {
			if results == nil {
				errors.BadRequest(w, "No Content")
				return
			}
			commonResponse.SuccessStatus[[]models.NFT](w, results)
		}
	} else {
		errors.BadRequest(w, "")
	}
}

func GetTXNByBlockchainAndIdentifier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["nftidentifier"] != "" || vars["blockchain"] != "" {
		results, err := marketplaceBusinessFacade.GetTXNByBlockchainAndIdentifier(vars["nftidentifier"], vars["blockchain"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[[]models.TXN](w, results)
		}
	} else {
		errors.BadRequest(w, "")
	}
}

func CreateTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var tags models.Tags
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tags)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
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
	results, err1 := marketplaceBusinessFacade.GetAllTags()

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

func GetTagsByNFTName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	results, err1 := marketplaceBusinessFacade.GetTagsByNFTName(vars["nftName"])
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

func UpdateMinter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var updateObj requestDtos.UpdateMint
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateObj)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	} else {
		_, err1 := marketplaceBusinessFacade.UpdateNFT(updateObj)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			message := "Minter updated successfully."
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}

func UpdateTXN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var updateObj requestDtos.UpdateMintTXN
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateObj)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	} else {
		_, err1 := marketplaceBusinessFacade.UpdateNFTTXN(updateObj)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			message := "Minter updated successfully."
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}

func SaveNFTStory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var story models.NFTStory
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&story)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}

	err = validations.ValidateInsertNftStory(story)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err := marketplaceBusinessFacade.StoreNFTStory(story)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}

func GetNFTStory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["nftidentifier"] != "" || vars["blockchain"] != "" {
		results, err := marketplaceBusinessFacade.GetNFTStory(vars["nftidentifier"], vars["blockchain"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[[]models.NFTStory](w, results)
		}
	} else {
		errors.BadRequest(w, "")
	}
}

func GetNFTByCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	var pagination requestDtos.NFTsForMatrixView
	pagination.Blockchain = vars["blockchain"]
	var CollectionToSearch = vars["collection"]
	var pubKey = vars["pubkey"]
	pgsize, err1 := strconv.Atoi(vars["pagesize"])
	if err1 != nil {
		errors.BadRequest(w, "Requested invalid page size.")
		return
	}
	pagination.PageSize = int32(pgsize)
	requestedPage, err2 := strconv.Atoi(vars["requestedPage"])
	if err2 != nil {
		errors.BadRequest(w, "Requested page error")
		return
	}
	pagination.RequestedPage = int32(requestedPage)
	pagination.SortbyFeild = "blockchain"
	logs.InfoLogger.Println("Received pagination requested: ", pagination)
	results, err := marketplaceBusinessFacade.GetNFTByCollection(pagination, CollectionToSearch, pubKey)
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
		commonResponse.SuccessStatus[models.Paginateresponse](w, results)
	}
}

/**
 **Description:Retrieves all nfts for the specified blockchain in a paginated format
 **Returns:Object ID of the new OTP created
 */
func GetPaginatedNFTs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	var pagination requestDtos.NFTsForMatrixView
	pagination.Blockchain = vars["blockchain"]
	pgsize, err1 := strconv.Atoi(vars["pagesize"])
	if err1 != nil {
		errors.BadRequest(w, "Requested invalid page size.")
		return
	}
	pagination.PageSize = int32(pgsize)
	requestedPage, err2 := strconv.Atoi(vars["requestedPage"])
	if err2 != nil {
		errors.BadRequest(w, "Requested page error")
	}
	pagination.RequestedPage = int32(requestedPage)
	pagination.SortbyFeild = "blockchain"
	results, err := marketplaceBusinessFacade.GetNFTPagination(pagination)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if pagination.RequestedPage < 0 {
			errors.NotFound(w, "Requested page size should be greater than zero")
			return
		}
		if results.PaginationInfo.TotalPages < pagination.RequestedPage {
			errors.NotFound(w, "requested page does not exist")
			return
		}
		commonResponse.SuccessStatus[models.Paginateresponse](w, results)
	}
}

/**
 **Description: Retrieves nfts that are having the ON SALE status in a paginated format
 **Returns:Paginated NfT Data
 */
func GetPaginatedNFTbySellingStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	var pagination requestDtos.NFTsForMatrixView
	if vars["sellingstatus"] == "ON SALE" {
		pagination.Blockchain = vars["blockchain"]
		pagesize, err1 := strconv.Atoi(vars["pagesize"])
		if err1 != nil {
			errors.BadRequest(w, "Requested invalid page size.")
			return
		}
		pagination.PageSize = int32(pagesize)
		requestedPage, err2 := strconv.Atoi(vars["requestedPage"])
		if err2 != nil {
			errors.BadRequest(w, "Requested page error")
		}
		pagination.RequestedPage = int32(requestedPage)
		pagination.SortbyFeild = vars["sellingstatus"]
		results, err := marketplaceBusinessFacade.GetPaginatedNFTbySellingStatus(pagination)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			if pagination.RequestedPage < 0 {
				errors.NotFound(w, "Requested page size should be greater than zero")
				return
			}
			if results.PaginationInfo.TotalPages < int32(requestedPage) {
				errors.NotFound(w, "requested page does not exist")
				return
			}
			commonResponse.SuccessStatus[models.Paginateresponse](w, results)
		}
	}
}

/**
 **Description:function is used to paginate and return block chain specific nfts which are either trending or under hotpicks
 **Returns:Paginated nft data
 */
func GetPaginatedNFTforstatusFilters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	var pagination requestDtos.NFTsForMatrixView
	pagination.Blockchain = vars["blockchain"]
	pgsize, err1 := strconv.Atoi(vars["pagesize"])
	if err1 != nil {
		errors.BadRequest(w, "Requested invalid page size.")
		return
	}
	pagination.PageSize = int32(pgsize)
	requestedPage, err2 := strconv.Atoi(vars["requestedPage"])
	if err2 != nil {
		errors.BadRequest(w, "Requested page error")
	}

	pagination.RequestedPage = int32(requestedPage)
	if vars["type"] == "hotpicks" {
		pagination.SortbyFeild = "hotpicks"
	} else if vars["type"] == "trending" {
		pagination.SortbyFeild = "trending"
	}
	results, err := marketplaceBusinessFacade.GetPaginatedNFTbyStatusFilter(pagination)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if pagination.RequestedPage < 0 {
			errors.NotFound(w, "Requested page size should be greater than zero")
			return
		}
		if results.PaginationInfo.TotalPages < pagination.RequestedPage {
			errors.NotFound(w, "requested page does not exist")
			return
		}
		commonResponse.SuccessStatus[models.Paginateresponse](w, results)
	}
}

/**
 **Description:function is used to paginate and return block chain specific nfts which are either trending or under hotpicks that are on Sale
 **Returns:Paginated nft data
 */
func GetPaginatedOnSaleNFTforstatusFilters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	var pagination requestDtos.NFTsForMatrixView
	pagination.Blockchain = vars["blockchain"]
	pgsize, err1 := strconv.Atoi(vars["pagesize"])
	if err1 != nil {
		errors.BadRequest(w, "Requested invalid page size.")
		return
	}
	pagination.PageSize = int32(pgsize)
	requestedPage, err2 := strconv.Atoi(vars["requestedPage"])
	if err2 != nil {
		errors.BadRequest(w, "Requested page error")
	}

	pagination.RequestedPage = int32(requestedPage)
	if vars["type"] == "hotpicks" {
		pagination.SortbyFeild = "hotpicks"
	} else if vars["type"] == "trending" {
		pagination.SortbyFeild = "trending"
	}
	results, err := marketplaceBusinessFacade.GetPaginatedOnSaleNFTbyStatusFilter(pagination)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if pagination.RequestedPage < 0 {
			errors.NotFound(w, "Requested page size should be greater than zero")
			return
		}
		if results.PaginationInfo.TotalPages < pagination.RequestedPage {
			errors.NotFound(w, "requested page does not exist")
			return
		}
		commonResponse.SuccessStatus[models.Paginateresponse](w, results)
	}
}

/**
 **Description: Get nfts that are on trending and hotpicks
 **Returns : Paginated nft data
 */
func GetBestCreations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	var bestcreations requestDtos.NFTsForMatrixView
	bestcreations.Blockchain = vars["blockchain"]
	PageSize, pageSizeerr := strconv.Atoi(vars["pagesize"])
	if pageSizeerr != nil {
		errors.BadRequest(w, "Requested page error")
		return
	}
	bestcreations.PageSize = int32(PageSize)
	RequestedPage, pageReqeustederr := strconv.Atoi(vars["requestedPage"])
	if pageReqeustederr != nil {
		errors.BadRequest(w, "Requested page error")
		return
	}
	bestcreations.RequestedPage = int32(RequestedPage)
	results, err := marketplaceBusinessFacade.GetPaginatedResponseforBestCreations(bestcreations)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		if bestcreations.RequestedPage < 0 {
			errors.NotFound(w, "Requested page size should be greater than zero")
			return
		}
		if results.PaginationInfo.TotalPages < bestcreations.RequestedPage {
			errors.NotFound(w, "requested page does not exist")
			return
		}
		commonResponse.SuccessStatus[models.Paginateresponse](w, results)
	}
}

/**
 **Description:Calcluate the avg rating for creators and return creators having a rating > 4
 **Returns : Paginated creator information(name,email,publickey and avg rating)
 */
func GetBestCreators(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	var creatorPaginationInfo requestDtos.CreatorInfoforMatrixView
	result, err := marketplaceBusinessFacade.GetBestCreators()
	if err != nil {
		errors.BadRequest(w, "Failed :"+err.Error())
	}
	_, updateErr := marketplaceBusinessFacade.UpdateBestCreators(result)
	if updateErr != nil {
		errors.BadRequest(w, updateErr.Error())
		return
	}
	PageSize, pageSizeerr := strconv.Atoi(vars["pagesize"])
	if pageSizeerr != nil {
		errors.BadRequest(w, "Requested page error")
		return
	}
	creatorPaginationInfo.PageSize = int32(PageSize)
	RequestedPage, pageReqeustederr := strconv.Atoi(vars["requestedPage"])
	if pageReqeustederr != nil {
		errors.BadRequest(w, "Requested page error")
		return
	}
	creatorPaginationInfo.RequestedPage = int32(RequestedPage)
	res, err := marketplaceBusinessFacade.GetPaginatedBestCreators(creatorPaginationInfo)
	if err != nil {
		errors.BadRequest(w, "failed to get data : "+err.Error())
	}
	if RequestedPage < 0 {
		errors.NotFound(w, "Requested page size should be greater than zero")
		return
	}
	if res.PaginationInfo.TotalPages < int32(RequestedPage) {
		errors.NotFound(w, "requested page does not exist")
		return
	}
	commonResponse.SuccessStatus[models.PaginatedCreatorInfo](w, res)
}

func GetImagebyID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if len(vars["id"]) != 0 {
		result, err := marketplaceBusinessFacade.GetThumbnailbyID(vars["id"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[models.ThumbNail](w, result)
		}
	} else {
		errors.BadRequest(w, "")
	}
}

func GetProfileContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	var pagination requestDtos.NFTsForMatrixView
	pagination.SortbyFeild = vars["pubkey"]
	pagination.Blockchain = vars["blockchain"]
	filterBy := vars["filter"]
	pgsize, err1 := strconv.Atoi(vars["pagesize"])
	if err1 != nil {
		errors.BadRequest(w, "Requested invalid page size.")
		return
	}
	pagination.PageSize = int32(pgsize)
	requestedPage, err2 := strconv.Atoi(vars["requestedPage"])
	if err2 != nil {
		errors.BadRequest(w, "Requested page error")
		return
	}
	pagination.RequestedPage = int32(requestedPage)

	logs.InfoLogger.Println("Received pagination requested: ", pagination)
	results, err := marketplaceBusinessFacade.GetUserProfileContent(pagination, filterBy)
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
			errors.BadRequest(w, "No Content for account: "+vars["pubkey"])
			return
		}
		commonResponse.SuccessStatus[models.Paginateresponse](w, results)
	}
}

func SaveNFTFromWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var wnft models.WalletNFT
	ps := middleware.WalletUserHasPermissionToMint(r.Header.Get("Authorization"))
	if !ps.Status {
		w.WriteHeader(http.StatusUnauthorized)
		logs.ErrorLogger.Println("Status Unauthorized")
		return
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&wnft)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	//Creator is initial owner of NFT
	wnft.NFTOwner = wnft.NFTCreator
	OTPencode := commonMethods.StringToSHA256(wnft.OTP)
	wnft.OTP = OTPencode
	err = validations.ValidateWalletNft(wnft)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err := marketplaceBusinessFacade.StoreWalletNFT(wnft)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}

func GetAllWalletNFTs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	results, err1 := marketplaceBusinessFacade.GetAllWalletNFTs()

	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		if len(results) > 0 {
			for i, result := range results {
				thumbnailUrl := "https://tracified.sirv.com/Spins/RURI%20Gems/" + result.ShopID + "/" + result.ShopID + ".jpg"
				results[i].Thumbnail = thumbnailUrl
			}
		} else {
			ErrorMessage := "No NFTs found"
			errors.BadRequest(w, ErrorMessage)
			return
		}
		commonResponse.SuccessStatus(w, results)
		return
	}
}

func GetWalletNFTsbyPK(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	results, err1 := marketplaceBusinessFacade.GetAllWalletNFTsbyPK(vars["publickey"])
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {

		commonResponse.SuccessStatus(w, results)
		return
	}
}

func GetNFTByBlockchainAndIdentifier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["nftidentifier"] != "" || vars["blockchain"] != "" {
		results, err := marketplaceBusinessFacade.GetNFTByBlockchainAndIdentifier(vars["nftidentifier"], vars["blockchain"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[models.NFT](w, results)
		}
	} else {
		errors.BadRequest(w, "")
	}
}

func SaveContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var contracts models.ContractInfo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&contracts)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}

	err = validations.ValidateInsertContract(contracts)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err := marketplaceBusinessFacade.StoreContracts(contracts)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}

func GetContractByUserAndBC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")

	vars := mux.Vars(r)
	if len(vars["user"]) != 0 && len(vars["blockchain"]) != 0 {
		result, err := marketplaceBusinessFacade.GetContractbyBlockchainAndUser((vars["blockchain"]), (vars["user"]))
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[[]models.ContractInfo](w, result)
		}
	} else {
		errors.BadRequest(w, "")
	}
}

func UpdateWalletNFTOwner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	var updaterequest requestDtos.WalletNFTUpdateOwner
	decoder := json.NewDecoder(r.Body)
	decodeErr := decoder.Decode(&updaterequest)
	if decodeErr != nil {
		errors.BadRequest(w, decodeErr.Error())
		return
	}
	rst, err := marketplaceBusinessFacade.UpdateWalletNFTOwner(updaterequest)
	if err != nil {
		errors.BadRequest(w, decodeErr.Error())
		return
	}
	if rst.ID.String() == "" {
		errors.BadRequest(w, "Record does not exist")
		return
	}
	commonResponse.SuccessStatus[string](w, rst.ID.Hex())
}

// GetMintedNFTInfoTenant handles HTTP requests to retrieve minted NFT information
// for a specific tenantID and sends a JSON response.
func GetMintedNFTInfoTenant(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json;")

	// Get the tenantID from request parameters
	vars := mux.Vars(r)
	tenantID := vars["tenantID"]

	// Check if the tenantID is valid using the validatetenant function
	if _validatetenant(tenantID) {
		// Call a function to get minted NFT identifier for the given tenantID
		rst, err := customizedNFTFacade.GetMintedNFTIdentifierForWallet(vars["tenantID"])
		if err != nil {
			// Handle error and send a bad request response if there's an issue with fetching data
			logs.InfoLogger.Printf("Error fetching data for tenantID %s: %v", vars["tenantID"], err)
			errors.InternalError(w, "failed to get data please try again.")
			return
		}
		if len(rst) == 0 {
			commonResponse.SuccessStatus[string](w, "no minted NFTs")
			return
		}
		// Send a successful response with the result using the commonResponse.SuccessStatus function
		commonResponse.SuccessStatus[[]string](w, rst)
	} else {
		// Send a bad request response for an invalid tenantID
		errors.BadRequest(w, "Invalid TenantID")
	}
}

// validatetenant checks if the provided tenantID is valid by comparing it
// against a list of valid tenantIDs obtained from an environment variable.
func _validatetenant(tenantID string) bool {
	// Get the list of valid tenant IDs from an environment variable
	list := commons.GoDotEnvVariable("MINTED_WALLETNFT_TENANTS")

	// Split the comma-separated list into individual tenantIDs
	tenantIDs := strings.Split(list, ",")

	// Check if the provided tenantID exists in the list
	for _, i := range tenantIDs {
		if tenantID == i {
			// TenantID is valid
			return true
		}
	}

	// TenantID is not valid
	return false
}
