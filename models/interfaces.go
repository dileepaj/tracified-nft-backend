package models

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SaveType interface {
	WatchList | User | NFT | Ownership | Offer | Widget | NFTComposerProject
}

type WidgetSaveResponse interface {
	string | responseDtos.WidgetSaveResponse
}

type InsertManyType interface {
	[]Widget
}

type FindOneType interface {
	string | primitive.ObjectID
}
