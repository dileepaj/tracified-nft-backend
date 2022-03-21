package nftComposerBusinessFacade

import (
	nftcomposerrepository "github.com/dileepaj/tracified-nft-backend/database/repository/nftComposerRepository"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

var widgetRepository nftcomposerrepository.WidgetRepository

func SaveWidgetList(widgets []models.Widget) (string, error) {
	return widgetRepository.SaveWidgetList(widgets)
}

func SaveWidget(widget models.Widget) (string, error) {
	return widgetRepository.SaveWidget(widget)
}

func FindWidgetAndUpdateQuery(widget requestDtos.RequestWidget)(models.Widget,error){
		update := bson.M{
		"$set": bson.M{"query": widget.Query},
	}
	return widgetRepository.FindWidgetAndUpdate("widgetid",widget.WidgetId,update)
}

func FindWidgetById(id string)(models.Widget,error){
	return widgetRepository.FindWidgetId("widgetid",id)
}