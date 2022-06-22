package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Partner struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Topic       string             `json:"topic" bson:"topic"`
	WebLink     string             `json:"weblink" bson:"weblink"`
	CompanyName string             `json:"companyname" bson:"companyname"`
	Image       string             `json:"image" bson:"image"`
	Description string             `json:"desc" bson:"desc"`
}
