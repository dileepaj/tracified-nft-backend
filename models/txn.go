package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TXN struct {
	Id            primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	Blockchain    string             `json:"Blockchain" bson:"blockchain" `
	NFTIdentifier string             `json:"NFTIdentifier" bson:"nftidentifier"`
	Status        string             `json:"Status" bson:"status"`
	NFTName       string             `json:"NFTName" bson:"nftname"`
	ImageURL      string             `json:"ImageURL" bson:"imageurl"`
	NFTTxnHash    string             `json:"NFTTxnHash" bson:"nfttxnhash"`
	Time          string             `json:"Time" bson:"Time"`
}
