package commonResponse

import (
	"encoding/json"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

type resultType interface {
	string | []string | responseDtos.ResponseNFTMakeSale | []models.NFT|responseDtos.QueryResult| []responseDtos.ResponseProject
}

func SuccessStatus[T resultType](w http.ResponseWriter, result T) {
	w.WriteHeader(http.StatusOK)
	response := responseDtos.ResultResponse{
		Status: http.StatusOK,
		Response: result,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
}
