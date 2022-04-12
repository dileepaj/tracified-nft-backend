package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type NFTCollection struct {
	Id             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId         string             `json:"userid" bson:"userid" validate:"required"`
	Timestamp      primitive.DateTime `json:"timestamp" bson:"timestamp"`
	CollectionName string             `json:"collectionname" bson:"collectionname" validate:"required"`
}