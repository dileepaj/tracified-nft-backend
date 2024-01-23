package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCollection(collection models.NFTCollection) (string, error) {
	return CollectionRepository.SaveCollection(collection)
}

func getCollectionProjection() bson.D {
	projection := bson.D{
		{Key: "userid", Value: 1},
		{Key: "collectionname", Value: 1},
		{Key: "organizationname", Value: 1},
		{Key: "publickey", Value: 1},
		{Key: "ispublic", Value: 1},
		{Key: "cid", Value: 1},
		{Key: "images", Value: 1},
	}
	return projection
}
func GetNFTCountInCollection(collectionName string) (int64, error) {
	return CollectionRepository.GetNFTCountInCollection(collectionName)
}

func GetAllCollectionsPaginated(pagination requestDtos.CollectionPagination) (models.CollectionPaginationResponse, error) {
	filter := bson.M{
		"ispublic": true,
	}
	var collections []models.NFTCollection
	projection := getCollectionProjection()
	response, err := CollectionRepository.PaginateCollectionResponse(filter, projection, pagination.PageSize, pagination.RequestedPage, "collections", "_id", collections, pagination.SortType)
	if err != nil {
		logs.ErrorLogger.Println("Error occurred :", err.Error())
		return response, err
	}
	return response, nil

}
func GetAllCollections() ([]models.NFTCollection, error) {
	return CollectionRepository.GetAllPublicCollections()
}

func GetCollectionByUserPK(userid string) ([]models.NFTCollection, error) {
	return CollectionRepository.FindCollectionbyUserPK("userid", userid)
}

func GetCollectionByUserIDPaginated(pagination requestDtos.CollectionPagination, publickey string) (models.CollectionPaginationResponse, error) {
	//if pubkey is given send pubkey and get endorsment ID and and attach it to userID in filter. if no response attach to publckey filter in Filter(backwards compatiablity)
	//TODO : write function to get endorsment ID

	//temp filter until endorsment check funtion is made
	filter := bson.M{
		"publickey": publickey,
	}
	var collections []models.NFTCollection
	projection := getCollectionProjection()
	response, err := CollectionRepository.PaginateCollectionResponse(filter, projection, pagination.PageSize, pagination.RequestedPage, "collections", "_id", collections, pagination.SortType)
	if err != nil {
		logs.ErrorLogger.Println("Error occurred :", err.Error())
		return response, err
	}
	return response, nil
}

func GetCollectionByPublicKey(pk string) ([]models.NFTCollection, error) {
	return CollectionRepository.FindCollectionbyPublickey("publickey", pk)
}

func UpdateCollection(update requestDtos.UpdateCollection) (models.NFTCollection, error) {
	return CollectionRepository.UpdateCollection(update)
}
func DeleteCollectionByUserPK(collection requestDtos.DeleteCollectionByUserPK) error {
	return CollectionRepository.DeleteCollection(collection)
}

func GetCollectionByUserPKByMail(userid string, publickey string) ([]models.NFTCollection, error) {
	return CollectionRepository.FindCollectionByKeyAndMail("userid", userid, "publickey", publickey)
}

func UpdateCollectionVisibility(UpdateObject requestDtos.UpdateCollectionVisibility) (models.NFTCollection, error) {
	update := bson.M{
		"$set": bson.M{"ispublic": UpdateObject.IsPublic},
	}
	return CollectionRepository.UpdateCollectionVisibility(UpdateObject, update)
}

func FindCollectionbyUserID(userid string) (models.NFTCollection, error) {
	return CollectionRepository.FindCollectionbyUserID(userid, "userid")
}

func UpdateCollectionDetails(id primitive.ObjectID, UpdateObject models.NFTCollection) (models.NFTCollection, error) {
	return CollectionRepository.UpdateCollectionDetails(id, UpdateObject)
}

func IsCollectionNameTaken(name string) (bool, error) {
	return CollectionRepository.IsCollectionNameTaken(name)
}
