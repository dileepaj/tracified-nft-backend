package marketplaceBusinessFacade

import (
	// "fmt"

	// "github.com/dileepaj/tracified-nft-backend/database/repository/marketplaceRepository"
	// "github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	// "github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"

	"log"

	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateFavourites(favs models.Favourite) (string, error) {
	return FavouriteRepository.SaveFavourite(favs)
}

func GetFavouritesByBlockchainAndIdentifier(blockchain string, id string) ([]models.Favourite, string, error) {
	log.Println("params :", blockchain, id)
	return FavouriteRepository.GetFavouritesByBlockchainAndIdentifier("blockchain", blockchain, "nftidentifier", id)
}

func GetFavouritesByUserPK(userid string) (models.Favourite, error) {
	return FavouriteRepository.FindFavouritesbyUserPK(userid)

}

func GetAllFavourites() ([]models.Favourite, error) {
	return FavouriteRepository.GetAllFavourites()
}

func UpdateHotPicks(nft models.Hotpicks) (models.NFT, error) {
	update := bson.M{
		"$set": bson.M{"hotpicks": nft.HotPicks},
	}
	return nftRepository.UpdateHotPicks("nftidentifier", nft.NFTIdentifier, update)
}
