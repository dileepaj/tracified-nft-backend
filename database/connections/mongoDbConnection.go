package connections

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoSession mongo.Session
var DbName="nftBackendQa"

func GetMongoSession() (mongo.Session, error) {

	connectionString := os.Getenv("BE_MONGOLAB_URI")
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
