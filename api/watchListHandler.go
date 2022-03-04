package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/controllers"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
)

func CreateWatchList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var createWatchListObject models.WatchList
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createWatchListObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	fmt.Println(createWatchListObject)
	err = validations.ValidateInsertWatchList(createWatchListObject)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		_, err1 := controllers.CreateWatchList(createWatchListObject)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			message := "SAVED WatchList"
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}