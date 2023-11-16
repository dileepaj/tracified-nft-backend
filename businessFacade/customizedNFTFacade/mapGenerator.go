package customizedNFTFacade

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services/mapGenerator"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveMap(mapdata []models.MapInfo) (string, error) {
	generatedMap := mapGenerator.GenerateMap(mapdata)
	var newMap models.GeneratedMap
	newMap.MapTemplate = generatedMap
	rst, mapSaveErr := mapRepository.SaveMap(newMap)
	if mapSaveErr != nil {
		logs.ErrorLogger.Println("Failed to save map : ", mapSaveErr.Error())
		return "", mapSaveErr
	}
	return rst, nil
}

func GetMapByID(id string) (string, error) {
	rst, err := mapRepository.GetMapByID(id)
	if err != nil {
		logs.ErrorLogger.Println("failed to retrive map : ", err.Error())
		return rst, err
	}
	return rst, nil
}

func UpdateMap(mapdata models.UpdateMap) (string, error) {
	generateMap := mapGenerator.GenerateMap(mapdata.MapData)
	var newMap models.GeneratedMap
	newMap.MapTemplate = generateMap

	update := bson.M{
		"$set": bson.M{"template": newMap.MapTemplate},
	}

	rst, updateError := mapRepository.UpdateMap(mapdata.Id, update)
	if updateError != nil {
		logs.ErrorLogger.Println("Failed to update Map: " + updateError.Error())
		return "failed", updateError
	}
	return string(rst.Id.Hex()), nil
}
