package apiHandler

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	nftcomposercontroller "github.com/dileepaj/tracified-nft-backend/controllers/nftComposerController"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services/htmlGeneretorService/htmlgenerator"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
)

//handel the HTML generate POST request(create HTML NFT)
func HTMLFileGenerator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var CreateHTMLObject models.HtmlGenerator
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&CreateHTMLObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	//retrive the generated html template
	results, err := htmlgenerator.GenerateNFTTemplate(CreateHTMLObject)
	//convert result to byte Array
	ByteResults := base64.StdEncoding.EncodeToString([]byte(results))
	if err != nil {
		ErrorMessage := err.Error()
		errors.BadRequest(w, ErrorMessage)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(ByteResults)
		if err != nil {
			logs.ErrorLogger.Println(err)
		}
		return
	}
}

func SaveHTMLData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var CreateHTMLOfNFTObject models.HtmlGenerator
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&CreateHTMLOfNFTObject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateInsertHtmlNft(CreateHTMLOfNFTObject)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		_, err1 := nftcomposercontroller.SaveHtmlContentData(CreateHTMLOfNFTObject)
		//convert result to byte Array
		if err1 != nil {
			ErrorMessage := err1.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			//retrive the generated html template
			results, err := htmlgenerator.GenerateNFTTemplate(CreateHTMLOfNFTObject)
			//convert result to byte Array
			ByteResults := base64.StdEncoding.EncodeToString([]byte(results))
			if err != nil {
				ErrorMessage := err.Error()
				errors.BadRequest(w, ErrorMessage)
				return
			} else {
				w.WriteHeader(http.StatusOK)
				err := json.NewEncoder(w).Encode(ByteResults)
				if err != nil {
					logs.ErrorLogger.Println(err)
				}
				return
			}
		}
	}
}
