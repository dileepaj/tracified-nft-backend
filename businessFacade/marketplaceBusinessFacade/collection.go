package marketplaceBusinessFacade

import (
	"fmt"
	"log"

	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
)

func CreateCollection(collection models.NFTCollection) (string, error) {
	log.Println("------------------------------------testing 7 ---------------------------------------------------")
	return CollectionRepository.SaveCollection(collection)
}
func GetAllCollections() ([]models.NFTCollection, error) {
	fmt.Println("Calling repo...")
	return CollectionRepository.GetAllCollections()
}
func GetCollectionByUserPK(userid string) ([]models.NFTCollection, error) {
	log.Println("---------------------------------------test 2------------------------------", userid)
	return CollectionRepository.FindCollectionbyUserPK("userid", userid)

}

func GetCollectionById(_id string) ([]models.NFTCollection, error) {
	log.Println("---------------------------------------test 2------------------------------", _id)
	return CollectionRepository.FindCollectionbyId(_id)
}

func UpdateCollection(update requestDtos.UpdateCollection) (responseDtos.ResponseCollectionUpdate, error) {
	return CollectionRepository.UpdateCollection(update)
}
func DeleteCollectionByUserPK(collection requestDtos.DeleteCollectionByUserPK) error {
	log.Println("-----------------------------------------test 4----------------------------------------")
	return CollectionRepository.DeleteCollectionByUserPK(collection)
}

func DeleteCollectionById(collection requestDtos.DeleteCollectionById) error {
	log.Println("-----------------------------------------test 5----------------------------------------")
	return CollectionRepository.DeleteCollectionById(collection)
}
