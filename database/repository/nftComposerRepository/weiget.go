package nftComposerRepository

import (
	"context"
	"time"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WidgetRepository struct{}

var Widget = "widget"

// Save the multiple widgets
func (r *WidgetRepository) SaveWidgetList(widgetList []models.Widget) (string, error) {
	rst, err := repository.InsertMany[[]models.Widget](widgetList, Widget)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return "SAVED", nil
	}
}

// Save the widgets return the object Id
func (r *WidgetRepository) SaveWidget(widget models.Widget) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := connections.Connect().Collection(Widget).InsertOne(ctx, widget)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return "", err
	}
	return widget.WidgetId, nil
}

func (r *WidgetRepository) FindWidgetAndUpdate(findBy string,id string,update primitive.M) (models.Widget, error) {
	var widgetResponse models.Widget
	rst := repository.FindOneAndUpdate(findBy, id, update, Widget)
	if rst != nil {
		err := rst.Decode(&widgetResponse)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return widgetResponse, err
		}
		return widgetResponse, nil
	} else {
		return widgetResponse, nil
	}
}

func (r *WidgetRepository) FindWidgetId(idName string, id string) (models.Widget, error) {
	var widget models.Widget
	rst := repository.FindOne[string](idName, id, Widget)
	err := rst.Decode(&widget)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return widget, err
	} else {
		return widget, nil
	}
}
