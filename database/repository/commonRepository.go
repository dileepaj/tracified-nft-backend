package repository

import (
	"context"
	"log"
	"time"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Save[T models.SaveType](model T, collection string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rst, err := connections.Connect().Collection(collection).InsertOne(ctx, model)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return "", err
	}
	log.Println("----------------------data saved-------------------------")
	id := rst.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func InsertMany[T models.InsertManyType](model T, collection string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var docs []interface{}
	for _, t := range model {
		docs = append(docs, t)
	}
	rst, err := connections.Connect().Collection(collection).InsertOne(ctx, model)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return "Error while inserting widgets", err
	}
	id := rst.InsertedID.(primitive.ObjectID)
	return id.String(), nil
}

func FindById(idName string, id string, collection string) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	rst, err := connections.Connect().Collection(collection).Find(context.TODO(), bson.D{{idName, id}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

func FindOne[T models.FindOneType](idName string, id T, collection string) *mongo.SingleResult {
	rst := connections.Connect().Collection(collection).FindOne(context.TODO(), bson.D{{idName, id}})
	return rst
}

func FindById1AndId2(idName1 string, id1 string, idName2 string, id2 string, collection string) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	rst, err := connections.Connect().Collection(collection).Find(context.TODO(), bson.D{{idName1, id1}, {idName2, id2}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

func FindById1AndNotId2(idName1 string, id1 string, idName2 string, id2 string, collection string) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	rst, err := connections.Connect().Collection(collection).Find(context.TODO(), bson.D{{idName1, id1}, {idName2, id2}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

func FindByFieldInMultipleValus(fields string, tags []string, collection string) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	rst, err := connections.Connect().Collection(collection).Find(context.TODO(), bson.D{{fields, bson.D{{"$in", tags}}}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

func FindOneAndUpdate(findBy string, value string, update primitive.M, collection string) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	err := connections.Connect().Collection(collection).FindOneAndUpdate(ctx, bson.M{findBy: value}, update, &opt)
	return err
}
