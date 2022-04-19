package connections

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoSession mongo.Session
var DbName="nftBackendQa"

func GetMongoSession() (mongo.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	connectionString := os.Getenv("BE_MONGOLAB_URI")
	if mgoSession == nil {
		var err error
		mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
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
