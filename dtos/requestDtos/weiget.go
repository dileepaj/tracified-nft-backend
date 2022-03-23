package requestDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type RequestWidget struct {
	WidgetId string `json:"widgetId" bson:"widgetid"  validate:"required"`
	Query    string `json:"Query" bson:"query" validate:"required"`
}

type UpdateWidgetRequest struct {
	WidgetId    string             `json:"widgetId" bson:"widgetid"  validate:"required"`
	Timestamp   primitive.DateTime `json:"Timestamp" bson:"timestamp" validate:"required"`
	BatchId     string             `json:"BatchId" bson:"bathid"`
	ProductId   string             `json:"productId" bson:"productid"`
	ProductName string             `json:"productName" bson:"productname"`
	TenentId    string             `json:"TenentId" bson:"tenentid" validate:"required"`
	OTPType     string             `json:"OTPType" bson:"otptype" validate:"required"`
	ArtifactId  string             `json:"ArtifactId" bson:"artifactid"`
	WidgetType  string             `json:"WidgetType" bson:"widgettype" validate:"required"`
}
