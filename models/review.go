package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NFTIdentifier string             `json:"nftidentifier" bson:"nftidentifier" validate:"required"`
	UserID        string             `json:"userid" bson:"userid" validate:"required"`
	Status        string             `json:"status" bson:"status"`
	Rating        float32            `json:"rating" bson:"rating"`
	Description   string             `json:"description" bson:"description"`
	Timestamp     string             `json:"timestamp" bson:"timestamp,omitempty"`
}

type CreatorsList struct {
	NftIdentifier    string
	UserID           string
	Star1Ratings     int
	Star_1_5_Ratings int
	Star_2_Ratings   int
	Star_2_5Ratings  int
	Star_3_Ratings   int
	Star_3_5_Ratings int
	Star_4_Ratings   int
	Star_4_5_Ratings int
	Star_5_Ratings   int
	TotalStars       int
	AvgRating        float32
}

type ReviewsforPagination struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NFTIdentifier string             `json:"nftidentifier" bson:"nftidentifier" validate:"required"`
	UserID        string             `json:"userid" bson:"userid" validate:"required"`
	Rating        float32            `json:"rating" bson:"rating"`
	Description   string             `json:"description" bson:"description"`
}

type ReviewPaginatedResponse struct {
	ReviewContent  []ReviewsforPagination `json:"reviewcontent" bson:"reviewcontent"`
	PaginationInfo PaginationTemplate
}
