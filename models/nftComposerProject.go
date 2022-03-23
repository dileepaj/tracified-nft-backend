package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chart struct {
	WidgetId   string      `json:"WidgetId" bson:"widgetid" validate:"required"`
	ProjectId  string      `json:"ProjectId" bson:"projectid" validate:"required"`
	ChartTitle string      `json:"ChartTitle" bson:"charttitle"`
	KeyTitle   string      `json:"Keytitle" bson:"keytitle"`
	ChartData  []ChartData `json:"ChartData" bson:"chartdata"`
	XAxis      string      `json:"Xaxis" bson:"xaxis"`
	YAxis      string      `json:"Yaxis" bson:"yaxis"`
	FontColor  string      `json:"Fontcolor" bson:"fontcolor"`
	FontSize   string      `json:"Fontsize" bson:"fontsize"`
	Width      string      `json:"Width" bson:"width"`
	Height     string      `json:"Height" bson:"height"`
	Type       string      `json:"Type" bson:"type" validate:"required"`
}
type ChartAndWidget struct {
	Chart  Chart  `json:"Chart" bson:"chart"`
	Widget Widget `json:"Widget" bson:"widget"`
}

type TableWithWidget struct {
	Table Table  `json:"Table" bson:"table"`
	Widget Widget `json:"Widget" bson:"widget"`
}

type BotWithWidget struct {
	ProofBot ProofBotData `json:"ProofBot" bson:"proofbot"`
	Widget   Widget       `json:"Widget" bson:"widget"`
}
type ChartData struct {
	Name   string `json:"Name" bson:"name"`
	Key    string `json:"Key" bson:"key"`
	Value  string `json:"Value" bson:"value"`
	Radius string `json:"Radius" bson:"radius"`
	Color  string `json:"Color" bson:"color"`
}
type Table struct {
	WidgetId     string `json:"WidgetId" bson:"widgetid" validate:"required"`
	ProjectId    string `json:"ProjectId" bson:"projectid" validate:"required"`
	TableTitle   string `json:"TableTitle" bson:"tabletitle"`
	TableContent string `json:"TableContent" bson:"Tablecontent" validate:"required"`
}

type StatData struct {
	Key   string `json:"Key" bson:"key"`
	Value string `json:"Value" bson:"value"`
	Color string `json:"Color" bson:"color"`
}

type StataArray struct {
	WidgetId  string     `json:"WidgetId" bson:"widgetid" validate:"required"`
	ProjectId string     `json:"ProjectId" bson:"projectid" validate:"required"`
	Title     string     `json:"Title" bson:"title"`
	StatData  []StatData `json:"StatData" bson:"statdata"`
}

type BotUrl struct {
	Type string   `json:"Type" bson:"type" validate:"required"`
	Urls []string `json:"Urls" bson:"urls" validate:"required"`
}

type BotBatch struct {
	BatchTitle  string   `json:"BatchTitle" bson:"batchtitle"`
	Title       string   `json:"Title" bson:"title"`
	BatchId     string   `json:"BatchId" bson:"batchid" `
	TenentId    string   `json:"TenentId" bson:"tenentid"`
	ProductId   string   `json:"productId" bson:"productid"`
	ProductName string   `json:"ProductName" bson:"productname"`
	BotUrls     []BotUrl `json:"BotUrls" bson:"boturls"`
}
type ProofBotData struct {
	WidgetId  string             `json:"WidgetId" bson:"widgetid" validate:"required"`
	ProjectId string             `json:"ProjectId" bson:"projectid" validate:"required"`
	BotTitle  string             `json:"BotTitle" bson:"bottitle"`
	Timestamp primitive.DateTime `json:"Timestamp" bson:"timestamp" validate:"required"`
	NFTType   string             `json:"NFTType" bson:"nfttype" validate:"required"`
	Batch     []BotBatch         `json:"Batch" bson:"batch"`
}

type ImageData struct {
	WidgetId    string `json:"WidgetId" bson:"widgetid" validate:"required"`
	ProjectId   string `json:"ProjectId" bson:"projectid" validate:"required"`
	Title       string `json:"Title" bson:"title"`
	Type        string `json:"Type" bson:"type" validate:"required"`
	Base64Image string `json:"Base64Image" bson:"base64image" validate:"required"`
}

type ContentOrderData struct {
	WidgetId        string `json:"WidgetId" bson:"widgetid" validate:"required"`
	Type            string `json:"Type" bson:"type" validate:"required"`
	CardOrderNumber int    `json:"cardOrderNumber" bson:"cardordernumber"`
}

type NFTContent struct {
	BarCharts    []Chart        `json:"BarCharts" bson:"barcharts"`
	PieCharts    []Chart        `json:"PieCharts" bson:"piecharts"`
	BubbleCharts []Chart        `json:"BubbleCharts" bson:"bubblecharts"`
	Stats        []StataArray   `json:"Stats" bson:"stats"`
	Tables       []Table        `json:"Tables" bson:"tables"`
	Images       []ImageData    `json:"Images" bson:"images"`
	ProofBot     []ProofBotData `json:"ProofBot" bson:"proofbot"`
}

type HtmlGenerator struct {
	Id               primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	ProjectId        string             `json:"ProjectId" bson:"projectid" validate:"required"`
	ProjectName      string             `json:"ProjectName" bson:"projectname" validate:"required"`
	NFTName          string             `json:"NFTName" bson:"nftname" validate:"required"`
	UserId           string             `json:"UserId" bson:"userid" validate:"required"`
	TenentId         string             `json:"TenentId" bson:"tenentid" validate:"required"`
	TenentName       string             `json:"TenentName" bson:"tenentname"`
	Timestamp        primitive.DateTime `json:"Timestamp" bson:"timestamp" validate:"required"`
	CreatorName      string             `json:"CreatorName" bson:"creatorname"`
	ContentOrderData []ContentOrderData `json:"ContentOrderData" bson:"Contentorderdata" validate:"required"`
	NftContent       NFTContent         `json:"NftContent" bson:"nftcontent" validate:"required"`
}
type NFTComposerProject struct {
	Id               primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	ProjectId        string             `json:"ProjectId" bson:"projectid" validate:"required"`
	ProjectName      string             `json:"ProjectName" bson:"projectname" validate:"required"`
	NFTName          string             `json:"NFTName" bson:"nftname" validate:"required"`
	UserId           string             `json:"UserId" bson:"userid" validate:"required"`
	TenentId         string             `json:"TenentId" bson:"tenentid" validate:"required"`
	TenentName       string             `json:"TenentName" bson:"tenentname"`
	Timestamp        primitive.DateTime `json:"Timestamp" bson:"timestamp" validate:"required"`
	CreatorName      string             `json:"CreatorName" bson:"creatorname"`
	ContentOrderData []ContentOrderData `json:"ContentOrderData" bson:"Contentorderdata" validate:"required"`
}

// pie.bar.,bubble.Table
type Widget struct {
	Id          primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	ProjectId   string             `json:"ProjectId" bson:"projectid" validate:"required"`
	WidgetId    string             `json:"WidgetId" bson:"widgetid" validate:"required"`
	ArtifactId  string             `json:"ArtifactId" bson:"artifactid"`
	Timestamp   primitive.DateTime `json:"Timestamp" bson:"timestamp" validate:"required"`
	BatchId     string             `json:"BatchId" bson:"bathid"`
	ProductId   string             `json:"productId" bson:"productid"`
	ProductName string             `json:"productName" bson:"productname"`
	TenentId    string             `json:"TenentId" bson:"tenentid" validate:"required"`
	UserId      string             `json:"UserId" bson:"userid" validate:"required"`
	OTP         string             `json:"OTP" bson:"otp"`
	OTPType     string             `json:"OTPType" bson:"otptype"`
	Query       string             `json:"Query" bson:"query"`
	WidgetType  string             `json:"WidgetType" bson:"widgettype"`
}

type ProjectDetail struct {
	Project      NFTComposerProject
	BarCharts    []ChartAndWidget `json:"BarCharts" bson:"barcharts"`
	PieCharts    []ChartAndWidget `json:"PieCharts" bson:"piecharts"`
	BubbleCharts []ChartAndWidget `json:"BubbleCharts" bson:"bubblecharts"`
	Stats        []StataArray     `json:"Stats" bson:"stats"`
	Tables       []TableWithWidget          `json:"Tables" bson:"tables"`
	Images       []ImageData      `json:"Images" bson:"images"`
	ProofBot     []BotWithWidget   `json:"ProofBot" bson:"proofbot"`
}
