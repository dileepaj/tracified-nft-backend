package responseDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResponseCollectionUpdate struct {
	Id               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId           string             `json:"userid" bson:"userid" `
	Timestamp        primitive.DateTime `json:"timestamp" bson:"timestamp"`
	CollectionName   string             `json:"collectionname" bson:"collectionname" `
	OrganizationName string             `json:"organizationname" bson:"organizationname" `
}
