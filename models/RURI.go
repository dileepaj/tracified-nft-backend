package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SvgCreator struct {
	Id          primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	ProductName string             `json:"productname" bson:"productname" `
}
