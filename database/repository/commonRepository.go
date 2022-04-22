package repository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
	Common crud operations with MongoDB
	Can execute crud oprations Passing DB name and colletion name
	Save
	FindOne
	Find
	FindOneAndUpdate
	Update
	UpdateOne
	FindById
	FindById1AndId2
	FindIdOneAndNotId2
	FindByMultipleValue
	Remove
*/

// Insert Document
func Save[T models.SaveType](model T, collection string) (string, error) {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())

	rst, err := session.Client().Database(connections.DbName).Collection(collection).InsertOne(context.TODO(), model)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return "", err
	}
	id := rst.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

// Insert Documets
func InsertMany[T models.InsertManyType](model T, collection string) (string, error) {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())

	var docs []interface{}
	for _, t := range model {
		docs = append(docs, t)
	}
	rst, err := session.Client().Database(connections.DbName).Collection(collection).InsertOne(context.TODO(), model)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return "Error while inserting widgets", err
	}
	id := rst.InsertedID.(primitive.ObjectID)
	return id.String(), nil
}

// Find Documets by Id
func FindById(idName string, id string, collection string) (*mongo.Cursor, error) {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	findOptions.SetProjection(bson.M{"otp": 0})
	rst, err := session.Client().Database(connections.DbName).Collection(collection).Find(context.TODO(), bson.D{{idName, id}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

// Find One Documets
func FindOne[T models.FindOneType](idName string, id T, collection string) *mongo.SingleResult {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())

	findOptions := options.FindOne()
	findOptions.SetProjection(bson.M{"otp": 0})
	rst := session.Client().Database(connections.DbName).Collection(collection).FindOne(context.TODO(), bson.D{{idName, id}}, findOptions)
	return rst
}

// Retrive the Documents filter bt the Id1 and Id2
func FindById1AndId2(idName1 string, id1 string, idName2 string, id2 string, collection string) (*mongo.Cursor, error) {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	rst, err := session.Client().Database(connections.DbName).Collection(collection).Find(context.TODO(), bson.D{{idName1, id1}, {idName2, id2}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

// Retrive all document bt Id1 and Not equal to Id2
func FindById1AndNotId2(idName1 string, id1 string, idName2 string, id2 string, collection string) (*mongo.Cursor, error) {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	rst, err := session.Client().Database(connections.DbName).Collection(collection).Find(context.TODO(), bson.D{{idName1, id1}, {idName2, id2}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

// Find Document using multiple ids
func FindByFieldInMultipleValus(fields string, tags []string, collection string) (*mongo.Cursor, error) {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	rst, err := session.Client().Database(connections.DbName).Collection(collection).Find(context.TODO(), bson.D{{fields, bson.D{{"$in", tags}}}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

// Find Documet and Update
func FindOneAndUpdate(findBy string, value string, update primitive.M, projectionData primitive.M, collection string) *mongo.SingleResult {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())

	after := options.After
	projection := projectionData
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Projection:     &projection,
	}
	rst := session.Client().Database(connections.DbName).Collection(collection).FindOneAndUpdate(context.TODO(), bson.M{findBy: value}, update, &opt)
	return rst
}

// Delete Document
func Remove(idName string, id, collection string) (int64, error) {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	result, err := session.Client().Database(connections.DbName).Collection(collection).DeleteMany(context.TODO(), bson.M{idName: id})
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return 0, err
	}
	return result.DeletedCount, nil
}
