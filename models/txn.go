package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TXN struct {
	Id            primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	Blockchain    string             `json:"Blockchain" bson:"blockchain" `
	Timestamp     primitive.DateTime `json:"Timestamp" bson:"timestamp,omitempty"`
	NFTIdentifier string             `json:"NFTIdentifier" bson:"nftidentifier"`
	Status        string             `json:"Status" bson:"status"`
	NFTName       string             `json:"NFTName" bson:"nftname"`
	ImageURL      string             `json:"ImageURL" bson:"imageurl"`
	NFTTxnHash    string             `json:"NFTTxnHash" bson:"nfttxnhash"`
}
