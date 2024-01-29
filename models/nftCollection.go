package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type NFTCollection struct {
	Id               primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	UserId           string             `json:"UserId" bson:"userid" `
	Timestamp        primitive.DateTime `json:"Timestamp" bson:"timestamp,omitempty"`
	CollectionName   string             `json:"CollectionName" bson:"collectionname"`
	OrganizationName string             `json:"OrganizationName" bson:"organizationname"`
	IsPublic         bool               `json:"ispublic" bson:"ispublic"`
	CID              string             `json:"cid" bson:"cid"`
	Images           []ImageObject      `json:"images" bson:"images"`
	NFTCount         int64              `json:"nftcount" bson:"nftcount"`
}
type CollectionPaginationResponse struct {
	Content        []NFTCollection
	PaginationInfo PaginationTemplate
}
