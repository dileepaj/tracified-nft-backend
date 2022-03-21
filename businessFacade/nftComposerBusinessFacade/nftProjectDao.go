package nftComposerBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
)

func SaveProject(project models.NFTComposerProject) (string, error) {
	return nftProjectRepository.SaveNFTComposerProject(project)
}

func GetRecntProjects(userid string) ([]responseDtos.ResponseProject, error) {
	return nftProjectRepository.FindNFTProjectById("userid",userid)
}

func GetRecntProjectDetails(userid string) ([]responseDtos.ResponseProject, error) {
	return nftProjectRepository.FindNFTProjectById("userid",userid)
}