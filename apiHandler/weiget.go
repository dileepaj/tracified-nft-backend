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

//Save the weiget data in a DB
func SaveWeiget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var createWeigetObject models.Weiget
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createWeigetObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateInsertWeiget(createWeigetObject)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err1 := nftcomposercontroller.SaveWeiget(createWeigetObject)
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