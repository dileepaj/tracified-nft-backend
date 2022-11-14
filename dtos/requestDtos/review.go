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

type ReviewFiltering struct {
	Filterby      string `json:"filterby" bson:"filterby"`
	FilterType    int    `json:"filtertype" bson:"filtertype"`
	NFTIdentifier string `json:"nftidentifier" bson:"nftidentifier" validate:"required"`
	PageSize      int32  `json:"pagesize" bson:"pagesize"`
	RequestedPage int32  `json:"requestedPage" bson:"requestedPage" `
}
