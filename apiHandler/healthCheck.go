package apiHandler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/gorilla/context"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp := responseDtos.HealthCheckResponse{
		Note:    "Tracified nft backend up and running",
		Time:    time.Now().Format("Mon Jan _2 15:04:05 2006"),
		Version: "0",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encodeErr := json.NewEncoder(w).Encode(resp)
	if encodeErr != nil {
		logs.ErrorLogger.Println("failed to encode JSON: ", encodeErr.Error())
	}

}
