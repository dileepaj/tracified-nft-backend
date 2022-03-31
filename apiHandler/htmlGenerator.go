package apiHandler

import (
	"encoding/json"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/nftComposerBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/middleware"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
)

// handel the HTML generate POST request(Generatee HTML NFT)
func HTMLFileGenerator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		var generateHTMLRequest models.HtmlGenerator
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&generateHTMLRequest)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
		}
		err = validations.ValidateHtmlGenerator(generateHTMLRequest)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			result, err := nftComposerBusinessFacade.GenerateHTMLFile(generateHTMLRequest)
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
