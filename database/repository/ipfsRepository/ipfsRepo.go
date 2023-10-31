package ipfsRepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/commons"
	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IpfsRepository struct{}

var Collection = "ipfsfiles"
var DbName = commons.GoDotEnvVariable("DATABASE_NAME")

func (r *IpfsRepository) SaveFileDetails(ipfsObj models.TracifiedDataPacket) (string, error) {
	return repository.Save[models.TracifiedDataPacket](ipfsObj, Collection)
}

func (r *IpfsRepository) GetTdpDetails(key string, id string) (models.TracifiedDataPacket, error) {
	var tdpDetailsArray models.TracifiedDataPacket
	rst := repository.FindOne(key, id, Collection)
	if rst != nil {
		err := rst.Decode(&tdpDetailsArray)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return tdpDetailsArray, nil
		}
		return tdpDetailsArray, nil
	} else {
		return tdpDetailsArray, nil
	}
}

func (r *IpfsRepository) UpdateFileDetails(tdpId string, updateObj models.TracifiedDataPacket) (models.TracifiedDataPacket, error) {
	var response models.TracifiedDataPacket
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session : ", err.Error())
	}
	defer session.EndSession(context.TODO())
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	update := bson.M{"$set": updateObj}
	rst := session.Client().Database(DbName).Collection(Collection).FindOneAndUpdate(context.TODO(), bson.M{"tdpid": tdpId}, update, &opt)
	if rst != nil {
		err := rst.Decode(&response)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while updating data from DB : ", err.Error())
			return response, err
		}
		return response, nil
	} else {
		return response, nil
	}
}
