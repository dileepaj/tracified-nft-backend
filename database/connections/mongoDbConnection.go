package connections

import (
	"context"
	"log"

	"github.com/dileepaj/tracified-nft-backend/commons"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DbName = commons.GoDotEnvVariable("DATABASE_NAME")

func GetMongoSession() (mongo.Session, error) {
	connectionString := commons.GoDotEnvVariable("BE_MONGOLAB_URI")
	var err error
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}
	mgoSession, err := mongoClient.StartSession()
	if err != nil {
		log.Println("Error while connecting to the DB : " + err.Error())
		return nil, err
	}
	defer func() {
		if err = mongoClient.Disconnect(context.TODO()); err != nil {
			logrus.Error("mongoClient Disconnect ", err)
		}
	}()
	return mgoSession, nil
}

func GetSessionClient(collection string) *mongo.Collection {
	session, err := GetMongoSession()
	if err != nil {
		log.Println("Error while getting session " + err.Error())
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	return session.Client().Database(DbName).Collection(collection)
}