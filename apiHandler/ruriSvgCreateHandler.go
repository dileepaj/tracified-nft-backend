package apiHandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/marketplaceBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/middleware"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func RuriSVGGenerator(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		var generateSVGRequest models.SvgCreator
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&generateSVGRequest)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
		}
		err = validations.ValidateSVGGenerator(generateSVGRequest)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			result, err := marketplaceBusinessFacade.GenerateSVGFileForRURI(generateSVGRequest)
			if err != nil {
				errors.BadRequest(w, err.Error())
			}
			commonResponse.SuccessStatus[string](w, result)
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}


func GetRURISVGByBatchID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	vars := mux.Vars(r)
	fmt.Println(vars["itemID"])
	results, err1 := marketplaceBusinessFacade.GetRURISVGByBatchID(vars["itemID"])
	if err1 != nil {
		ErrorMessage := err1.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}

}
