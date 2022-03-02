package commonRepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/connections"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct{}

func (mongoRepository *MongoRepository)FindById(idName1 string, id1 string,collection string) (*mongo.Cursor, error) {
	if (idName1 != "" && collection !="") {
		findOptions := options.Find()
		findOptions.SetSort(bson.D{{"timestamp", -1}})
		rst, err := connections.Connect().Collection(collection).Find(context.TODO(), bson.D{{idName1, id1}}, findOptions)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nil, err
		}
		return rst, nil
	} else {
		return nil, nil
	}
}