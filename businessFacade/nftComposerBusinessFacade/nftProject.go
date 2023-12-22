package nftComposerBusinessFacade

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dileepaj/tracified-nft-backend/constants"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services/composerimgservice"
	"github.com/dileepaj/tracified-nft-backend/utilities/commonResponse"
	"github.com/dileepaj/tracified-nft-backend/utilities/errors"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

func SaveProject(project models.NFTComposerProject) (string, error) {
	return nftProjectRepository.SaveNFTComposerProject(project)
}

func SaveChart(project models.Chart) (string, error) {
	return nftProjectRepository.SaveChart(project)
}

func SaveTable(project models.Table) (string, error) {
	return nftProjectRepository.SaveTable(project)
}

func SaveImages(project models.ImageData) (string, error) {
	return nftProjectRepository.SaveImage(project)
}

func SaveTimeline(timeline models.Timeline) (string, error) {
	return nftProjectRepository.SaveTimeline(timeline)
}

func SaveProofBot(project models.ProofBotData) (string, error) {
	return nftProjectRepository.SaveProofBot(project)
}

func SaveStats(project models.StataArray) (string, error) {
	return nftProjectRepository.SaveStat(project)
}

func GetRecntProjects(userid string) ([]models.NFTComposerProject, error) {
	return nftProjectRepository.FindNFTProjectById("userid", userid)
}

func GetRecntProjectDetails(projectId string) (models.ProjectDetail, string) {
	var nftProject models.ProjectDetail
	var barchart []models.ChartAndWidget
	var piechart []models.ChartAndWidget
	var bubblechart []models.ChartAndWidget
	var tableWithWidget []models.TableWithWidget
	resultProject, err := nftProjectRepository.FindNFTProjectOneById("projectid", projectId)
	if err != nil {
		return nftProject, err.Error()
	} else if resultProject.ProjectId == "" {
		return nftProject, "Invalid ProjectId"
	} else {

		resultCharts, err2 := nftProjectRepository.FindChartById("projectid", resultProject.ProjectId)
		resultTables, err3 := nftProjectRepository.FindTableById("projectid", resultProject.ProjectId)
		resultStats, err4 := nftProjectRepository.FindStatById("projectid", resultProject.ProjectId)
		resultImages, err5 := nftProjectRepository.FindImagesById("projectid", resultProject.ProjectId)
		ProoBotData, err6 := nftProjectRepository.FindProofBotById("projectid", resultProject.ProjectId)
		resultTimeline, err7 := nftProjectRepository.FindTimelineById("projectid", resultProject.ProjectId)

		if err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil || err7 != nil {
			return nftProject, err.Error()
		}

		if len(resultCharts) != 0 {
			for _, chart := range resultCharts {
				if chart.Type == "BarChart" {
					resultWidget, err1 := FindWidgetByWidgetId(chart.WidgetId)
					if err != nil {
						logs.ErrorLogger.Println(err1)
					}
					bar := models.ChartAndWidget{
						Chart:  chart,
						Widget: resultWidget,
					}
					barchart = append(barchart, bar)
				} else if chart.Type == "PieChart" {
					resultWidget, err1 := FindWidgetByWidgetId(chart.WidgetId)
					if err != nil {
						logs.ErrorLogger.Println(err1)
					}
					pie := models.ChartAndWidget{
						Chart:  chart,
						Widget: resultWidget,
					}
					piechart = append(piechart, pie)
				} else if chart.Type == "BubbleChart" {
					resultWidget, err1 := FindWidgetByWidgetId(chart.WidgetId)
					if err != nil {
						logs.ErrorLogger.Println(err1)
					}
					bubble := models.ChartAndWidget{
						Chart:  chart,
						Widget: resultWidget,
					}
					bubblechart = append(bubblechart, bubble)
				}
			}
		}

		if len(resultTables) != 0 {
			for _, table := range resultTables {
				resultWidget, err1 := FindWidgetByWidgetId(table.WidgetId)
				if err != nil {
					logs.ErrorLogger.Println(err1)
				}
				table := models.TableWithWidget{
					Table:  table,
					Widget: resultWidget,
				}
				tableWithWidget = append(tableWithWidget, table)
			}
		}
		responseProject := models.ProjectDetail{
			Project:      resultProject,
			BarCharts:    barchart,
			PieCharts:    piechart,
			BubbleCharts: bubblechart,
			Stats:        resultStats,
			Tables:       tableWithWidget,
			Images:       resultImages,
			ProofBot:     ProoBotData,
			Timeline:     resultTimeline,
		}

		return responseProject, ""
	}
}

func UpdateProject(w http.ResponseWriter, updateProject requestDtos.UpdateProjectRequest) {
	rst, err := nftProjectRepository.UpdateProject(updateProject)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else if rst.ProjectId == "" {
		commonResponse.NoContent(w, "Invalid Product Id")
	} else {
		commonResponse.SuccessStatus(w, rst)
	}
}

func UpdateChart(w http.ResponseWriter, updateChart requestDtos.UpdateChartRequest) {
	fmt.Println(updateChart)
	rst, err := nftProjectRepository.UpdateChart(updateChart)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else if rst.ProjectId == "" {
		commonResponse.NoContent(w, "Invalid WidgetId")
	} else {
		commonResponse.SuccessStatus(w, rst)
	}
}

func UpdateTable(w http.ResponseWriter, updateTable requestDtos.UpdateTableRequest) {
	rst, err := nftProjectRepository.UpdateTable(updateTable)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else if rst.ProjectId == "" {
		commonResponse.NoContent(w, "Invalid WidgetId")
	} else {
		commonResponse.SuccessStatus(w, rst)
	}
}

func UpdateProofBot(w http.ResponseWriter, updateProofBot requestDtos.UpdateProofBotRequest) {
	rst, err := nftProjectRepository.UpdateProofBot(updateProofBot)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else if rst.ProjectId == "" {
		commonResponse.NoContent(w, "Invalid WidgetId")
	} else {
		commonResponse.SuccessStatus(w, rst)
	}
}

func UpdateImages(w http.ResponseWriter, updateImage requestDtos.UpdateImageRequest) {
	//append timestamp to key before uploading to IPFS
	timestamp := time.Now().Format("20060102150405") //YYYYMMDDHHMMSS
	updatedImageTitle := updateImage.Title + timestamp

	//Upload new image to IPFS
	cidHash, errWhenUploadingImageToIpfs := composerimgservice.UploadImageToIpfsWithFolder(constants.ImageWidget, updateImage.Base64Image, updateImage.ProjectId, updateImage.WidgetId, updateImage.TenetId, updatedImageTitle)
	if errWhenUploadingImageToIpfs != nil {
		logs.ErrorLogger.Println(errWhenUploadingImageToIpfs.Error())
		errors.BadRequest(w, errWhenUploadingImageToIpfs.Error())
	}

	//Add the new CID hash to update Object
	updateObj := requestDtos.SaveUpdatedImage{
		WidgetId:    updateImage.WidgetId,
		Title:       updateImage.Title,
		Type:        updateImage.Type,
		Base64Image: updateImage.Base64Image,
		ProjectId:   updateImage.ProjectId,
		Cid:         cidHash,
	}

	rst, err := nftProjectRepository.UpdateImage(updateObj)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else if rst.ProjectId == "" {
		commonResponse.NoContent(w, "Invalid WidgetId")
	} else {
		commonResponse.SuccessStatus(w, rst)
	}
}

func UpdateTimeline(w http.ResponseWriter, updateTimeline requestDtos.UpdateTimelineRequest) {
	rst, err := nftProjectRepository.UpdateTimeline(updateTimeline)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else if rst.ProjectId == "" {
		commonResponse.NoContent(w, "Invalid WidgetId")
	} else {
		commonResponse.SuccessStatus(w, rst)
	}
}

func UpdateStats(w http.ResponseWriter, updateStat requestDtos.UpdateStatsRequest) {
	rst, err := nftProjectRepository.UpdateStats(updateStat)
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else if rst.ProjectId == "" {
		commonResponse.NoContent(w, "Invalid WidgetId")
	} else {
		commonResponse.SuccessStatus(w, rst)
	}
}

func RemoveProjet(w http.ResponseWriter, id string) {
	rst, err := repository.Remove("projectid", id, "nftComopserProject")
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else if rst <= 0 {
		commonResponse.NoContent(w, "Invalid ProjectId")
	} else {
		_, err1 := repository.Remove("projectid", id, "tables")
		if err1 != nil {
			logs.ErrorLogger.Fatalln(err1)
		}
		_, err2 := repository.Remove("projectid", id, "charts")
		if err2 != nil {
			logs.ErrorLogger.Fatalln(err1)
		}
		_, err3 := repository.Remove("projectid", id, "images")
		if err3 != nil {
			logs.ErrorLogger.Fatalln(err1)
		}
		_, err4 := repository.Remove("projectid", id, "proofbot")
		if err4 != nil {
			logs.ErrorLogger.Fatalln(err1)
		}
		_, err5 := repository.Remove("projectid", id, "stats")
		if err5 != nil {
			logs.ErrorLogger.Fatalln(err1)
		}
		_, err6 := repository.Remove("projectid", id, "widget")
		if err6 != nil {
			logs.ErrorLogger.Fatalln(err1)
		}
		_, err7 := repository.Remove("projectid", id, "timeline")
		if err7 != nil {
			logs.ErrorLogger.Fatalln(err1)
		}
		commonResponse.SuccessStatus(w, id)
	}
}

func RemoveTable(w http.ResponseWriter, id string) {
	rst, err := repository.Remove("widgetid", id, "tables")
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else if rst <= 0 {
		commonResponse.NoContent(w, "Invalid WidgetId")
	} else {
		_, err := repository.Remove("widgetid", id, "widget")
		if err != nil {
			logs.ErrorLogger.Fatalln(err)
		}
		commonResponse.SuccessStatus(w, id)
	}
}

func RemoveChart(w http.ResponseWriter, id string) {
	rst, err := repository.Remove("widgetid", id, "charts")
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else if rst <= 0 {
		commonResponse.NoContent(w, "Invalid WidgetId")
	} else {
		_, err := repository.Remove("widgetid", id, "widget")
		if err != nil {
			logs.ErrorLogger.Fatalln(err)
		}
		commonResponse.SuccessStatus(w, id)
	}
}

func RemoveProofBot(w http.ResponseWriter, id string) {
	rst, err := repository.Remove("widgetid", id, "proofbot")
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else if rst <= 0 {
		commonResponse.NoContent(w, "Invalid WidgetId")
	} else {
		_, err := repository.Remove("widgetid", id, "widget")
		if err != nil {
			logs.ErrorLogger.Fatalln(err)
		}
		commonResponse.SuccessStatus(w, id)
	}
}

func RemoveImage(w http.ResponseWriter, id string) {
	rst, err := repository.Remove("widgetid", id, "images")
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else if rst <= 0 {
		commonResponse.NoContent(w, "Invalid WidgetId")
	} else {
		commonResponse.SuccessStatus(w, id)
	}
}

func RemoveStats(w http.ResponseWriter, id string) {
	rst, err := repository.Remove("widgetid", id, "stats")
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else if rst <= 0 {
		commonResponse.NoContent(w, "Invalid WidgetId")
	} else {
		commonResponse.RespondWithJSON(w, http.StatusOK, responseDtos.WidgetIdResponse{
			WidgetId: id,
		})
	}
}

func RemoveTimeline(w http.ResponseWriter, id string) {
	rst, err := repository.Remove("widgetid", id, "timeline")
	if err != nil {
		errors.BadRequest(w, err.Error())
	} else if rst <= 0 {
		commonResponse.NoContent(w, "Invalid WidgetId")
	} else {
		commonResponse.RespondWithJSON(w, http.StatusOK, responseDtos.WidgetIdResponse{
			WidgetId: id,
		})
	}
}
