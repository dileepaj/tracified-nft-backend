package apiHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	ipfsbusinessfacade "github.com/dileepaj/tracified-nft-backend/businessFacade/ipfsBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/businessFacade/marketplaceBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/commons"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/gorilla/mux"
)

// CreateCollection handles creating a new collection.
// It decodes the request body into a CreateCollectionObject,
// validates the data, checks if the collection name is available,
// uploads the collection to IPFS, and returns the IPFS CID
// or an error response.
func CreateCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var createCollectionObject models.IpfsObjectForCollections
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createCollectionObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		errors.BadRequest(w, "Invalid visibility setting.")
		return
	}
	err = validations.ValidateInsertCollection(createCollectionObject.CollectionDetails)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		availableRst, err := marketplaceBusinessFacade.IsCollectionNameTaken(createCollectionObject.CollectionDetails.CollectionName)
		if err != nil {
			logs.ErrorLogger.Println("Error while checking if collection name is taken: ", err.Error())
			errors.InternalError(w, "Somthing Went Wrong!")
			return
		}
		if availableRst {
			logs.ErrorLogger.Println("Collection name " + createCollectionObject.CollectionDetails.CollectionName + " already taken!")
			errors.BadRequest(w, "Collection name "+createCollectionObject.CollectionDetails.CollectionName+" already taken!")
			return
		}
		cid, ipfserr := ipfsbusinessfacade.UploadCollectionsToIpfs(createCollectionObject)
		if ipfserr != nil {
			ErrorMessage := ipfserr.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			commonResponse.SuccessStatus[string](w, cid)
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

/**
** /collection/{publickey}?limit=10?page=1?sort=-1/1
 */
func GetCollectionByPublicKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	var paginationRequest requestDtos.CollectionPagination

	pgsize, err1 := strconv.Atoi(r.URL.Query().Get("limit"))
	if err1 != nil || pgsize <= 0 {
		_pgsize, envErr := strconv.Atoi(commons.GoDotEnvVariable("PAGINATION_DEFUALT_LIMIT"))
		if envErr != nil {
			errors.InternalError(w, "Something went wrong")
			logs.ErrorLogger.Println("Failed to load value from env: ", envErr.Error())
			return
		}
		pgsize = _pgsize
	}
	paginationRequest.PageSize = int32(pgsize)

	requestedPage, err2 := strconv.Atoi(r.URL.Query().Get("page"))
	if err2 != nil || requestedPage <= -1 {
		_requestedPage, envErr := strconv.Atoi(commons.GoDotEnvVariable("PAGINATION_DEFAULT_PAGE"))
		if envErr != nil {
			errors.InternalError(w, "Something went wrong")
			logs.ErrorLogger.Println("Failed to load value from env: ", envErr.Error())
			return
		}
		requestedPage = _requestedPage
	}
	paginationRequest.RequestedPage = int32(requestedPage)

	sort, err := strconv.Atoi(r.URL.Query().Get("sort"))
	if err != nil || sort != -1 && sort != 1 {
		_sort, envErr := strconv.Atoi(commons.GoDotEnvVariable("PAGINATION_DEFAULT_SORT"))
		if envErr != nil {
			errors.InternalError(w, "Something went wrong")
			logs.ErrorLogger.Println("Failed to load value from env: ", envErr.Error())
			return
		}
		sort = _sort
	}
	paginationRequest.SortType = sort
	results, err1 := marketplaceBusinessFacade.GetCollectionByPublicKeyPaginated(paginationRequest, vars["userid"])
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		if results.Content == nil {
			errorMessage := "User has no collections"
			errors.BadRequest(w, errorMessage)
			return
		}
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}

}

/**
** /collection?limit=10?page=1?sort=-1/1
 */
func GetAllCollections(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	log.Println("calling func Get All Collections....")
	var paginationRequest requestDtos.CollectionPagination
	pgsize, err1 := strconv.Atoi(r.URL.Query().Get("limit"))
	if err1 != nil || pgsize <= 0 {
		_pgsize, envErr := strconv.Atoi(commons.GoDotEnvVariable("PAGINATION_DEFUALT_LIMIT"))
		if envErr != nil {
			errors.InternalError(w, "Something went wrong")
			logs.ErrorLogger.Println("Failed to load value from env: ", envErr.Error())
			return
		}
		pgsize = _pgsize
	}
	paginationRequest.PageSize = int32(pgsize)

	requestedPage, err2 := strconv.Atoi(r.URL.Query().Get("page"))
	if err2 != nil || requestedPage <= -1 {
		_requestedPage, envErr := strconv.Atoi(commons.GoDotEnvVariable("PAGINATION_DEFAULT_PAGE"))
		if envErr != nil {
			errors.InternalError(w, "Something went wrong")
			logs.ErrorLogger.Println("Failed to load value from env: ", envErr.Error())
			return
		}
		requestedPage = _requestedPage
	}
	paginationRequest.RequestedPage = int32(requestedPage)

	sort, err := strconv.Atoi(r.URL.Query().Get("sort"))
	if err != nil || sort != -1 && sort != 1 {
		_sort, envErr := strconv.Atoi(commons.GoDotEnvVariable("PAGINATION_DEFAULT_SORT"))
		if envErr != nil {
			errors.InternalError(w, "Something went wrong")
			logs.ErrorLogger.Println("Failed to load value from env: ", envErr.Error())
			return
		}
		sort = _sort
	}
	paginationRequest.SortType = sort

	results, err1 := marketplaceBusinessFacade.GetAllCollectionsPaginated(paginationRequest)
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		if results.Content == nil {
			errorMessage := "No collections"
			errors.BadRequest(w, errorMessage)
			return
		}
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

func UpdateCollectionVisibility(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var UpdateObject requestDtos.UpdateCollectionVisibility
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&UpdateObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		errors.BadRequest(w, err.Error())
		return
	} else {
		result, err := marketplaceBusinessFacade.UpdateCollectionVisibility(UpdateObject)
		if err != nil {
			logs.WarningLogger.Println("Failed to update visibility of collection : ", err.Error())
			errors.BadRequest(w, err.Error())
			return
		}
		commonResponse.SuccessStatus[models.NFTCollection](w, result)
	}
}
