package requestDtos

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateNFTRequest struct {
	NFT       models.NFT
	Ownership models.Ownership
}

type UpdateNFTSALERequest struct {
	NFTIdentifier  string             `json:"nftidentifier" bson:"nftidentifier" validate:"required"`
	Timestamp      primitive.DateTime `json:"timestamp" bson:"timestamp"`
	CurrentPrice   string             `json:"currentprice" bson:"currentprice" `
	SellingStatus  string             `json:"sellingstatus" bson:"sellingstatus"` //ONSALE,NOTSALE,NOTLISTED
	SellingType    string             `json:"sellingtype" bson:"sellingtype" `    //NOTLISTED
	MarketContract string             `json:"marketcontract" bson:"marketcontract"`
	CurrentOwnerPK string             `json:"currentownerpk" bson:"currentownerpk"`
}

type UpdateMint struct {
	Imagebase64   string `json:"imagebase64" bson:"imagebase64" `
	NFTIssuerPK   string `json:"nftissuerpk" bson:"nftissuerpk" `
	NFTIdentifier string `json:"nftidentifier" bson:"nftidentifier" `
	NFTTxnHash    string `json:"nfttxnhash" bson:"nfttxnhash" `
}

type UpdateMintTXN struct {
	Imagebase64 string `json:"imagebase64" bson:"imagebase64" `
	NFTTxnHash  string `json:"nfttxnhash" bson:"nfttxnhash" `
}
