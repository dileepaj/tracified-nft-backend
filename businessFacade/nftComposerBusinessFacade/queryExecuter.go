package nftComposerBusinessFacade

import (
	"net/http"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"

	"github.com/dileepaj/tracified-nft-backend/services/otpService"
	"github.com/dileepaj/tracified-nft-backend/services/queryLanguageService"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
)

func QueryExecuter(w http.ResponseWriter, widget requestDtos.RequestWidget, token string) {
	var queryResult responseDtos.QueryResult
	result, err := FindWidgetByWidgetIdWithOTP(widget.WidgetId)
	if err != nil {
		commonResponse.NoContent(w, err.Error())
	}
	if result.OTPType == "Batch" || result.OTPType == "Artifact" {
		if strings.HasPrefix(result.OTP, `{"url":`) {
			OtpJson, err := otpService.GetOtpJSON(result.OTP, token)
			if err != nil {
				commonResponse.RespondWithJSON(w, http.StatusInternalServerError, "Invalid Json fromat")
			} else {
				queryResult = queryLanguageService.QueryExecuter(widget.Query, result,OtpJson )

				if strings.HasPrefix(queryResult.Result, "\nError") {
					commonResponse.SuccessStatus(w, "invalid Query")
				} else {
					_, err1 := FindWidgetAndUpdateQuery(widget)
					if err1 != nil {
						commonResponse.NoContent(w, err1.Error())
					} else {
						commonResponse.SuccessStatus(w, queryResult)
					}
				}
			}
		} else {
			commonResponse.RespondWithJSON(w, http.StatusInternalServerError, "Invalid tracibility Data")
		}
	} else {
		commonResponse.RespondWithJSON(w, http.StatusInternalServerError, "Invalid tracibility Data type")
	}
}
