package nftComposerBusinessFacade

import (
	nftcomposerrepository "github.com/dileepaj/tracified-nft-backend/database/repository/nftComposerRepository"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
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

func FindWidgetById(id string)(models.Widget,error){
	return widgetRepository.FindWidgetId("widgetid",id)
}