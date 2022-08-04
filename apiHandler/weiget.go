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
	"github.com/dileepaj/tracified-nft-backend/utilities/middleware"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

// Save the widget data in a DB
func SaveWidget(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	
	if ps.Status {
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
			result,statsCode, err := nftComposerBusinessFacade.SaveWidget(createWidgetResponse,r.Header.Get("Authorization"))
			if err != nil {
				commonResponse.RespondWithJSON(w,statsCode,err.Error())
			} else {
				commonResponse.RespondWithJSON(w, http.StatusOK, result)
			}
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func UpdateWidget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
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
			result,codeStatus, err := nftComposerBusinessFacade.ChangeWidget(updateWidgetResponse,r.Header.Get("Authorization"))
			if err != nil {
				commonResponse.RespondWithJSON(w,codeStatus, err.Error())
			} else {
				commonResponse.SuccessStatus(w, result)
			}
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

// Find project by user ID
func GetWidgetDetails(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		vars := mux.Vars(r)
		if vars["widgetId"] != "" {
			results, err1 := nftComposerBusinessFacade.FindWidgetByWidgetId(vars["widgetId"])
			if err1 != nil {
				errors.BadRequest(w, err1.Error())
			} else {
				commonResponse.RespondWithJSON(w ,200, results);
			}
		} else {
			errors.BadRequest(w, "")
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}