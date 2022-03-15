package apiHandler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	resp := responseDtos.HealthCheckResponse{
		Note:    "Tracified nft backend up and running",
		Time:    time.Now().Format("Mon Jan _2 15:04:05 2006"),
		Version: "0",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
