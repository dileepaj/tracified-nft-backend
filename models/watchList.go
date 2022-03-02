package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type WatchList struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NFTIdentifier string             `json:"nftidentifier" bson:"nftidentifier" validate:"required"`
	List          []string           `json:"list" bson:"list"`
}
