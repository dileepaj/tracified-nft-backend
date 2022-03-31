package marketplaceRepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

type WatchListRepository struct{}

var WatchList = "watchlist"

func (r *WatchListRepository) FindNFTIdentifieryByUserId(userId string) ([]string, error) {
	var watchLists []string
	rst, err := repository.FindById("list", userId, WatchList)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return watchLists, err
	}
	for rst.Next(context.TODO()) {
		var watchList models.WatchList
		err = rst.Decode(&watchList)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return watchLists, err
		}
		watchLists = append(watchLists, watchList.NFTIdentifier)
	}
	return watchLists, nil
}

func (r *WatchListRepository) SaveWatchList(watchList models.WatchList) (string, error) {
	return repository.Save[models.WatchList](watchList, WatchList)
}
