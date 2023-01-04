package marketplaceRepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *WatchListRepository) FindWatchListbyUserPK(idName string, id string) ([]models.WatchList, error) {
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

func (r *WatchListRepository) FindWatchlistbyUserPKandNFTIdentifier(blockchain string, userpk string, nftID string) (models.WatchList, error) {
	var watchlist models.WatchList
	rst, err := repository.FindById1Id2Id3("blockchain", blockchain, "nftidentifier", nftID, "user", userpk, WatchList)
	if err != nil {
		return watchlist, err
	}
	for rst.Next(context.TODO()) {
		decodeErr := rst.Decode(&watchlist)
		if decodeErr != nil {
			logs.ErrorLogger.Print("failed to decode watchlist data : ", err.Error())
			return watchlist, err
		}
	}
	return watchlist, nil
}

func (r *WatchListRepository) RemoveUserFromWatchlist(objectID primitive.ObjectID) (int64, error) {
	result, err := connections.GetSessionClient(WatchList).DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return 0, err
	}
	return result.DeletedCount, nil
}

func (r *WatchListRepository) GetAllWatchLists() ([]models.WatchList, error) {
	var watchlist []models.WatchList
	findOptions := options.Find()
	result, err := connections.GetSessionClient(WatchList).Find(context.TODO(), bson.D{{}}, findOptions)
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

func (r *WatchListRepository) FindWatchListsByBlockchainAndIdentifier(idName string, id string, idName2 string, id2 string) ([]models.WatchList, string, error) {
	var watchlists []models.WatchList
	rst, err := repository.FindById1AndNotId2(idName, id, idName2, id2, WatchList)
	if err != nil {
		return watchlists, id2, err
	}
	for rst.Next(context.TODO()) {
		var watchlist models.WatchList
		err = rst.Decode(&watchlist)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())

			return watchlists, id2, err
		}
		watchlists = append(watchlists, watchlist)
	}
	return watchlists, id2, nil
}
