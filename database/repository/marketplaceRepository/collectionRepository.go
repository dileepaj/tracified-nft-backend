package marketplaceRepository

import (
	"context"
	"fmt"
	"time"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
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

func (r *CollectionRepository) SaveSVG(svg models.SVG) (string, error) {
	return repository.Save[models.SVG](svg, Svg)
}

func (repository *CollectionRepository) FindCollectionbyUserPK(idName1 string, id1 string) ([]models.NFTCollection, error) {
	var collections []models.NFTCollection
	if idName1 != "" {
		findOptions := options.Find()
		rst, err := connections.Connect().Collection("collections").Find(context.TODO(), bson.D{{idName1, id1}}, findOptions)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return collections, err
		}
		for rst.Next(context.TODO()) {
			var collection models.NFTCollection
			err = rst.Decode((&collection))
			if err != nil {
				logs.ErrorLogger.Println(err.Error())
				return collections, err
			}
			collections = append(collections, collection)
		}
		return collections, nil
	} else {
		return collections, nil
	}
}

func (repository *CollectionRepository) GetAllCollections() ([]models.NFTCollection, error) {
	var collections []models.NFTCollection
	findOptions := options.Find()
	findOptions.SetLimit(10)
	rst, err := connections.Connect().Collection("collections").Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return collections, err
	}
	for rst.Next(context.TODO()) {
		var collection models.NFTCollection
		err = rst.Decode((&collection))
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return collections, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

func (repository *CollectionRepository) UpdateCollection(collection requestDtos.UpdateCollection) (responseDtos.ResponseCollectionUpdate, error) {
	var responseCollection responseDtos.ResponseCollectionUpdate
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	update := bson.M{
		"$set": bson.M{"collectionname": collection.CollectionName},
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := connections.Connect().Collection("collections").FindOneAndUpdate(ctx, bson.M{"_id": collection.Id}, update, &opt).Decode(&responseCollection)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	return responseCollection, err
}

func (repository *CollectionRepository) DeleteCollectionByUserPK(collection requestDtos.DeleteCollectionByUserPK) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := connections.Connect().Collection("collections").DeleteOne(ctx, bson.M{"userid": collection.UserId})
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	fmt.Printf("Delete One removed %v document(s)\n", result.DeletedCount)
	return err
}

func (repository *CollectionRepository) DeleteCollectionById(collection requestDtos.DeleteCollectionById) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := connections.Connect().Collection("collections").DeleteOne(ctx, bson.M{"_id": collection.Id})
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
	return err
}

func (repository *CollectionRepository) FindCollectionbyId(_id string) ([]models.NFTCollection, error) {
	var collections []models.NFTCollection
	if _id != "" {
		findOptions := options.Find()
		objectId, err := primitive.ObjectIDFromHex(_id)
		if err != nil {
			logs.ErrorLogger.Println("Invalid id")
			return collections, nil
		}
		rst, err := connections.Connect().Collection("collections").Find(context.TODO(), bson.D{{"_id", objectId}}, findOptions)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return collections, err
		}
		for rst.Next(context.TODO()) {
			var collection models.NFTCollection
			err = rst.Decode((&collection))
			if err != nil {
				logs.ErrorLogger.Println(err.Error())
				return collections, err
			}
			collections = append(collections, collection)

		}
		return collections, nil
	} else {
		return collections, nil
	}
}
