package apiHandler

import (
	"encoding/json"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/nftComposerBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/gorilla/mux"
)

// save html version of NFt and save it's json verson on DB
func SaveProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")

	var CreateProjectRequest models.NFTComposerProject
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&CreateProjectRequest)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateInsertHtmlNft(CreateProjectRequest)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err1 := nftComposerBusinessFacade.SaveProject(CreateProjectRequest)
		if err1 != nil {
			errors.BadRequest(w, err1.Error())
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}

// Find project by user ID
func GetRecentProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["userId"] != "" {
		result, err1 := nftComposerBusinessFacade.GetRecntProjects(vars["userId"])
		if err1 != nil {
			errors.BadRequest(w, err1.Error())
		} else {
			commonResponse.SuccessStatus[[]responseDtos.ResponseProject](w, result)
		}
	} else {
		errors.BadRequest(w, "")
	}
}

// Find project by user ID
func GetRecentProjectDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	vars := mux.Vars(r)
	if vars["projectId"] != "" {
		results, err1 := nftComposerBusinessFacade.GetRecntProjectDetails(vars["projectId"])
		if err1 != nil {
			errors.BadRequest(w, err1.Error())
		} else {
			commonResponse.SuccessStatus[[]responseDtos.ResponseProject](w, results)
		}
	} else {
		errors.BadRequest(w, "")
	}
}
