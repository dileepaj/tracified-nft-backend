package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chart struct {
	WidgetId   string      `json:"WidgetId" bson:"widgetid" validate:"required"`
	ProjectId  string      `json:"ProjectId" bson:"projectid" validate:"required"`
	ChartTitle string      `json:"ChartTitle" bson:"charttitle"`
	ChartData  []ChartData `json:"ChartData" bson:"chartdata"`
	Color      []string    `json:"Color" bson:"color"`
	XAxis      string      `json:"XAxis" bson:"xaxis"`
	YAxis      string      `json:"YAxis" bson:"yaxis"`
	FontColor  string      `json:"FontColor" bson:"fontcolor"`
	FontSize   float32     `json:"FontSize" bson:"fontsize"`
	Width      float32     `json:"Width" bson:"width"`
	Height     float32     `json:"Height" bson:"height"`
	Type       string      `json:"Type" bson:"type" validate:"required"`
	Domain     []float32   `json:"Domain" bson:"domain"`
	ChartImage string      `json:"ChartImage" bson:"chartimage" validate:"required"`
}
type ChartAndWidget struct {
	Chart  Chart  `json:"Chart" bson:"chart"`
	Widget Widget `json:"Widget" bson:"widget"`
}

type TableWithWidget struct {
	Table  Table  `json:"Table" bson:"table"`
	Widget Widget `json:"Widget" bson:"widget"`
}

type BotWithWidget struct {
	ProofBot ProofBotData `json:"ProofBot" bson:"proofbot"`
	Widget   Widget       `json:"Widget" bson:"widget"`
}
type ChartData struct {
	Name  string  `json:"Name" bson:"name"`
	X     float32 `json:"X" bson:"x"`
	Value float32 `json:"Value" bson:"value"`
	Y     float32 `json:"Y" bson:"y"`
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

type ProofURL struct {
	Type string `json:"Type" bson:"type"`
	Urls string `json:"Url" bson:"url"`
}

type ProofData struct {
	BatchId           string   `json:"BatchId" bson:"batchid" `
	GatewayIdentifier string   `json:"GatewayIdentifier" bson:"gatewayidentifier" validate:"required"`
	TxnType           string   `json:"TxnType" bson:"txntype"`
	TxnHash           string   `json:"TxnHash" bson:"txnhash"`
	AvailableProofs   []string `json:"AvailableProofs" bson:"availableproofs"`
	Urls              []ProofURL
	Timestamp         string
}
type ProofBotData struct {
	WidgetId    string             `json:"WidgetId" bson:"widgetid" validate:"required"`
	ProjectId   string             `json:"ProjectId" bson:"projectid" validate:"required"`
	Timestamp   primitive.DateTime `json:"Timestamp" bson:"timestamp" validate:"required"`
	ArtifactId  string             `json:"ArtifactId" bson:"artifactid"`
	ProductId   string             `json:"ProductId" bson:"productid"`
	ProductName string             `json:"ProductName" bson:"productname"`
	TenentId    string             `json:"TenentId" bson:"tenentid" validate:"required"`
	OTPType     string             `json:"OTPType" bson:"otptype"`
	WidgetType  string             `json:"WidgetType" bson:"widgettype" validate:"required"`
	Title       string             `json:"Title" bson:"title"`
	Data        []ProofData
}

type ImageData struct {
	WidgetId    string `json:"WidgetId" bson:"widgetid" validate:"required"`
	ProjectId   string `json:"ProjectId" bson:"projectid" validate:"required"`
	Title       string `json:"Title" bson:"title"`
	Type        string `json:"Type" bson:"type" validate:"required"`
	Base64Image string `json:"Base64Image" bson:"base64image" validate:"required"`
}

type ContentOrderData struct {
	WidgetId string `json:"WidgetId" bson:"widgetid" validate:"required"`
	Type     string `json:"Type" bson:"type" validate:"required"`
}

type NFTContent struct {
	BarCharts    []Chart        `json:"BarCharts" bson:"barcharts"`
	PieCharts    []Chart        `json:"PieCharts" bson:"piecharts"`
	BubbleCharts []Chart        `json:"BubbleCharts" bson:"bubblecharts"`
	Stats        []StataArray   `json:"Stats" bson:"stats"`
	Tables       []Table        `json:"Tables" bson:"tables"`
	Images       []ImageData    `json:"Images" bson:"images"`
	ProofBot     []ProofBotData `json:"ProofBot" bson:"proofbot"`
	TimeLine     []Timeline     `json:"Timeline" bson:"timeline"`
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
	Description      string             `json:"Description" bson:"description" validate:"required"`
	UserId           string             `json:"UserId" bson:"userid" validate:"required"`
	TenentId         string             `json:"TenentId" bson:"tenentid" validate:"required"`
	TenentName       string             `json:"TenentName" bson:"tenentname"`
	Timestamp        primitive.DateTime `json:"Timestamp" bson:"timestamp" validate:"required"`
	CreatorName      string             `json:"CreatorName" bson:"creatorname"`
	ContentOrderData []ContentOrderData `json:"ContentOrderData" bson:"Contentorderdata" `
}

// pie.bar.,bubble
type Widget struct {
	Id          primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	ProjectId   string             `json:"ProjectId" bson:"projectid" validate:"required"`
	WidgetId    string             `json:"WidgetId" bson:"widgetid" validate:"required"`
	ArtifactId  string             `json:"ArtifactId" bson:"artifactid"`
	Timestamp   primitive.DateTime `json:"Timestamp" bson:"timestamp" validate:"required"`
	BatchId     string             `json:"BatchId" bson:"bathid"`
	ProductId   string             `json:"ProductId" bson:"productid"`
	ProductName string             `json:"ProductName" bson:"productname"`
	TenentId    string             `json:"TenentId" bson:"tenentid" validate:"required"`
	UserId      string             `json:"UserId" bson:"userid" validate:"required"`
	OTP         string             `json:"OTP" bson:"otp"`
	OTPType     string             `json:"OTPType" bson:"otptype"`
	Query       string             `json:"Query" bson:"query"`
	WidgetType  string             `json:"WidgetType" bson:"widgettype"`
}

type ProjectDetail struct {
	Project      NFTComposerProject
	BarCharts    []ChartAndWidget  `json:"BarCharts" bson:"barcharts"`
	PieCharts    []ChartAndWidget  `json:"PieCharts" bson:"piecharts"`
	BubbleCharts []ChartAndWidget  `json:"BubbleCharts" bson:"bubblecharts"`
	Stats        []StataArray      `json:"Stats" bson:"stats"`
	Tables       []TableWithWidget `json:"Tables" bson:"tables"`
	Images       []ImageData       `json:"Images" bson:"images"`
	ProofBot     []ProofBotData    `json:"ProofBot" bson:"proofbot"`
	Timeline     []Timeline        `json:"Timeline" bson:"timeline"`
}

type Timeline struct {
	Id           primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	ProjectId    string             `json:"ProjectId" bson:"projectid" validate:"required"`
	BatchId      string             `json:"BatchId" bson:"batchid"`
	WidgetId     string             `json:"WidgetId" bson:"widgetid" validate:"required"`
	ArtifactId   string             `json:"ArtifactId" bson:"artifactid"`
	Timestamp    primitive.DateTime `json:"Timestamp" bson:"timestamp" validate:"required"`
	ProductId    string             `json:"ProductId" bson:"productid"` // item id
	ProductName  string             `json:"ProductName" bson:"productname"`
	Title        string             `json:"Title" bson:"title"`
	TimelineData []TimelineData
	WidgetType   string `json:"WidgetType" bson:"widgettype"`
}

type TimelineData struct {
	Title       string   `json:"Title" bson:"title"`
	Icon        string   `json:"Icon" bson:"icon" `
	SubTitle    string   `json:"SubTitle" bson:"subtitle" `
	Description string   `json:"Description" bson:"description"`
	Images      []string `json:"Images" bson:"images"`
	Children    []Children
}
type Children struct {
	NewTDP    bool   `json:"NewTDP" bson:"newtdp"`
	Timestamp string `json:"Timestamp" bson:"timestamp"`
	Key       string `json:"Key" bson:"Key" `
	Value     string `json:"Value" bson:"value" `
}
