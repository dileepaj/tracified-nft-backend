package apiHandler

import (
	"encoding/json"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/nftComposerBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
)

// Save the widget data in a DB
func SaveWidget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	var createWidgetResponse models.Widget
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createWidgetResponse)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateWidget(createWidgetResponse)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err := nftComposerBusinessFacade.SaveWidget(createWidgetResponse)
		if err != "" {
			errors.BadRequest(w, err)
		} else {
			commonResponse.RespondWithJSON(w, http.StatusOK, result)
		}
	}
}

func UpdateWidget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	var updateWidgetResponse requestDtos.UpdateWidgetRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateWidgetResponse)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateUpdateWidget(updateWidgetResponse)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err := nftComposerBusinessFacade.ChangeWidget(updateWidgetResponse)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			commonResponse.SuccessStatus(w, result)
		}
	}
}
