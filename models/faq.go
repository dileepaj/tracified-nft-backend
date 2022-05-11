package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Faq struct {
	QuestionID primitive.ObjectID `json:"questionID" bson:"_id,omitempty"`
	Question   string             `json:"question" bson:"question"`
	Answers    []string           `json:"answers" bson:"answers"`
}
