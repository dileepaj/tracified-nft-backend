package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/models"
)

func CreateWatchList(watchList models.WatchList) (string, error) {
	return watchListRepository.SaveWatchList(watchList)
}

func FindNFTIdentifieryByUserId(userId string) ([]string, error) {
	return watchListRepository.FindNFTIdentifieryByUserId(userId)
}

func GetWatchListByUserPK(userId string) (models.WatchList, error) {
	return watchListRepository.FindWatchListbyUserPK(userId)
}

func GetWatchListsbyBlockchain(blockchain string) ([]models.WatchList, error) {
	return watchListRepository.FindWatchListsByBlockchain("blockchain", blockchain)
}

func GetAllWatchLists() ([]models.Favourite, error) {
	return FavouriteRepository.GetAllFavourites()
}
