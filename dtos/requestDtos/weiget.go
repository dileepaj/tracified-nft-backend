package requestDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type RequestWidget struct {
	Id    primitive.ObjectID `json:"Id" bson:"_id"  validate:"required"`
	Query string             `json:"Query" bson:"query" validate:"required"`
}
