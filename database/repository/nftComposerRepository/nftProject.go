package nftComposerRepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

type NFTComposerProjectRepository struct{}

var NFTComposerProject = "nftComopserProject"

/**
Save the Json that used to create HTML file of NFT
**/
func (r *NFTComposerProjectRepository) SaveNFTComposerProject(project models.NFTComposerProject) (string, error) {
	return repository.Save[models.NFTComposerProject](project, NFTComposerProject)
}

func (r *NFTComposerProjectRepository) FindNFTProjectById(idName string, id string) ([]responseDtos.ResponseProject, error) {
	var projects []responseDtos.ResponseProject

	rst, err := repository.FindById(idName, id, NFTComposerProject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return projects, err
	}
	for rst.Next(context.TODO()) {
		var project responseDtos.ResponseProject
		err = rst.Decode(&project)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return projects, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func (r *NFTComposerProjectRepository) FindNFTProjectOneById(idName string, id string) (models.NFTComposerProject, error) {
	var project models.NFTComposerProject
	rst := repository.FindOne(idName, id, NFTComposerProject)
	if rst != nil {
		err := rst.Decode(&project)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return project, err
		}
		return project, nil
	} else {
		return project, nil
	}
}
