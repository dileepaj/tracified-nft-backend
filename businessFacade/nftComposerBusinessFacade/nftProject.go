package nftComposerBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
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

func SaveProofBot(project models.ProofBotData) (string, error) {
	return nftProjectRepository.SaveProofBot(project)
}

func SaveStats(project models.StataArray) (string, error) {
	return nftProjectRepository.SaveStat(project)
}

func GetRecntProjects(userid string) ([]responseDtos.ResponseProject, error) {
	return nftProjectRepository.FindNFTProjectById("userid", userid)
}

func GetRecntProjectDetails(projectId string) (models.ProjectDetail, string) {
	var nftProject models.ProjectDetail
	var barchart []models.Chart
	var piechart []models.Chart
	var bubblechart []models.Chart
	resultProject, err := nftProjectRepository.FindNFTProjectOneById("projectid", projectId)
	if err != nil {
		return nftProject, err.Error()
	} else if resultProject.ProjectId == "" {
		return nftProject, "Invalid ProjectId"
	} else {
		resultWidgets, err1 := nftProjectRepository.FindWidgetsById("projectid", resultProject.ProjectId)
		resultCharts, err2 := nftProjectRepository.FindChartById("projectid", resultProject.ProjectId)
		resultTables, err3 := nftProjectRepository.FindTableById("projectid", resultProject.ProjectId)
		resultStats, err4 := nftProjectRepository.FindStatById("projectid", resultProject.ProjectId)
		resultImages, err5 := nftProjectRepository.FindImagesById("projectid", resultProject.ProjectId)
		resultProofbot, err6 := nftProjectRepository.FindProofBotById("projectid", resultProject.ProjectId)
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
			return nftProject, err.Error()
		}

		if len(resultCharts) != 0 {
			for _, chart := range resultCharts {
				if chart.Type == "BarChart" {
					barchart = append(barchart, chart)
				} else if chart.Type == "PieChart" {
					piechart = append(piechart, chart)
				} else if chart.Type == "BubbleChart" {
					bubblechart = append(bubblechart, chart)
				}
			}
		}

		responseProject := models.ProjectDetail{
			Project:      resultProject,
			Widgets:      resultWidgets,
			BarCharts:    barchart,
			PieCharts:    piechart,
			BubbleCharts: bubblechart,
			Stats:        resultStats,
			Tables:       resultTables,
			Images:       resultImages,
			ProofBot:     resultProofbot,
		}
		return responseProject, ""
	}
}
