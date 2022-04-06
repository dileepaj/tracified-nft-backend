package responseDtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResponseProject struct {
	Id          primitive.ObjectID `json:"ProjectId" bson:"_id"`
	ProjectName string             `json:"ProjectName" bson:"projectname"`
	Timestamp   primitive.DateTime `json:"Timestamp" bson:"timestamp"`
}

type UpdareProjectResponse struct {
	ProjectId string
	Error     string
}
type WidgetIdResponse struct{
	WidgetId     string `json:"WidgetId" bson:"widgetid"`
}