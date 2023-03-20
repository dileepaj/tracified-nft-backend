package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tags struct {
	Id        primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	UserId    string             `json:"UserId" bson:"userid" `
	Timestamp primitive.DateTime `json:"Timestamp" bson:"timestamp,omitempty"`
	NFTName   string             `json:"NFTName" bson:"nftName"`
	Tags      []string           `json:"Tags" bson:"tags"`
}
