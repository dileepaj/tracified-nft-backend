package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateFavourites(favs models.Favourite) (string, error) {
	return FavouriteRepository.SaveFavourite(favs)
}

func GetFavouritesByBlockchainAndIdentifier(blockchain string, id string) ([]models.Favourite, string, error) {
	return FavouriteRepository.GetFavouritesByBlockchainAndIdentifier("blockchain", blockchain, "nftidentifier", id)
}

func GetFavouritesByUserPK(userid string) ([]models.Favourite, error) {
	return FavouriteRepository.FindFavouritesbyUserPK("user", userid)

}

func VerifyFavouriteTogglebUserPK(blockchain string, userpk string, nftidentifier string) (models.Favourite, error) {
	return FavouriteRepository.FindFavouritesbyUserPKandNFTIdentifier(blockchain, userpk, nftidentifier)
}
func RemoveUserFromFavourites(objectID primitive.ObjectID) (int64, error) {
	return FavouriteRepository.RemoveUserFromFavourites(objectID)
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
