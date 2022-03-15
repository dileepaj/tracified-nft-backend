package apiHandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/controllers/nftController"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var createUserObject models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createUserObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	fmt.Println(createUserObject)
	err = validations.ValidateInsertUser(createUserObject)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		_, err1 := nftController.CreateUser(createUserObject)
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			message := "user created"
			err = json.NewEncoder(w).Encode(message)
			if err != nil {
				logs.ErrorLogger.Println(err)
			}
			return
		}
	}
}