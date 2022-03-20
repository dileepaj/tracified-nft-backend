package apiHandler

import (
	"encoding/json"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/nftComposerBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
)

// handel the HTML generate POST request(Generatee HTML NFT)
func HTMLFileGenerator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	var generateHTMLRequest models.NFTComposerProject
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&generateHTMLRequest)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateInsertHtmlNft(generateHTMLRequest)
	if err != nil {
		errors.BadRequest(w, err.Error())
	}
	result, err := nftComposerBusinessFacade.GenerateHTMLFile(generateHTMLRequest)
	if err != nil {
		errors.BadRequest(w, err.Error())
	}
	commonResponse.SuccessStatus[string](w, result)
}