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
	logs.InfoLogger.Println("data retreived for sale : ", makeSaleRequestObject)
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
			logs.InfoLogger.Println("Data to be sent back : ", results)
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

func GetNFTbyStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["sellingstatus"] != "" {
		results, err := marketplaceBusinessFacade.GetNFTBySellingStatus(vars["sellingstatus"])
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
		} else {
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
	if vars["collection"] != "" {
		results, err := marketplaceBusinessFacade.GetNFTByCollection(vars["collection"])
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus[[]models.NFT](w, results)
		}
	} else {
		errors.BadRequest(w, "")
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
	logs.InfoLogger.Println("Received pagination requested: ", pagination)
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
	logs.InfoLogger.Println("END Pagination")
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
