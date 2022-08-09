package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Favourite struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NFTIdentifier string             `json:"nftidentifier" bson:"nftidentifier" `
	Blockchain    string             `json:"blockchain" bson:"blockchain" `
	User          string             `json:"user" bson:"user"`
}
type Hotpicks struct {
	NFTIdentifier string `json:"nftidentifier" bson:"nftidentifier" `
	HotPicks      bool   `json:"hotpicks" bson:"hotpicks" `
}
