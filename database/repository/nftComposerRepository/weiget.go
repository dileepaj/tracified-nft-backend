package nftComposerRepository

import (
	"context"
	"time"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *WidgetRepository) FindWidgetAndUpdate(findBy string, id string, update primitive.M) (models.Widget, error) {
	var widgetResponse models.Widget
	projection := bson.M{"otp": 0}
	rst := repository.FindOneAndUpdate(findBy, id, update, projection, Widget)
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

func (r *WidgetRepository) FindWidgetOneByIdWithOtp(idName string, id string) (models.Widget, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var widget models.Widget
	rst := connections.Connect().Collection(Widget).FindOne(ctx, bson.D{{idName, id}})
	err := rst.Decode(&widget)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return models.Widget{}, err
	} else {
		return widget, nil
	}
}

func (r *WidgetRepository) FindWidgetOneById(idName string, id string) (models.Widget, error) {
	var widget models.Widget
	rst := repository.FindOne[string](idName, id, Widget)
	err := rst.Decode(&widget)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return models.Widget{}, err
	} else {
		return widget, nil
	}
}

func (r *NFTComposerProjectRepository) FindWidgetsById(idName string, id string) ([]models.Widget, error) {
	var widgets []models.Widget
	rst, err := repository.FindById(idName, id, Widget)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return widgets, err
	}
	for rst.Next(context.TODO()) {
		var widgetResult models.Widget
		err = rst.Decode(&widgetResult)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return widgets, err
		}
		widgets = append(widgets, widgetResult)
	}
	return widgets, nil
}
