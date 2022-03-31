package apiHandler

import (
	"encoding/json"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/nftComposerBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/middleware"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
)

func QueryExecuter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
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
			nftComposerBusinessFacade.QueryExecuter(w, requestWidget)
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}
