package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
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
