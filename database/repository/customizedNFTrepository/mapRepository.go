package customizedNFTrepository

import (
	"context"
	"fmt"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		logs.WarningLogger.Println("Error Occured when trying to convert hex string in to Object(ID) in GetMapByID : MapRepository: ", err.Error())
	}
	rst, sessionerr := connections.GetSessionClient(mapCollection).Find(context.TODO(), bson.M{"_id": objectId})
	if sessionerr != nil {
		return mapData.MapTemplate, sessionerr
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&mapData)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection faq in GetMapByID:MapRepository.go: ", err.Error())
			return mapData.MapTemplate, err
		}
	}
	return mapData.MapTemplate, nil
}

func (r *MapRepository) UpdateMap(Id primitive.ObjectID, update primitive.M) (models.GeneratedMap, error) {
	var mapResult models.GeneratedMap
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient(mapCollection).FindOneAndUpdate(context.TODO(), bson.M{"_id": Id}, update, &opt)
	if rst != nil {
		err := rst.Decode(&mapResult)
		if err != nil {
			logs.ErrorLogger.Println("Error occur while updatirng map in database", err.Error())
			return mapResult, err
		}
		return mapResult, nil
	}
	UpdateError := fmt.Errorf("failed to get result from database")
	logs.ErrorLogger.Println("Failed to get data from data base.")
	return mapResult, UpdateError
}
