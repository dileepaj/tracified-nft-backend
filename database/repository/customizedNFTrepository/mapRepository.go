package customizedNFTrepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MapRepository struct{}

var mapCollection = "MapData"

func (r *MapRepository) SaveMap(Generatedmap models.GeneratedMap) (string, error) {
	return repository.Save(Generatedmap, mapCollection)
}

func (r *MapRepository) GetMapByID(id string) (string, error) {
	var mapData models.GeneratedMap
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logs.WarningLogger.Println("Error Occurred when trying to convert hex string in to Object(ID) in GetMapByID : MapRepository: ", err.Error())
	}
	rst, sessionerr := connections.GetSessionClient(mapCollection).Find(context.TODO(), bson.M{"_id": objectId})
	if sessionerr != nil {
		return mapData.MapTemplate, sessionerr
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&mapData)
		if err != nil {
			logs.ErrorLogger.Println("Error occurred while retrieving data from collection faq in GetMapByID:MapRepository.go: ", err.Error())
			return mapData.MapTemplate, err
		}
	}
	return mapData.MapTemplate, nil
}
