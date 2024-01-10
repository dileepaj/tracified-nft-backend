package marketplaceRepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/commons"
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

var DbName = commons.GoDotEnvVariable("DATABASE_NAME")
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

func (r *CollectionRepository) GetAllPublicCollections() ([]models.NFTCollection, error) {
	var collections []models.NFTCollection
	findOptions := options.Find()
	result, err := connections.GetSessionClient(Collection).Find(context.TODO(), bson.M{"ispublic": true}, findOptions)
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
	update := bson.M{
		"$set": bson.M{"collectionname": collection.CollectionName},
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient(Collection).FindOneAndUpdate(context.TODO(), bson.M{"_id": collection.Id}, update, &opt)
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
	result, err := connections.GetSessionClient(Collection).DeleteOne(context.TODO(), bson.M{"_id": collection.UserId})
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
	rst := connections.GetSessionClient("svg").FindOneAndUpdate(context.TODO(), bson.M{"_id": objectID}, update, &opt)
	if rst != nil {
		err := rst.Decode((&svgUpdateResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error Occured while retreving data from collection faq in UpdateSVGBlockchain:collectionRepo.go: ", err.Error())
			return svgUpdateResponse, err
		}
		return svgUpdateResponse, err
	} else {
		return svgUpdateResponse, err
	}

}

func (r *CollectionRepository) FindCollectionByKeyAndMail(idName1 string, id1 string, idName2 string, id2 string) ([]models.NFTCollection, error) {
	var collection []models.NFTCollection
	rst, err := repository.FindById1AndId2(idName1, id1, idName2, id2, Collection)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return collection, err
	}
	for rst.Next(context.TODO()) {
		var collections models.NFTCollection
		err = rst.Decode(&collections)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return collection, err
		}
		collection = append(collection, collections)
	}
	return collection, nil
}

func (r *CollectionRepository) UpdateCollectionVisibility(UpdateObject requestDtos.UpdateCollectionVisibility, update primitive.M) (models.NFTCollection, error) {
	var CollcetionUpdateResponse models.NFTCollection
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("collections").FindOneAndUpdate(context.TODO(), bson.M{"_id": UpdateObject.Id}, update, &opt)
	if rst != nil {
		err := rst.Decode((&CollcetionUpdateResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error Occured while Update Collection visibility", err.Error())
			return CollcetionUpdateResponse, err
		}
		return CollcetionUpdateResponse, err
	}
	return CollcetionUpdateResponse, nil
}

func (r *CollectionRepository) FindCollectionByKeyAndMailAndName(idName1 string, id1 string, idName2 string, id2 string, idName3 string, id3 string) (models.NFTCollection, error) {
	var collection models.NFTCollection
	rst, err := repository.FindById1Id2Id3(id1, idName1, id2, idName2, id3, idName3, Collection)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return collection, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&collection)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection document in GetDocumentByID:documentRepository.go: ", err.Error())
			return collection, err
		}
	}
	return collection, nil
}

func (r *CollectionRepository) UpdateCollectionDetails(Id primitive.ObjectID, updateObj models.NFTCollection) (models.NFTCollection, error) {
	var response models.NFTCollection
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session : ", err.Error())
	}
	defer session.EndSession(context.TODO())
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	update := bson.M{"$set": updateObj}
	rst := session.Client().Database(DbName).Collection(Collection).FindOneAndUpdate(context.TODO(), bson.M{"_id": Id}, update, &opt)
	if rst != nil {
		err := rst.Decode(&response)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while updating data from DB : ", err.Error())
			return response, err
		}
		return response, nil
	} else {
		return response, nil
	}
}
