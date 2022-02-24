package errors

import (
	"encoding/json"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/wrappers/responseWrappers"
)

func BadRequest(w http.ResponseWriter,message string){
    w.WriteHeader(http.StatusBadRequest)
	 response:=responseWrappers.ErrorResponse{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad request",
	}
	var err=json.NewEncoder(w).Encode(response)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
}

func NotFound(w http.ResponseWriter,message string){
    w.WriteHeader(http.StatusNotFound)
	response:=responseWrappers.ErrorResponse{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not found",
	}
	var err=json.NewEncoder(w).Encode(response)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
}

func InternalError(w http.ResponseWriter,message string){
    w.WriteHeader(http.StatusInternalServerError)
	response:=responseWrappers.ErrorResponse{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal server error",
	}
	var err=json.NewEncoder(w).Encode(response)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
}