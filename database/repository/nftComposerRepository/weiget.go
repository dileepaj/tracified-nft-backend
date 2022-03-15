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

type WidgetRepository struct{}

//Save the multiple widgets
func (r *WidgetRepository) SaveWidgetList(widgetList []models.Widget) (string, error) {
	var docs []interface{}
	for _, t := range widgetList {
		docs = append(docs, t)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := connections.Connect().Collection("widget").InsertMany(ctx, docs)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return "", err
	} else {
		return "SAVED", nil
	}
}

//Save the widgets return the object Id
func (r *WidgetRepository) SaveWidget(widget models.Widget) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := connections.Connect().Collection("widget").InsertOne(ctx, widget)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return widget.WidgetId, err
	} else {
		id := result.InsertedID.(primitive.ObjectID).Hex()
		return id, nil
	}
}

func (r *WidgetRepository) FindWidgetAndUpdate(widget requestDtos.RequestWidget) (models.Widget, error) {
	var widgetResponse models.Widget
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	update := bson.M{
		"$set": bson.M{"query": widget.Query},
	}
	err := connections.Connect().Collection("widget").FindOneAndUpdate(ctx, bson.M{"_id": widget.Id}, update, &opt).Decode(&widgetResponse)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return widgetResponse, err
	} else {
		return widgetResponse, nil
	}
}

func (r *WidgetRepository) FindWidgetById(id primitive.ObjectID) (models.Widget, error) {
	var widget models.Widget
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := connections.Connect().Collection("widget").FindOne(ctx, bson.M{"_id": id}).Decode(&widget)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return widget, err
	} else {
		return widget, nil
	}

}
