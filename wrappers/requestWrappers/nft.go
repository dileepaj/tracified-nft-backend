package requestWrappers

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateNFTRequest struct {
	NFT models.NFT
	Ownership models.Ownership
}

type UpdateNFTSALERequest struct {
	NFTIdentifier	string `json:"nftidentifier" bson:"nftidentifier" validate:"required"`
	Timestamp       primitive.DateTime `json:"timestamp" bson:"timestamp"`
	CurrentPrice    string `json:"currentprice" bson:"currentprice" `
	SellingStatus   string `json:"sellingstatus" bson:"sellingstatus" validate:"required"` //ONSALE,NOTSALE,NOTLISTED
	SellingType     string `json:"sellingtype" bson:"sellingtype" ` //NOTLISTED
	MarketContract	string `json:"smartcontract" bson:"smartcontract"`
}