package nftcomposerrepository

import (
	"context"
	"time"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WeigetRepository struct{}

//Save the multiple weigets
func (r *WeigetRepository) SaveWeigetList(weigetList []models.Weiget) (string, error) {
	var docs []interface{}
	for _, t := range weigetList {
		docs = append(docs, t)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := connections.Connect().Collection("weiget").InsertMany(ctx, docs)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return "", err
	} else {
		return "SAVED", nil
	}
}

//Save the weigets return the object Id
func (r *WeigetRepository) SaveWeiget(weiget models.Weiget) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := connections.Connect().Collection("weiget").InsertOne(ctx, weiget)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return weiget.WeigetId, err
	} else {
		id := result.InsertedID.(primitive.ObjectID).Hex()
		return id, nil
	}
}

func (r *WeigetRepository) FindWeigetAndUpdate(weiget requestDtos.RequestWeiget) (models.Weiget, error) {
	var weigetResponse models.Weiget
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	update := bson.M{
		"$set": bson.M{"query": weiget.Query},
	}
	err := connections.Connect().Collection("weiget").FindOneAndUpdate(ctx, bson.M{"_id": weiget.Id}, update, &opt).Decode(&weigetResponse)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return weigetResponse, err
	} else {
		return weigetResponse, nil
	}
}

func (r *WeigetRepository) FindWeigetById(id primitive.ObjectID) (models.Weiget, error) {
	var weiget models.Weiget
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := connections.Connect().Collection("weiget").FindOne(ctx, bson.M{"_id": id}).Decode(&weiget)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return weiget, err
	} else {
		return weiget, nil
	}

}
