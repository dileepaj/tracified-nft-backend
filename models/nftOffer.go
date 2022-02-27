package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Incident struct {
	Type    string `json:"type" bson:"type" validate:"required"`
	Price   string `json:"price" bson:"price" validate:"required,number"`
	SellePK string `json:"seller" bson:"required,seller"`
	BuyerPK string `json:"buyer" bson:"buyer"`
}

type Offer struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NFTIdentifier   string             `json:"nftidentifier" bson:"nftidentifier" validate:"required"`
	Blockchain      string             `json:"blockchain" bson:"blockchain" validate:"required`
	Timestamp       primitive.DateTime `json:"timestamp" bson:"timestamp"`
	Price           string             `json:"currentprice" bson:"currentprice,price"`
	SellingStatus   string             `json:"sellingstatus" bson:"sellingstatus" validate:"required"`
	SellingType     string             `json:"sellingtype" bson:"sellingtype" validate:"required"`
	TransactionHash string             `json:"txnhash" bson:"txnhash" validate:"required,txnhash"`
	IncidentType    Incident
}
