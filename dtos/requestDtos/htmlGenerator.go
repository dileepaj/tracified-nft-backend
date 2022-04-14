package requestDtos

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HtmlGeneratorRequest struct {
	NFTComposerProject models.NFTComposerProject
	WidgetDetails      []models.Widget
}

type UpdateProjectRequest struct {
	ProjectId        string                    `json:"ProjectId" bson:"projectid" validate:"required"`
	ProjectName      string                    `json:"ProjectName" bson:"projectname" validate:"required"`
	NFTName          string                    `json:"NFTName" bson:"nftname" validate:"required"`
	TenentId         string                    `json:"TenentId" bson:"tenentid" validate:"required"`
	TenentName       string                    `json:"TenentName" bson:"tenentname"`
	Timestamp        primitive.DateTime        `json:"Timestamp" bson:"timestamp" validate:"required"`
	CreatorName      string                    `json:"CreatorName" bson:"creatorname"`
	ContentOrderData []models.ContentOrderData `json:"ContentOrderData" bson:"Contentorderdata" validate:"required"`
}
type UpdateChartRequest struct {
	WidgetId   string             `json:"WidgetId" bson:"widgetid" validate:"required"`
	ChartTitle string             `json:"charttitle" bson:"charttitle"`
	ChartData  []models.ChartData `json:"ChartData" bson:"chartdata"`
	XAxis      string             `json:"Xaxis" bson:"xaxis"`
	YAxis      string             `json:"Yaxis" bson:"yaxis"`
	FontColor  string             `json:"Fontcolor" bson:"fontcolor"`
	FontSize   float32             `json:"Fontsize" bson:"fontsize"`
	Width      float32             `json:"Width" bson:"width"`
	Height     float32             `json:"Height" bson:"height"`
}
type UpdateTableRequest struct {
	WidgetId     string `json:"WidgetId" bson:"widgetid" validate:"required"`
	TableTitle   string `json:"TableTitle" bson:"tabletitle"`
	TableContent string `json:"TableContent" bson:"Tablecontent" validate:"required"`
}
type UpdateStatsRequest struct {
	WidgetId string            `json:"WidgetId" bson:"widgetid" validate:"required"`
	Title    string            `json:"Title" bson:"title"`
	StatData []models.StatData `json:"StatData" bson:"statdata"`
}
type UpdateImageRequest struct {
	WidgetId    string `json:"WidgetId" bson:"widgetid" validate:"required"`
	Title       string `json:"Title" bson:"title"`
	Type        string `json:"Type" bson:"type" validate:"required"`
	Base64Image string `json:"Base64Image" bson:"base64image" validate:"required"`
}
type UpdateProofBotRequest struct {
	WidgetId    string             `json:"WidgetId" bson:"widgetid" validate:"required"`
	Timestamp   primitive.DateTime `json:"Timestamp" bson:"timestamp" validate:"required"`
	ArtifactId  string             `json:"ArtifactId" bson:"artifactid"`
	ProductId   string             `json:"ProductId" bson:"productid"`
	ProductName string             `json:"ProductName" bson:"productname"`
	TenentId    string             `json:"TenentId" bson:"tenentid" validate:"required"`
	OTPType     string             `json:"OTPType" bson:"otptype"`
	WidgetType  string             `json:"WidgetType" bson:"widgettype"`
	Data        []models.ProofData
}

type UpdateTimelineRequest struct {
	WidgetId     string             `json:"WidgetId" bson:"widgetid" validate:"required"`
	ArtifactId   string             `json:"ArtifactId" bson:"artifactid"`
	Timestamp    primitive.DateTime `json:"Timestamp" bson:"timestamp" validate:"required"`
	ProductId    string             `json:"productId" bson:"productid"` // item id
	ProductName  string             `json:"productName" bson:"productname"`
	TimelineData []models.TimelineData
}