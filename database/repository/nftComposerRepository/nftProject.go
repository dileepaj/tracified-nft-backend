package nftComposerRepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/constants"
	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services/composerimgservice"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NFTComposerProjectRepository struct{}

var NFTComposerProject = "nftComopserProject"

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
	//Call method to upload image to IPFS
	cidHash, errWhenUploadingImageToIpfs := composerimgservice.UploadImageToIpfsWithFolder(constants.ImageWidget, image.Base64Image, image.ProjectId, image.WidgetId, image.TenetId, image.Title)
	if errWhenUploadingImageToIpfs != nil {
		logs.ErrorLogger.Println(errWhenUploadingImageToIpfs.Error())
		return "", errWhenUploadingImageToIpfs
	}

	//Populate the saving object with the IPFS hash
	saveData := models.SaveImageData{
		WidgetId:    image.WidgetId,
		ProjectId:   image.ProjectId,
		Title:       image.Title,
		Base64Image: image.Base64Image,
		Cid:         cidHash,
	}

	return repository.Save[models.SaveImageData](saveData, "images")
}

func (r *NFTComposerProjectRepository) SaveTimeline(timeline models.Timeline) (string, error) {
	return repository.Save[models.Timeline](timeline, "timeline")
}

func (r *NFTComposerProjectRepository) FindChartById(idName string, id string) ([]models.Chart, error) {
	var charts []models.Chart
	rst, err := repository.FindById(idName, id, "charts")
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
	rst, err := repository.FindById(idName, id, "tables")
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
	rst, err := repository.FindById(idName, id, "stats")
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
	rst, err := repository.FindById(idName, id, "images")
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

func (r *NFTComposerProjectRepository) FindTimelineById(idName string, id string) ([]models.Timeline, error) {
	var Images []models.Timeline
	rst, err := repository.FindById(idName, id, "timeline")
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return Images, err
	}
	for rst.Next(context.TODO()) {
		var image models.Timeline
		err = rst.Decode(&image)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return Images, err
		}
		Images = append(Images, image)
	}
	return Images, nil
}

func (r *NFTComposerProjectRepository) FindNFTProjectById(idName string, id string) ([]models.NFTComposerProject, error) {
	var projects []models.NFTComposerProject
	rst, err := repository.FindById(idName, id, NFTComposerProject)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return projects, err
	}
	for rst.Next(context.TODO()) {
		var project models.NFTComposerProject
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

func (r *NFTComposerProjectRepository) UpdateProject(update requestDtos.UpdateProjectRequest) (models.NFTComposerProject, error) {
	var project models.NFTComposerProject
	pByte, err := bson.Marshal(update)
	if err != nil {
		return project, err
	}
	var updateNew bson.M
	err = bson.Unmarshal(pByte, &updateNew)
	if err != nil {
		return project, err
	}
	rst := connections.GetSessionClient(NFTComposerProject).FindOneAndUpdate(context.TODO(), bson.M{"projectid": update.ProjectId}, bson.D{{Key: "$set", Value: updateNew}})
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

func (r *NFTComposerProjectRepository) UpdateChart(chart requestDtos.UpdateChartRequest) (models.Chart, error) {
	var chartdata models.Chart
	pByte, err := bson.Marshal(chart)
	if err != nil {
		return chartdata, err
	}
	var updateNew bson.M
	err = bson.Unmarshal(pByte, &updateNew)
	if err != nil {
		return chartdata, err
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("charts").FindOneAndUpdate(context.TODO(), bson.M{"widgetid": chart.WidgetId}, bson.D{{Key: "$set", Value: updateNew}}, &opt)
	if rst != nil {
		err := rst.Decode(&chartdata)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return chartdata, err
		}
		return chartdata, nil
	} else {
		return chartdata, nil
	}
}

func (r *NFTComposerProjectRepository) UpdateTable(table requestDtos.UpdateTableRequest) (models.Table, error) {
	var tabledata models.Table
	pByte, err := bson.Marshal(table)
	if err != nil {
		return tabledata, err
	}
	var updateNew bson.M
	err = bson.Unmarshal(pByte, &updateNew)
	if err != nil {
		return tabledata, err
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("tables").FindOneAndUpdate(context.TODO(), bson.M{"widgetid": table.WidgetId}, bson.D{{Key: "$set", Value: updateNew}}, &opt)
	if rst != nil {
		err := rst.Decode(&tabledata)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return tabledata, err
		}
		return tabledata, nil
	} else {
		return tabledata, nil
	}
}

func (r *NFTComposerProjectRepository) UpdateImage(image requestDtos.UpdateImageRequest) (models.ImageData, error) {
	var imageData models.ImageData
	pByte, err := bson.Marshal(image)
	if err != nil {
		return imageData, err
	}
	var updateNew bson.M
	err = bson.Unmarshal(pByte, &updateNew)
	if err != nil {
		return imageData, err
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("images").FindOneAndUpdate(context.TODO(), bson.M{"widgetid": image.WidgetId}, bson.D{{Key: "$set", Value: updateNew}}, &opt)
	if rst != nil {
		err := rst.Decode(&imageData)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return imageData, err
		}
		return imageData, nil
	} else {
		return imageData, nil
	}
}

func (r *NFTComposerProjectRepository) UpdateTimeline(timeline requestDtos.UpdateTimelineRequest) (models.Timeline, error) {
	var timelineData models.Timeline
	pByte, err := bson.Marshal(timeline)
	if err != nil {
		return timelineData, err
	}
	var updateNew bson.M
	err = bson.Unmarshal(pByte, &updateNew)
	if err != nil {
		return timelineData, err
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("timeline").FindOneAndUpdate(context.TODO(), bson.M{"widgetid": timeline.WidgetId}, bson.D{{Key: "$set", Value: updateNew}}, &opt)
	if rst != nil {
		err := rst.Decode(&timelineData)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return timelineData, err
		}
		return timelineData, nil
	} else {
		return timelineData, nil
	}
}

func (r *NFTComposerProjectRepository) UpdateProofBot(proofbot requestDtos.UpdateProofBotRequest) (models.ProofBotData, error) {
	var bot models.ProofBotData
	pByte, err := bson.Marshal(proofbot)
	if err != nil {
		return bot, err
	}
	var updateNew bson.M
	err = bson.Unmarshal(pByte, &updateNew)
	if err != nil {
		return bot, err
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("proofbot").FindOneAndUpdate(context.TODO(), bson.M{"widgetid": proofbot.WidgetId}, bson.D{{Key: "$set", Value: updateNew}}, &opt)
	if rst != nil {
		err := rst.Decode(&bot)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return bot, err
		}
		return bot, nil
	} else {
		return bot, nil
	}
}

func (r *NFTComposerProjectRepository) UpdateStats(stat requestDtos.UpdateStatsRequest) (models.StataArray, error) {
	var stastData models.StataArray
	pByte, err := bson.Marshal(stat)
	if err != nil {
		return stastData, err
	}
	var updateNew bson.M
	err = bson.Unmarshal(pByte, &updateNew)
	if err != nil {
		return stastData, err
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("stats").FindOneAndUpdate(context.TODO(), bson.M{"widgetid": stat.WidgetId}, bson.D{{Key: "$set", Value: updateNew}}, &opt)
	if rst != nil {
		err := rst.Decode(&stastData)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return stastData, err
		}
		return stastData, nil
	} else {
		return stastData, nil
	}
}
