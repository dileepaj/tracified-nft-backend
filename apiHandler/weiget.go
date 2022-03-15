package apiHandler

import (
	"encoding/json"
	"net/http"

	nftcomposercontroller "github.com/dileepaj/tracified-nft-backend/controllers/nftComposerController"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
)

//Save the widget data in a DB
func SaveWidget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var createWidgetObject models.Widget
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createWidgetObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateInsertWidget(createWidgetObject)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err1 := nftcomposercontroller.SaveWidget(createWidgetObject)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(result)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}