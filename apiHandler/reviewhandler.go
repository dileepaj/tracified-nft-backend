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

/*
These functions decode the json sent via api and passes data to the review businessfacade class
*/
func CreateReview(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var requestCreateReview models.Review
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestCreateReview)
	if err != nil {
		logs.ErrorLogger.Println("Error occured while decoding JSON in CreateReview(reviewHandler): ", err.Error())
	}
	err = validations.ValidatReview(requestCreateReview)
	if err != nil {
		errors.BadRequest(W, err.Error())
	} else {
		result, err1 := marketplaceBusinessFacade.CreateReview(requestCreateReview)
		if err1 != nil {
			errors.BadRequest(W, err.Error())
		} else {
			commonResponse.SuccessStatus[string](W, result)
		}
	}
}

func GetNFTReviewByNFT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	results, err1 := marketplaceBusinessFacade.GetReviewByNFT(vars["nftidentifier"])
	if err1 != nil {
		errors.BadRequest(w, err1.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while encoding JSON in GetNFTReviewByNFT(reviewHandler): ", err.Error())
		}
	}
}

func GetAllReviews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	results, err := marketplaceBusinessFacade.GetAllReviews()
	if err != nil {
		ErrorMessage := err.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while encoding JSON in GetAllReviews(reviewHandler): ", err.Error())
		}
		return
	}
}

func UpdateReviewStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var updateReviewStatus requestDtos.UpdateReviewStatus
	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode((&updateReviewStatus))
	if err != nil {
		logs.ErrorLogger.Println("Error occured while decoding JSON in UpdateReviewStatus(reviewHandler):", err.Error())
	} else {
		_, err := marketplaceBusinessFacade.UpdateReviewStatus(updateReviewStatus)
		if err != nil {
			Errormsg := err.Error()
			errors.BadRequest(w, Errormsg)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			message := "Review Status has been Updated"
			err := json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println("Error occured while encoding JSON in UpdateReviewStatus(reviewHandler):", err.Error())
			}
			return
		}
	}
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var deleteReviewObj requestDtos.DeleteReview
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&deleteReviewObj)
	if err != nil {
		logs.ErrorLogger.Println("Error occured while decoding JSON in DeleteReview(reviewHandler):", err.Error())
	} else {
		err1 := marketplaceBusinessFacade.DeleteReview(deleteReviewObj)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			message := "review has been deleted"
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println("Error occured while encode JSON in UpdateReviewStatus(reviewHandler):", err.Error())
			}
			return
		}
	}
}

func GetReviewbyFilter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	var reviewfilter requestDtos.ReviewFiltering
	if vars["filtertype"] == "high" {
		reviewfilter.Filterby = "rating"
		reviewfilter.FilterType = -1

	} else if vars["filtertype"] == "low" {
		reviewfilter.Filterby = "rating"
		reviewfilter.FilterType = 1

	} else if vars["filtertype"] == "latest" {
		reviewfilter.Filterby = "_id"
		reviewfilter.FilterType = -1
	}
	pgsize, err1 := strconv.Atoi(vars["pagesize"])
	if err1 != nil {
		errors.BadRequest(w, "invalid page size.")
		return
	}
	reviewfilter.PageSize = int32(pgsize)
	requestedPage, err2 := strconv.Atoi(vars["pageno"])
	if err2 != nil {
		errors.BadRequest(w, "Requested page error: "+err2.Error())
		return
	}
	if requestedPage < 0 {
		requestedPage = 0
	}
	reviewfilter.RequestedPage = int32(requestedPage)
	if vars["id"] == "" {
		errors.BadRequest(w, "nft identifier missing.")
		return
	}
	reviewfilter.NFTIdentifier = vars["id"]
	result, paginateErr := marketplaceBusinessFacade.GetReviewsbyFilter(reviewfilter)
	if paginateErr != nil {
		errors.BadRequest(w, paginateErr.Error())
		return
	}
	if reviewfilter.RequestedPage > result.PaginationInfo.TotalPages || reviewfilter.RequestedPage < 0 {
		errors.NotFound(w, "Requested page not found")
		return
	}
	commonResponse.SuccessStatus[models.ReviewPaginatedResponse](w, result)
}
