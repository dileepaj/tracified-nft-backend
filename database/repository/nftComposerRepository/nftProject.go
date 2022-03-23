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

func (r *NFTComposerProjectRepository) SaveChart(chart models.Chart) (string, error) {
	return repository.Save[models.Chart](chart, "charts")
}

func (r *NFTComposerProjectRepository) SaveTable(table models.Table) (string, error) {
	return repository.Save[models.Table](table, "tables")
}

func (r *NFTComposerProjectRepository) SaveStat(stat models.StataArray) (string, error) {
	return repository.Save[models.StataArray](stat, "stats")
}

func (r *NFTComposerProjectRepository) SaveProofBot(proofbot models.ProofBotData) (string, error) {
	return repository.Save[models.ProofBotData](proofbot, "proofbot")
}

func (r *NFTComposerProjectRepository) SaveImage(image models.ImageData) (string, error) {
	return repository.Save[models.ImageData](image, "Images")
}

func (r *NFTComposerProjectRepository) FindChartById(idName string, id string) ([]models.Chart, error) {
	var charts []models.Chart
	rst, err := repository.FindById(idName, id, "chart")
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return charts, err
	}
	for rst.Next(context.TODO()) {
		var chart models.Chart
		err = rst.Decode(&chart)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return charts, err
		}
		charts = append(charts, chart)
	}
	return charts, nil
}

func (r *NFTComposerProjectRepository) FindTableById(idName string, id string) ([]models.Table, error) {
	var tables []models.Table
	rst, err := repository.FindById(idName, id, "table")
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return tables, err
	}
	for rst.Next(context.TODO()) {
		var table models.Table
		err = rst.Decode(&table)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return tables, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func (r *NFTComposerProjectRepository) FindStatById(idName string, id string) ([]models.StataArray, error) {
	var stats []models.StataArray
	rst, err := repository.FindById(idName, id, "stat")
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return stats, err
	}
	for rst.Next(context.TODO()) {
		var stat models.StataArray
		err = rst.Decode(&stat)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return stats, err
		}
		stats = append(stats, stat)
	}
	return stats, nil
}

func (r *NFTComposerProjectRepository) FindProofBotById(idName string, id string) ([]models.ProofBotData, error) {
	var botdata []models.ProofBotData
	rst, err := repository.FindById(idName, id, "proofbot")
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return botdata, err
	}
	for rst.Next(context.TODO()) {
		var bot models.ProofBotData
		err = rst.Decode(&bot)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return botdata, err
		}
		botdata = append(botdata, bot)
	}
	return botdata, nil
}

func (r *NFTComposerProjectRepository) FindImagesById(idName string, id string) ([]models.ImageData, error) {
	var Images []models.ImageData
	rst, err := repository.FindById(idName, id, "image")
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return Images, err
	}
	for rst.Next(context.TODO()) {
		var image models.ImageData
		err = rst.Decode(&image)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return Images, err
		}
		Images = append(Images, image)
	}
	return Images, nil
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
