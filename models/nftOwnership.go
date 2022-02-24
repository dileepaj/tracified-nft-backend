package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ownership struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NFTIdentifier   string             `json:"nftidentifier" bson:"nftidentifier" validate:"required"`
	Blockchain      string             `json:"blockchain" bson:"blockchain" validate:"required"`
	Timestamp       string             `json:"timestamp" bson:"timestamp" validate:"required,datetime=Mon Jan 02 15:04:05 -0700 2006"`
	CurentOwnerPK   string             `json:"currentownerpk" bson:"currentownerpk" validate:"required"`
	PreviousOwnerPK string             `json:"previousownerpk" bson:"previousownerpk" `
	Status          string             `json:"status" bson:"status" validate:"required"` //ACTIVE ,EXPIRED, INCONTRACT
	OwnerRevisionID uint64             `json:"ownerrevisionid" bson:"ownerrevisionid" validate:"required"`
}
