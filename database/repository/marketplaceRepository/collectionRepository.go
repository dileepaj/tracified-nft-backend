package marketplaceRepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionRepository struct{}

var Collection = "collections"
var Svg = "svg"
var Txn = "txn"

func (r *CollectionRepository) SaveCollection(collection models.NFTCollection) (string, error) {
	return repository.Save[models.NFTCollection](collection, Collection)
}

func (r *CollectionRepository) FindCollectionbyUserPK(idName string, id string) ([]models.NFTCollection, error) {
	var collections []models.NFTCollection
	rst, err := repository.FindById(idName, id, Collection)
	if err != nil {
		return collections, err
	}
	for rst.Next(context.TODO()) {
		var collection models.NFTCollection
		err = rst.Decode(&collection)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return collections, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

func (r *CollectionRepository) FindCollectionbyPublickey(idName string, id string) ([]models.NFTCollection, error) {
	var collections []models.NFTCollection
	rst, err := repository.FindById(idName, id, Collection)
	if err != nil {
		return collections, err
	}
	for rst.Next(context.TODO()) {
		var collection models.NFTCollection
		err = rst.Decode(&collection)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return collections, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

func (r *CollectionRepository) GetAllCollections() ([]models.NFTCollection, error) {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session in getAllCollection : CollectionRepository.go : ", err.Error())
	}
	defer session.EndSession(context.TODO())

	var collections []models.NFTCollection
	findOptions := options.Find()
	result, err := session.Client().Database(connections.DbName).Collection(Collection).Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println("Error occured when trying to connect to DB and excute Find query in GetAllCollection:CollectionRepository.go: ", err.Error())
		return collections, err
	}
	for result.Next(context.TODO()) {
		var collection models.NFTCollection
		err = result.Decode(&collection)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection favourites in GetAllCollections:collectionsRepository.go: ", err.Error())
			return collections, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

func (r *CollectionRepository) UpdateCollection(collection requestDtos.UpdateCollection) (models.NFTCollection, error) {
	var responseCollectionStatus models.NFTCollection
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	update := bson.M{
		"$set": bson.M{"collectionname": collection.CollectionName},
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := session.Client().Database(connections.DbName).Collection("collections").FindOneAndUpdate(context.TODO(), bson.M{"_id": collection.Id}, update, &opt)
	if rst != nil {
		err := rst.Decode((&responseCollectionStatus))
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return responseCollectionStatus, err
		}
		return responseCollectionStatus, nil
	} else {
		return responseCollectionStatus, nil

	}
}

func (r *CollectionRepository) DeleteCollection(collection requestDtos.DeleteCollectionByUserPK) error {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	result, err := session.Client().Database(connections.DbName).Collection("collections").DeleteOne(context.TODO(), bson.M{"_id": collection.UserId})
	if err != nil {
		logs.ErrorLogger.Println("Error occured when Connecting to DB and executing DeleteOne Query in DeleteCollection(CollectionRepository): ", err.Error())
	}
	logs.InfoLogger.Println("collection deleted :", result.DeletedCount)
	return err

}

func (r *CollectionRepository) SaveSVG(svg models.SVG) (string, error) {
	return repository.Save[models.SVG](svg, Svg)
}

func (r *CollectionRepository) UpdateSVGBlockchain(id string, update primitive.M) (models.SVG, error) {
	var svgUpdateResponse models.SVG
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logs.WarningLogger.Println("Error Occured when trying to convert hex string in to Object(ID) in UpdateSVGBlockchain : collectionRepo: ", err.Error())
	}
	logs.InfoLogger.Println("Searching by ID: ", id)
	rst := session.Client().Database(connections.DbName).Collection("svg").FindOneAndUpdate(context.TODO(), bson.M{"_id": objectID}, update, &opt)
	if rst != nil {
		err := rst.Decode((&svgUpdateResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection faq in UpdateSVGBlockchain:collectionRepo.go: ", err.Error())
			return svgUpdateResponse, err
		}
		return svgUpdateResponse, err
	} else {
		return svgUpdateResponse, err
	}

}
