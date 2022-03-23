package nftComposerBusinessFacade

import (
	"net/http"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	querylanguageservice "github.com/dileepaj/tracified-nft-backend/services/queryLanguageService"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
)

func QueryExecuter(w http.ResponseWriter, widget requestDtos.RequestWidget) {
	result, err := FindWidgetByWidgetId(widget.WidgetId)
	if err != nil {
		commonResponse.NoContent(w, err.Error())
	} else {
		queryResult := querylanguageservice.QueryExecuter(widget.Query, result)
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
}
