package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SVG struct {
	Id             primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	Base64ImageSVG string             `json:"Base64ImageSVG" bson:"base64imagesvg" `
	Timestamp      primitive.DateTime `json:"Timestamp" bson:"timestamp,omitempty"`
	Hash           string             `json:"Hash" bson:"hash"`
}
