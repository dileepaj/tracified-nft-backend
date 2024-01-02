package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateCollection(collection models.NFTCollection) (string, error) {
	return CollectionRepository.SaveCollection(collection)
}
func GetAllCollections() ([]models.NFTCollection, error) {
	return CollectionRepository.GetAllCollections()
}

func GetCollectionByUserPK(userid string) ([]models.NFTCollection, error) {
	return CollectionRepository.FindCollectionbyUserPK("userid", userid)
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
		"$set": bson.M{"isprivate": UpdateObject.IsPrivate},
	}
	return CollectionRepository.UpdateCollectionVisibility(UpdateObject, update)
}
