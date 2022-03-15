package nftcomposercontroller

import (
	nftcomposerrepository "github.com/dileepaj/tracified-nft-backend/database/repository/nftComposerRepository"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var widgetRepository nftcomposerrepository.WidgetRepository

func SaveWidgetList(widgets []models.Widget) (string, error) {
	return widgetRepository.SaveWidgetList(widgets)
}

func SaveWidget(widget models.Widget) (string, error) {
	return widgetRepository.SaveWidget(widget)
}

func FindWidgetAndUpdateQuery(widget requestDtos.RequestWidget)(models.Widget,error){
	return widgetRepository.FindWidgetAndUpdate(widget)
}

func FindWidgetById(id primitive.ObjectID)(models.Widget,error){
	return widgetRepository.FindWidgetById(id)
}