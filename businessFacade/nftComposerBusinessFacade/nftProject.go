package nftComposerBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
)

func SaveProject(project models.NFTComposerProject) (string, error) {
	return nftProjectRepository.SaveNFTComposerProject(project)
}

func GetRecntProjects(userid string) ([]responseDtos.ResponseProject, error) {
	return nftProjectRepository.FindNFTProjectById("userid", userid)
}

func GetRecntProjectDetails(projectId string) (models.ProjectWithWidgets, string) {
	var nftProject models.ProjectWithWidgets
	resultProject, err := nftProjectRepository.FindNFTProjectOneById("projectid", projectId)
	if err != nil {
		return nftProject, err.Error()
	} else if resultProject.ProjectId == "" {
		return nftProject, "Invalid ProjectId"
	} else {
		resultWidgets, err := nftProjectRepository.FindWidgetsById("projectid", resultProject.ProjectId)
		if err != nil {
			return nftProject, err.Error()
		}
		responseProject := models.ProjectWithWidgets{
			NFTComposerProject: resultProject,
			WidgetDetails:      resultWidgets,
		}
		return responseProject,""
	}
}
