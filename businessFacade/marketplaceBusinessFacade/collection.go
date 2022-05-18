package marketplaceBusinessFacade

import (
	"fmt"

	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
)

func CreateCollection(collection models.NFTCollection) (string, error) {
	return CollectionRepository.SaveCollection(collection)
}

func CreateSVG(svg models.SVG) (string, error) {
	return CollectionRepository.SaveSVG(svg)
}

func GetAllCollections() ([]models.NFTCollection, error) {
	fmt.Println("Calling repo...")
	return CollectionRepository.GetAllCollections()
}
func GetCollectionByUserPK(userid string) ([]models.NFTCollection, error) {
	return CollectionRepository.FindCollectionbyUserPK("userid", userid)

}

func GetCollectionById(_id string) ([]models.NFTCollection, error) {
	return CollectionRepository.FindCollectionbyId(_id)
}

func UpdateCollection(update requestDtos.UpdateCollection) (responseDtos.ResponseCollectionUpdate, error) {
	return CollectionRepository.UpdateCollection(update)
}
func DeleteCollectionByUserPK(collection requestDtos.DeleteCollectionByUserPK) error {
	return CollectionRepository.DeleteCollectionByUserPK(collection)
}

func DeleteCollectionById(collection requestDtos.DeleteCollectionById) error {
	return CollectionRepository.DeleteCollectionById(collection)
}
