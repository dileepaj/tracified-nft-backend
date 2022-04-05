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
