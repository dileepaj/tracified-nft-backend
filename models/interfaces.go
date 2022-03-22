package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SaveType interface {
	WatchList | User | NFT | Ownership | Offer | Widget | NFTComposerProject | NFTCollection
}

type InsertManyType interface {
	[]Widget
}

type FindOneType interface {
	string | primitive.ObjectID
}
