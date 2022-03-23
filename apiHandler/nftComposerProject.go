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
	err = validations.ValidateNFTProject(CreateProjectRequest)
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

func SaveChart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")

	var chartRequest models.Chart
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&chartRequest)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateChart(chartRequest)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err1 := nftComposerBusinessFacade.SaveChart(chartRequest)
		if err1 != nil {
			errors.BadRequest(w, err1.Error())
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}

func SaveTable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")

	var tableRequest models.Table
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tableRequest)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateTable(tableRequest)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err1 := nftComposerBusinessFacade.SaveTable(tableRequest)
		if err1 != nil {
			errors.BadRequest(w, err1.Error())
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}

func SaveStat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")

	var statRequest models.StataArray
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&statRequest)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateStat(statRequest)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err1 := nftComposerBusinessFacade.SaveStats(statRequest)
		if err1 != nil {
			errors.BadRequest(w, err1.Error())
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}

func SaveProofBot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")

	var proofbotRequest models.ProofBotData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&proofbotRequest)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateProofBot(proofbotRequest)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err1 := nftComposerBusinessFacade.SaveProofBot(proofbotRequest)
		if err1 != nil {
			errors.BadRequest(w, err1.Error())
		} else {
			commonResponse.SuccessStatus[string](w, result)
		}
	}
}

func SaveImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")

	var imageRequest models.ImageData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&imageRequest)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	err = validations.ValidateImage(imageRequest)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else {
		result, err1 := nftComposerBusinessFacade.SaveImages(imageRequest)
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
		if err1 != "" {
			errors.BadRequest(w, err1)
		} else {
			commonResponse.SuccessStatus[models.ProjectDetail](w, results)
		}
	} else {
		errors.BadRequest(w, "")
	}
}
