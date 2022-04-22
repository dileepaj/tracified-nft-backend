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

/*
	function QueryExecuter() This fucntion takes widget details that user need to query and Auth token return the query result
	Process
		First find the widget from DB(widgets Documents) and assign it in to result and err(error), if widges does not exist give proper response
		Checked the OTPType. it should be "Batch" or "Artifact"
		Checked the OTP. it should be strat with `{"url":`
		call the GetOtpJSON
		Call the QueryExecuter
		if query result was corret update the widtes's quer field in the document
	?parameters
		w http.ResponseWriter
		widget requestDtos.RequestWidget its incle the query and widget Id
	?return
		Void

*/

func QueryExecuter(w http.ResponseWriter, widget requestDtos.RequestWidget, token string) {
	var queryResult responseDtos.QueryResult
	result, err := FindWidgetByWidgetIdWithOTP(widget.WidgetId)
	if err != nil {
		commonResponse.NoContent(w, err.Error())
	}
	if result.OTPType == "Batch" || result.OTPType == "Artifact" {
		if strings.HasPrefix(result.OTP, `{"url":`) {
			// To retrive the OTp from S3 bucket assign that result to OtpJson, err(error)
			OtpJson, err := otpService.GetOtpJSON(result.OTP, token)
			if err != nil {
				commonResponse.RespondWithJSON(w, http.StatusInternalServerError, "Invalid Json fromat")
			} else {
				// To retrive the query result for widget's OTPs
				queryResult = queryLanguageService.QueryExecuter(widget.Query, result, OtpJson)

				// checked whether result include error in query result
				if strings.HasPrefix(queryResult.Result, "\nError") {
					commonResponse.SuccessStatus(w, "invalid Query")
				} else {
					//update the widget's qury field
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
