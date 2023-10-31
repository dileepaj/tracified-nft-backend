package apiHandler

import (
	"encoding/json"
	"net/http"

	ipfsbusinessfacade "github.com/dileepaj/tracified-nft-backend/businessFacade/ipfsBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/commons"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
)

func UploadFilesToIpfs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var ipfsObject models.IpfsObjectForTDP
	decoder := json.NewDecoder(r.Body)

	errWhenDecodingBody := decoder.Decode(&ipfsObject)
	if errWhenDecodingBody != nil {
		logs.ErrorLogger.Println(errWhenDecodingBody)
	}

	//validate the body
	errWhenValidatingBody := validations.ValidateUploadIpfsFile(ipfsObject)
	if errWhenValidatingBody != nil {
		errors.BadRequest(w, errWhenValidatingBody.Error())
	} else {
		//calling the business facade
		cid, errWhenCallingTheFacade := ipfsbusinessfacade.UploadFilesToIpfs(ipfsObject)
		if errWhenCallingTheFacade != nil {
			ErrorMessage := errWhenCallingTheFacade.Error()
			errors.BadRequest(w, ErrorMessage)
			return
		} else {
			url := commons.GoDotEnvVariable("IPFSURL")
			response := models.IpfsResponse{
				Message: "File uploaded to IPFS",
				Url:     url + cid,
			}
			w.WriteHeader(http.StatusOK)
			errWhenEncodingMsg := json.NewEncoder(w).Encode(response)
			if errWhenEncodingMsg != nil {
				logs.ErrorLogger.Println(errWhenEncodingMsg)
			}
			return
		}

	}

}
