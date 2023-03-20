package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SaveType interface {
	WatchList | User | NFT | Ownership | Offer | Widget | NFTComposerProject | Chart | Table | StataArray | ProofBotData | ImageData | Timeline | Review | NewsLetter | Faq | NFTCollection | Tags | SVG | TXN | Favourite | Endorse | Partner | Document | []TXN | UserAuth | TDP | UserNFTMapping | NFTStory | UserQuestions | Subscription | GeneratedMap
}

type InsertManyType interface {
	[]Widget
}

type FindOneType interface {
	string | primitive.ObjectID
}
