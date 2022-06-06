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

func (r *WatchListRepository) FindWatchListbyUserPK(userpk string) (models.WatchList, error) {
	var watchList models.WatchList

	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	rst, err := session.Client().Database(connections.DbName).Collection("watchlist").Find(context.TODO(), bson.M{"creatoruserid": userpk})
	if err != nil {
		return watchList, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&watchList)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection watchlist in GetWatchlistByID:watchlistRepository.go: ", err.Error())
			return watchList, err
		}
	}
	return watchList, err
}

func (r *WatchListRepository) GetAllWatchLists() ([]models.WatchList, error) {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session in getAllWatchList : WatchRepository.go : ", err.Error())
	}
	defer session.EndSession(context.TODO())

	var watchlist []models.WatchList
	findOptions := options.Find()
	findOptions.SetLimit(10)
	result, err := session.Client().Database(connections.DbName).Collection(WatchList).Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println("Error occured when trying to connect to DB and excute Find query in GetAllWatchList:watchlistRepository.go: ", err.Error())
		return watchlist, err
	}
	for result.Next(context.TODO()) {
		var watchlists models.WatchList
		err = result.Decode(&watchlists)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection faq in GetAllWatchlist:watchRepository.go: ", err.Error())
			return watchlist, err
		}
		watchlist = append(watchlist, watchlists)
	}
	return watchlist, nil
}

func (r *WatchListRepository) FindWatchListsByBlockchain(idName string, id string) ([]models.WatchList, error) {
	var watchlists []models.WatchList
	rst, err := repository.FindById(idName, id, WatchList)
	if err != nil {
		return watchlists, err
	}
	for rst.Next(context.TODO()) {
		var watchlist models.WatchList
		err = rst.Decode(&watchlist)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return watchlists, err
		}
		watchlists = append(watchlists, watchlist)
	}
	return watchlists, nil
}
