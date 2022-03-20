package nftComposerBusinessFacade

import (
	"net/http"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	querylanguageservice "github.com/dileepaj/tracified-nft-backend/services/queryLanguageService"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
)

func QueryExecuter(w http.ResponseWriter, widget requestDtos.RequestWidget) {
	result, err := FindWidgetById(widget.WidgetId)
	if err != nil {
		errors.NoContent(w, err.Error())
	} else {
		queryResult := querylanguageservice.QueryExecuter(widget.Query, result)
		if strings.HasPrefix(queryResult.Result, "\nError") {
			commonResponse.SuccessStatus(w, "invalid Query")
		} else {
			_, err1 := FindWidgetAndUpdateQuery(widget)
			if err1 != nil {
				errors.NoContent(w, err1.Error())
				return
			} else {
				commonResponse.SuccessStatus[responseDtos.QueryResult](w, queryResult)
				return
			}
		}
	}
}
