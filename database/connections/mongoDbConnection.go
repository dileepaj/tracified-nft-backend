package connections

import (
	"context"
	"log"

	"github.com/dileepaj/tracified-nft-backend/commons"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoSession mongo.Session
var DbName=commons.GoDotEnvVariable("DATABASE_NAME")

func GetMongoSession() (mongo.Session, error) {

	connectionString := commons.GoDotEnvVariable("BE_MONGOLAB_URI")
	if mgoSession == nil {
		var err error
		mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
		if err != nil {
			return nil, err
		}
		mgoSession, err = mongoClient.StartSession()
		if err != nil {
			log.Println("Error while connecting to the DB : " + err.Error())
			return nil, err
		}
	}
	return mgoSession, nil
}
