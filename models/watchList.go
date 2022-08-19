package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type WatchList struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NFTIdentifier string             `json:"nftidentifier" bson:"nftidentifier" `
	Blockchain    string             `json:"blockchain" bson:"blockchain" `
	User          string             `json:"user" bson:"user" `
}
type Trending struct {
	NFTIdentifier string `json:"nftidentifier" bson:"nftidentifier" `
	Trending      bool   `json:"trending" bson:"trending" `
}
