package connections

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
build the Database connection with mongodb
**/
func Connect() *mongo.Database {
	connectionString := os.Getenv("BE_MONGOLAB_URI")
	clientOptions := options.Client().ApplyURI(connectionString)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("Error while connecting to the DB : " + err.Error())
	}
	return client.Database("nftBackendQa")
}