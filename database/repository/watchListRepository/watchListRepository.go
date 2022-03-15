package watchListRepository

import (
	"context"
	"fmt"
	"time"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WatchListRepository struct{}

func (repository *WatchListRepository) FindNFTIdentifieryByUserId(userId string) ([]string, error) {
	var watchLists []string
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	findOptions.SetProjection(bson.M{"nftidentifier": 1, "_id": 0})
	rst, err := connections.Connect().Collection("watchlist").Find(context.TODO(), bson.D{{"list", userId}}, findOptions)
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

func (repository *WatchListRepository) Save(watchList models.WatchList) (string, error) {
	fmt.Println(watchList)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rst, err := connections.Connect().Collection("watchlist").InsertOne(ctx, watchList)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return watchList.NFTIdentifier, err
	}
	var id = rst.InsertedID.(primitive.ObjectID)
	return id.String(), nil
}
