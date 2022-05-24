package marketplaceRepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (repository *WatchListRepository) FindWatchListbyUserPK(idName1 string, id1 string) ([]models.WatchList, error) {
	var watchlists []models.WatchList
	if idName1 != "" {
		findOptions := options.Find()
		rst, err := connections.Connect().Collection("watchlist").Find(context.TODO(), bson.D{{idName1, id1}}, findOptions)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return watchlists, err
		}
		for rst.Next(context.TODO()) {
			var watchlist models.WatchList
			err = rst.Decode((&watchlist))
			if err != nil {
				logs.ErrorLogger.Println(err.Error())
				return watchlists, err
			}
			watchlists = append(watchlists, watchlist)
		}
		return watchlists, nil
	} else {
		return watchlists, nil
	}
}

func (repository *WatchListRepository) GetAllWatchLists() ([]models.WatchList, error) {
	var watchlists []models.WatchList
	findOptions := options.Find()
	findOptions.SetLimit(10)
	rst, err := connections.Connect().Collection("watchlist").Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return watchlists, err
	}
	for rst.Next(context.TODO()) {
		var watchlist models.WatchList
		err = rst.Decode((&watchlist))
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return watchlists, err
		}
		watchlists = append(watchlists, watchlist)
	}
	return watchlists, nil
}
