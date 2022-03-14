package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chart struct {
	ChartTitle string      `json:"charttitle" bson:"charttitle"`
	WeigetId   string      `json:"WeigetId" bson:"weigetid"`
	KeyTitle   string      `json:"Keytitle" bson:"keytitle"`
	ChartData  []ChartData `json:"ChartData" bson:"chartdata"`
	XAxis      string      `json:"Xaxis" bson:"xaxis"`
	YAxis      string      `json:"Yaxis" bson:"yaxis"`
	FontColor  string      `json:"Fontcolor" bson:"fontcolor"`
	FontSize   string      `json:"Fontsize" bson:"fontsize"`
	Width      string      `json:"Width" bson:"width"`
	Height     string      `json:"Height" bson:"height"`
}

type ChartData struct {
	Name   string `json:"Name" bson:"name"`
	Key    string `json:"Key" bson:"key"`
	Value  string `json:"Value" bson:"value"`
	Radius string `json:"Radius" bson:"radius"`
	Color  string `json:"Color" bson:"color"`
}
type Table struct {
	WeigetId     string `json:"WeigetId" bson:"weigetid" validate:"required"`
	TableTitle   string `json:"TableTitle" bson:"tabletitle"`
	TableContent string `json:"TableContent" bson:"Tablecontent" validate:"required"`
}

type StatData struct {
	Key   string `json:"Key" bson:"key"`
	Value string `json:"Value" bson:"value"`
	Color string `json:"Color" bson:"color"`
}

type StataArray struct {
	Title    string     `json:"Title" bson:"title"`
	WeigetId string     `json:"WeigetId" bson:"weigetid" validate:"required"`
	StatData []StatData `json:"StatData" bson:"statdata"`
}

type BotUrl struct {
	Type string   `json:"Type" bson:"type" validate:"required"`
	Urls []string `json:"Urls" bson:"urls" validate:"required"`
}

type BotBatch struct {
	BatchTitle string   `json:"Batchtitle" bson:"batchtitle"`
	Title      string   `json:"Title" bson:"title"`
	BotUrls    []BotUrl `json:"BotUrls" bson:"boturls"`
}
type ProofBotData struct {
	BotTitle string     `json:"BotTitle" bson:"bottitle"`
	WeigetId string     `json:"WeigetId" bson:"weigetid" validate:"required"`
	Batch    []BotBatch `json:"Batch" bson:"batch"`
}

type ImageData struct {
	Title       string `json:"Title" bson:"title"`
	WeigetId    string `json:"WeigetId" bson:"weigetid" validate:"required"`
	Type        string `json:"Type" bson:"type" validate:"required"`
	Base64Image string `json:"Base64Image" bson:"base64image" validate:"required"`
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
	Id          primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	ProjectName string             `json:"ProjectName" bson:"projectname" validate:"required"`
	NFTName     string             `json:"NFTName" bson:"nftname" validate:"required"`
	UserId      string             `json:"UserId" bson:"userid" validate:"required"`
	Timestamp   primitive.DateTime `json:"Timestamp" bson:"timestamp" validate:"required"`
	CreatorName string             `json:"CreatorName" bson:"creatorname"`
	NftContent  NFTContent
}

type Weight struct {
	Id          primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	Timestamp   primitive.DateTime `json:"Timestamp" bson:"timestamp" validate:"required"`
	WeigetId    string             `json:"WeigetId" bson:"weigetid" validate:"required"`
	BactchId    string             `json:"BactchId" bson:"batchid"`
	ProductId   string             `json:"productid" bson:"productid"`
	ProductName string             `json:"ProductId" bson:"productname" validate:"required"`
	TenentName  string             `json:"tenentname" bson:"tenentname"`
	ProjectId   string             `json:"TenentName" bson:"projectid" validate:"required"`
	UserId      string             `json:"UserId" bson:"userid" validate:"required"`
	OTP         string             `json:"OTP" bson:"otp"`
	OTPType     string             `json:"OTPType" bson:"otptype"`
	Query       []string           `json:"Query" bson:"query"`
	WeigetType  string             `json:"WeigetType" bson:"weigettype"`
}
