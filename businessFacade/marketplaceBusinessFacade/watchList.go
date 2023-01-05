package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateWatchList(watchList models.WatchList) (string, error) {
	return watchListRepository.SaveWatchList(watchList)
}

func FindNFTIdentifieryByUserId(userId string) ([]string, error) {
	return watchListRepository.FindNFTIdentifieryByUserId(userId)
}

func GetWatchListByUserPK(userId string) ([]models.WatchList, error) {
	return watchListRepository.FindWatchListbyUserPK("user", userId)
}

func VerifyWatchListTogglebUserPK(blockchain string, userpk string, nftidentifier string) (models.WatchList, error) {
	return watchListRepository.FindWatchlistbyUserPKandNFTIdentifier(blockchain, userpk, nftidentifier)
}
func RemoveUserFromWatchlist(objectID primitive.ObjectID) (int64, error) {
	return watchListRepository.RemoveUserFromWatchlist(objectID)
}
func FindWatchListsByBlockchainAndIdentifier(blockchain string, id string) ([]models.WatchList, string, error) {
	return watchListRepository.FindWatchListsByBlockchainAndIdentifier("blockchain", blockchain, "nftidentifier", id)
}

func GetAllWatchLists() ([]models.Favourite, error) {
	return FavouriteRepository.GetAllFavourites()
}

func UpdateTrending(nft models.Trending) (models.NFT, error) {
	update := bson.M{
		"$set": bson.M{"trending": nft.Trending},
	}
	return nftRepository.UpdateTrending("nftidentifier", nft.NFTIdentifier, update)
}
