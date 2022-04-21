package apiHandler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/nftComposerBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/dileepaj/tracified-nft-backend/utilities/middleware"
	"github.com/dileepaj/tracified-nft-backend/utilities/validations"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

// save html version of NFt and save it's json verson on DB
func SaveProject(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
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
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func SaveChart(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
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
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func SaveTable(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
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
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func SaveStat(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
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
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func SaveProofBot(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
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
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func SaveImage(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
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
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func SaveTimeline(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		var timelineRequest models.Timeline
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&timelineRequest)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
		}
		err = validations.ValidateTimeline(timelineRequest)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			result, err1 := nftComposerBusinessFacade.SaveTimeline(timelineRequest)
			if err1 != nil {
				errors.BadRequest(w, err1.Error())
			} else {
				commonResponse.SuccessStatus[string](w, result)
			}
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}
// Find project by user ID
func GetRecentProjects(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	time.Sleep(1 * time.Second) //for fixing jwt token validation server time issue 
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		vars := mux.Vars(r)
		if vars["userId"] != "" {
			result, err1 := nftComposerBusinessFacade.GetRecntProjects(vars["userId"])
			if err1 != nil {
				errors.BadRequest(w, err1.Error())
			} else {
				commonResponse.SuccessStatus[[]models.NFTComposerProject](w, result)
			}
		} else {
			errors.BadRequest(w, "")
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

// Find project by user ID
func GetRecentProjectDetails(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
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
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		var projectRequest requestDtos.UpdateProjectRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&projectRequest)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
		}
		err = validations.ValidateUpdateProject(projectRequest)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			nftComposerBusinessFacade.UpdateProject(w, projectRequest)
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func UpdateChart(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")

	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		logs.ErrorLogger.Println(ps.TenantId)
		var updateChartRequest requestDtos.UpdateChartRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&updateChartRequest)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
		}
		err = validations.ValidateUpdateChart(updateChartRequest)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			nftComposerBusinessFacade.UpdateChart(w, updateChartRequest)
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func UpdateTable(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		var updateTableRequest requestDtos.UpdateTableRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&updateTableRequest)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
		}
		err = validations.ValidateUpdateTable(updateTableRequest)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			nftComposerBusinessFacade.UpdateTable(w, updateTableRequest)
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func UpdateProofBot(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")

	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		var updateProofBotRequest requestDtos.UpdateProofBotRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&updateProofBotRequest)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
		}
		err = validations.ValidateUpdateProofBot(updateProofBotRequest)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			nftComposerBusinessFacade.UpdateProofBot(w, updateProofBotRequest)
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func UpdateImage(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		var updateImageRequest requestDtos.UpdateImageRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&updateImageRequest)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
		}
		err = validations.ValidateUpdateImage(updateImageRequest)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			nftComposerBusinessFacade.UpdateImages(w, updateImageRequest)
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func UpdateTimeline(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		var updateTimelineRequest requestDtos.UpdateTimelineRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&updateTimelineRequest)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
		}
		err = validations.ValidateTimelineRequest(updateTimelineRequest)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			nftComposerBusinessFacade.UpdateTimeline(w, updateTimelineRequest)
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func UpdateStats(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json;")
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		var updateStatsRequest requestDtos.UpdateStatsRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&updateStatsRequest)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
		}
		err = validations.ValidateUpdateStats(updateStatsRequest)
		if err != nil {
			errors.BadRequest(w, err.Error())
		} else {
			nftComposerBusinessFacade.UpdateStats(w, updateStatsRequest)
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func RemoveProjet(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		vars := mux.Vars(r)
		if vars["projectId"] != "" {
			nftComposerBusinessFacade.RemoveProjet(w, vars["projectId"])
		} else {
			commonResponse.RespondWithJSON(w, http.StatusBadRequest, "")
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func RemoveChart(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		vars := mux.Vars(r)
		if vars["widgetId"] != "" {
			nftComposerBusinessFacade.RemoveChart(w, vars["widgetId"])
		} else {
			commonResponse.RespondWithJSON(w, http.StatusBadRequest, "")
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func RemoveTable(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		vars := mux.Vars(r)
		if vars["widgetId"] != "" {
			nftComposerBusinessFacade.RemoveTable(w, vars["widgetId"])
		} else {
			commonResponse.RespondWithJSON(w, http.StatusBadRequest, "")
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func RemoveProofBot(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		vars := mux.Vars(r)
		if vars["widgetId"] != "" {
			nftComposerBusinessFacade.RemoveProofBot(w, vars["widgetId"])
		} else {
			commonResponse.RespondWithJSON(w, http.StatusBadRequest, "")
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func RemoveImage(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		vars := mux.Vars(r)
		if vars["widgetId"] != "" {
			nftComposerBusinessFacade.RemoveImage(w, vars["widgetId"])
		} else {
			commonResponse.RespondWithJSON(w, http.StatusBadRequest, "")
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func RemoveStats(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		vars := mux.Vars(r)
		if vars["widgetId"] != "" {
			nftComposerBusinessFacade.RemoveStats(w, vars["widgetId"])
		} else {
			commonResponse.RespondWithJSON(w, http.StatusBadRequest, "")
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}

func RemoveTimeline(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	ps := middleware.HasPermissions(r.Header.Get("Authorization"))
	if ps.Status {
		vars := mux.Vars(r)
		if vars["widgetId"] != "" {
			nftComposerBusinessFacade.RemoveTimeline(w, vars["widgetId"])
		} else {
			commonResponse.RespondWithJSON(w, http.StatusBadRequest, "")
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	logs.ErrorLogger.Println("Status Unauthorized")
	return
}
