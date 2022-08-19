package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ownership struct {
	Id              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NFTIdentifier   string             `json:"nftidentifier" bson:"nftidentifier" validate:"required"`
	Timestamp       primitive.DateTime `json:"timestamp" bson:"timestamp"`
	CurentOwnerPK   string             `json:"currentownerpk" bson:"currentownerpk" validate:"required"`
	PreviousOwnerPK string             `json:"previousownerpk" bson:"previousownerpk" `
	Status          string             `json:"status" bson:"status" validate:"required"` //ACTIVE ,EXPIRED, INCONTRACT
	OwnerRevisionID uint64             `json:"ownerrevisionid" bson:"ownerrevisionid" validate:"required"`
}
