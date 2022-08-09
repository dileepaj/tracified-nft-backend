package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
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
