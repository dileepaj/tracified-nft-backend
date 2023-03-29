package commonResponse

import (
	"encoding/json"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

type resultType interface {
	models.Timeline | models.Table | models.Chart | models.StataArray | models.ProofBotData | models.ImageData | models.NFTComposerProject | responseDtos.UpdareProjectResponse | []models.NFTComposerProject | responseDtos.WidgetIdResponse | string | []string | responseDtos.ResponseNFTMakeSale | []models.NFT | responseDtos.QueryResult | []responseDtos.ResponseProject | models.Widget | models.ProjectDetail | models.NewsLetter | models.Faq | []models.SVG | []models.WatchList | []models.Tags | models.NFT | models.SVG | []models.Endorse | []models.Favourite | models.Partner | models.Document | []models.TXN | responseDtos.SVGforNFTResponse | []models.NFTStory | []models.UserQuestions | []responseDtos.GetPendingUserFAQ | models.Paginateresponse | models.PaginatedCreatorInfo | models.ReviewPaginatedResponse | models.ThumbNail | []models.ContractInfo | models.Response | models.SVGforNFTResponse
}

func SuccessStatus[T resultType](w http.ResponseWriter, result T) {
	w.WriteHeader(http.StatusOK)
	response := responseDtos.ResultResponse{
		Status:   http.StatusOK,
		Response: result,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
}

func NoContent(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNoContent)
	response := responseDtos.ErrorResponse{
		Message: message,
		Status:  http.StatusNoContent,
		Error:   "No documents in result",
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
}

func RespondWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}
