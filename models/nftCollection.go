package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type NFTCollection struct {
	Id               primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	UserId           string             `json:"UserId" bson:"userid" `
	Timestamp        primitive.DateTime `json:"Timestamp" bson:"timestamp,omitempty"`
	CollectionName   string             `json:"CollectionName" bson:"collectionname"`
	OrganizationName string             `json:"OrganizationName" bson:"organizationname"`
	PublicKey        string             `json:"PublicKey" bson:"publickey"`
	IsPrivate        bool               `json:"isprivate" bson:"isprivate"`
}
