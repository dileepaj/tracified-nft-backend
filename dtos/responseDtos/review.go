package responseDtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResponseReviewStatusUpdate struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Status      string             `json:"status" bson:"status"`
	Rating      float32            `json:"rating" bson:"rating"`
	Description string             `json:"description" bson:"description"`
}
