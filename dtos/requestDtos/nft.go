package requestDtos

import (
	"github.com/dileepaj/tracified-nft-backend/models"
)

type CreateNFTRequest struct {
	NFT       models.NFT
	Ownership models.Ownership
}

type UpdateNFTSALERequest struct {
	NFTIdentifier  string `json:"nftidentifier" bson:"nftidentifier"`
	Timestamp      string `json:"timestamp" bson:"timestamp"`
	CurrentPrice   string `json:"currentprice" bson:"currentprice" `
	SellingStatus  string `json:"sellingstatus" bson:"sellingstatus"` //ONSALE,NOTSALE,NOTLISTED
	SellingType    string `json:"sellingtype" bson:"sellingtype" `    //NOTLISTED
	MarketContract string `json:"marketcontract" bson:"marketcontract"`
	CurrentOwnerPK string `json:"currentownerpk" bson:"currentownerpk"`
	Royalty        string `json:"royalty" bson:"royalty"`
	Blockchain     string `json:"blockchain" bson:"blockchain"`
}

type UpdateMint struct {
	Imagebase64   string `json:"imagebase64" bson:"imagebase64" `
	NFTIssuerPK   string `json:"nftissuerpk" bson:"nftissuerpk" `
	NFTIdentifier string `json:"nftidentifier" bson:"nftidentifier" `
	NFTTxnHash    string `json:"nfttxnhash" bson:"nfttxnhash" `
	Blockchain    string `json:"blockchain" bson:"blockchain" `
}

type UpdateMintTXN struct {
	Imagebase64 string `json:"imagebase64" bson:"imagebase64" `
	NFTTxnHash  string `json:"nfttxnhash" bson:"nfttxnhash" `
	Blockchian  string `json:"blockchain" bson:"blockchain" `
}

type NFTsForMatrixView struct {
	Blockchain    string `json:"blockchain" bson:"blockchain"`
	PageSize      int32  `json:"pagesize" bson:"pagesize"`
	RequestedPage int32  `json:"requestedPage" bson:"requestedPage" `
	SortbyFeild   string `json:"sortbyfeild" bson:"sortbyfeild" `
}

type CreatorInfoforMatrixView struct {
	//CreatorPK     string `json:"userpk" bson:"userpk"`
	PageSize      int32 `json:"pagesize" bson:"pagesize"`
	RequestedPage int32 `json:"requestedPage" bson:"requestedPage" `
}
