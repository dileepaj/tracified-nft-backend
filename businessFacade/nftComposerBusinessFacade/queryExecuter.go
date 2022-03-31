package nftComposerBusinessFacade

import (
	"net/http"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/services/otpService"
	querylanguageservice "github.com/dileepaj/tracified-nft-backend/services/queryLanguageService"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
)

func QueryExecuter(w http.ResponseWriter, widget requestDtos.RequestWidget) {
	var queryResult responseDtos.QueryResult
	result, err := FindWidgetByWidgetIdWithOTP(widget.WidgetId)
	if err != nil {
		commonResponse.NoContent(w, err.Error())
	}
	if result.OTPType == "Batch" {
		if result.OTP == "" {
			commonResponse.RespondWithJSON(w, http.StatusNoContent, "Invalid id")
		} else if strings.HasPrefix(result.OTP, "[{") {
			queryResult = querylanguageservice.QueryExecuter(widget.Query, result)
		} else {
			commonResponse.RespondWithJSON(w, http.StatusInternalServerError, "Invalid tracibility Data fromat")
		}
	} else if result.OTPType == "Artifact" {
		artifactOTP, err := otpService.GetOtpForArtifact(result.ArtifactId, result.OTPType)
		if err != nil {
			commonResponse.NoContent(w, err.Error())
		} else if artifactOTP == "" {
			commonResponse.RespondWithJSON(w, http.StatusNoContent, "Invalid artifact id")
		} else if strings.HasPrefix(artifactOTP, "[{") {
			result.OTP = artifactOTP
			queryResult = querylanguageservice.QueryExecuter(widget.Query, result)
		} else {
			commonResponse.RespondWithJSON(w, http.StatusInternalServerError, "Invalid tracibility Data fromat")
		}

	} else {
		commonResponse.RespondWithJSON(w, http.StatusInternalServerError, "Invalid tracibility Data type")
	}
	if strings.HasPrefix(queryResult.Result, "\nError") {
		commonResponse.SuccessStatus(w, "invalid Query")
	} else {
		_, err1 := FindWidgetAndUpdateQuery(widget)
		if err1 != nil {
			commonResponse.NoContent(w, err1.Error())
			return
		} else {
			commonResponse.SuccessStatus(w, queryResult)
			return
		}
	}
}
