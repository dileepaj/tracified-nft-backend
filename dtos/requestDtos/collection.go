package requestDtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateCollection struct {
	Id             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CollectionName string             `json:"collectionname" bson:"collectionname"`
}

type DeleteCollectionById struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

type DeleteCollectionByUserPK struct {
	UserId string `json:"userid" bson:"userid"`
}

type UpdateCollectionVisibility struct {
	Id       primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	IsPublic bool               `json:"ispublic" bson:"ispublic"`
}
