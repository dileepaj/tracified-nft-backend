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
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		_, objIDerr := primitive.ObjectIDFromHex(nft.NFTID)
		if objIDerr != nil {
			errors.BadRequest(w, "Invalid NFT ID : "+objIDerr.Error())
			return
		}
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

// GetWalletNFTByStateAndCurrentOwner retrieves wallet NFTs based on their state and current owner.
// This function handles HTTP requests and responses, populating the response with the appropriate NFTs.
// It expects the following query parameters:
//   - blockchain: The blockchain to filter the NFTs.
//   - nftstatus: The state of the NFTs (optional). If not provided, all states will be considered.
//   - currentowner: The current owner's public key for filtering NFTs.
//   - pagesize: The number of NFTs to display per page.
//   - requestedPage: The requested page number for pagination.
func GetWalletNFTByStateAndCurrentOwner(w http.ResponseWriter, r *http.Request) {
	// Variables to store the state and its conversion error, if any.
	var StateToSearch int
	var stateErr error

	// Set the response header to indicate JSON content.
	w.Header().Set("Content-Type", "application/json;")

	// Structure to hold pagination information.
	var pagination requestDtos.WalletNFTsForMatrixView

	// Extract blockchain and NFT status from the query parameters.
	pagination.Blockchain = r.URL.Query().Get("blockchain")
	state := r.URL.Query().Get("nftstatus")

	// Convert the NFT status to an integer.
	if state != "" {
		StateToSearch, stateErr = strconv.Atoi(state)
		if stateErr != nil {
			// Handle invalid NFT state.
			errors.BadRequest(w, "Invalid NFT state")
			return
		}
	} else {
		StateToSearch = 0 // If NFT status is not provided, consider all states.
	}

	// Extract the current owner's public key from the query parameters.
	var pubKey = r.URL.Query().Get("currentowner")

	// Extract the page size from the query parameters and convert it to an integer.
	pgsize, err1 := strconv.Atoi(r.URL.Query().Get("pagesize"))
	if err1 != nil {
		// Handle invalid page size.
		errors.BadRequest(w, "Requested invalid page size.")
		return
	}

	// Assign the page size to the pagination structure.
	pagination.PageSize = int32(pgsize)

	// Extract the requested page number from the query parameters and convert it to an integer.
	requestedPage, err2 := strconv.Atoi(r.URL.Query().Get("requestedPage"))
	if err2 != nil {
		// Handle requested page error.
		errors.BadRequest(w, "Requested page error")
		return
	}

	// Assign the requested page number to the pagination structure.
	pagination.RequestedPage = int32(requestedPage)

	// Set the default sorting field to "blockchain".
	pagination.SortbyFeild = "blockchain"

	// Log the received pagination request.
	logs.InfoLogger.Println("Received pagination requested: ", pagination)

	// Call the business facade to get the wallet NFTs based on state and current owner.
	results, err := marketplaceBusinessFacade.GetWalletNFTByState(pagination, StateToSearch, pubKey)
	if err != nil {
		// Handle error by sending a bad request with the error message.
		errors.BadRequest(w, err.Error())
	} else {
		// Check for valid pagination parameters and NFT results.

		if pagination.PageSize <= 0 {
			// Handle invalid page size.
			errors.BadRequest(w, "Page size should be greater than zero")
			return
		}

		if pagination.RequestedPage < 0 {
			// Handle invalid requested page size.
			errors.BadRequest(w, "Requested page size should be greater than zero")
			return
		}

		if results.PaginationInfo.TotalPages < pagination.RequestedPage {
			// Handle non-existing requested page.
			errors.BadRequest(w, "requested page does not exist")
			return
		}

		if results.Content == nil {
			// Handle the case when there are no NFTs in the collection.
			errors.BadRequest(w, "Collection does not have any NFTs")
			return
		}

		// If everything is valid, send a successful response with the NFTs.
		commonResponse.SuccessStatus[models.PaginateWalletNFTResponse](w, results)
	}
}

// GetWalletNFTByStatusAndRequested retrieves wallet NFTs based on their status and request.
// This function handles HTTP requests and responses, populating the response with the appropriate NFTs.
// It expects the following query parameters:
//   - blockchain: The blockchain to filter the NFTs.
//   - nftstatus: The status of the NFTs (optional). If not provided, all statuses will be considered.
//   - nftrequested: The public key of the user who has requested the NFT for filtering NFTs.
//   - pagesize: The number of NFTs to display per page.
//   - requestedPage: The requested page number for pagination.
func GetWalletNFTByStatusAndRequested(w http.ResponseWriter, r *http.Request) {
	// Variables to store the state and its conversion error, if any.
	var StateToSearch int
	var stateErr error

	// Set the response header to indicate JSON content.
	w.Header().Set("Content-Type", "application/json;")

	// Structure to hold pagination information.
	var pagination requestDtos.WalletNFTsForMatrixView

	// Extract blockchain and NFT status from the query parameters.
	pagination.Blockchain = r.URL.Query().Get("blockchain")
	state := r.URL.Query().Get("nftstatus")

	// Convert the NFT status to an integer.
	if state != "" {
		StateToSearch, stateErr = strconv.Atoi(state)
		if stateErr != nil {
			// Handle invalid NFT status.
			errors.BadRequest(w, "Invalid State.")
			return
		}
	} else {
		StateToSearch = 0 // If NFT status is not provided, consider all statuses.
	}

	// Extract the requested public key from the query parameters.
	var pubKey = r.URL.Query().Get("nftrequested")

	// Extract the page size from the query parameters and convert it to an integer.
	pgsize, err1 := strconv.Atoi(r.URL.Query().Get("pagesize"))
	if err1 != nil {
		// Handle invalid page size.
		errors.BadRequest(w, "Requested invalid page size.")
		return
	}

	// Assign the page size to the pagination structure.
	pagination.PageSize = int32(pgsize)

	// Extract the requested page number from the query parameters and convert it to an integer.
	requestedPage, err2 := strconv.Atoi(r.URL.Query().Get("requestedPage"))
	if err2 != nil {
		// Handle requested page error.
		errors.BadRequest(w, "Requested page error")
		return
	}

	// Assign the requested page number to the pagination structure.
	pagination.RequestedPage = int32(requestedPage)

	// Set the default sorting field to "blockchain".
	pagination.SortbyFeild = "blockchain"

	// Log the received pagination request.
	logs.InfoLogger.Println("Received pagination requested: ", pagination)

	// Call the business facade to get the wallet NFTs based on status and requested public key.
	results, err := marketplaceBusinessFacade.GetWalletNFTByStateForRequested(pagination, StateToSearch, pubKey)
	if err != nil {
		// Handle error by sending a bad request with the error message.
		errors.BadRequest(w, err.Error())
	} else {
		// Check for valid pagination parameters and NFT results.

		if pagination.PageSize <= 0 {
			// Handle invalid page size.
			errors.BadRequest(w, "Page size should be greater than zero")
			return
		}

		if pagination.RequestedPage < 0 {
			// Handle invalid requested page size.
			errors.BadRequest(w, "Requested page size should be greater than zero")
			return
		}

		if results.PaginationInfo.TotalPages < pagination.RequestedPage {
			// Handle non-existing requested page.
			errors.BadRequest(w, "requested page does not exist")
			return
		}

		if results.Content == nil {
			// Handle the case when there are no NFTs in the collection.
			errors.BadRequest(w, "Collection does not have any NFTs")
			return
		}

		// If everything is valid, send a successful response with the NFTs.
		commonResponse.SuccessStatus[models.PaginateWalletNFTResponse](w, results)
	}
}
