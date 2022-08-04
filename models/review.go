package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NFTIdentifier string             `json:"nftidentifier" bson:"nftidentifier" validate:"required"`
	UserID        string             `json:"userid" bson:"userid" validate:"required"`
	Status        string             `json:"status" bson:"status"`
	Rating        float32            `json:"rating" bson:"rating"`
	Description   string             `json:"description" bson:"description"`
}
