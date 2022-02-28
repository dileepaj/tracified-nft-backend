package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Favourite struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NFTIdentifier string             `json:"nftidentifier" bson:"nftidentifier" validate:"required"`
	Blockchain    string             `json:"blockchain" bson:"blockchain" validate:"required"`
	List          []string           `json:"list" bson:"list"`
}
