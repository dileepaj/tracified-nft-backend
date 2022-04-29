package requestDtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateReviewStatus struct {
	Id     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Status string             `json:"status" bson:"status"`
}
type DeleteReview struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}
