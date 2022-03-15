package apiHandler

import (
	"encoding/json"
	"net/http"
	"strings"

	nftcomposercontroller "github.com/dileepaj/tracified-nft-backend/controllers/nftComposerController"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	querylanguageservice "github.com/dileepaj/tracified-nft-backend/services/queryLanguageService"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
)

//find weiegt using user id and projectid extract OTP from it and execute the quey for it
func FindOtpAndExecuteQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var requestWidget requestDtos.RequestWidget
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestWidget)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
	err = validations.ValidateQueryExecuter(requestWidget)
	if err != nil {
		errors.BadRequest(w, err.Error())
		return
	} else {
		//find the widget by Id and update query  filed
		result, err1 := nftcomposercontroller.FindWidgetById(requestWidget.Id)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.NoContent(w, ErrorMessage)
			return
		} else {
			//pass the widget model to querylanguage service get result
			queryResult := querylanguageservice.QueryExecuter(result)
			if strings.HasPrefix(queryResult.Result, "\nError") {
				errors.NoContent(w, "invalid Query")
			} else {
				_, err1 := nftcomposercontroller.FindWidgetAndUpdateQuery(requestWidget)
				if err1 != nil {
					errors.NoContent(w, err1.Error())
					return
				} else {
					w.WriteHeader(http.StatusOK)
					err = json.NewEncoder(w).Encode(queryResult)
					if err != nil {
						logs.ErrorLogger.Println(err)
					}
					return
				}
			}
		}
	}
}
