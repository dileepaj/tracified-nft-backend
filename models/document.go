package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
	TopicID primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Topic   string             `json:"topic" bson:"topic"`
	Answers []string           `json:"answers" bson:"answers"`
}
