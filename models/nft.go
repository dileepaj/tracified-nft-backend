package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NFT struct {
	Id                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NFTIdentifier     string             `json:"nftidentifier" bson:"nftidentifier" validate:"required"` // polygon eth ==> "tokenId - contractAddress ,stellar==> assetName - NFT issuer PK
	CreatorUserId     string             `json:"creatoruserid" bson:"creatoruserid" validate:"required"` // minter user id //stellar==> distributor user Id assect creator user id
	Blockchain        string             `json:"blockchain" bson:"blockchain" validate:"required"`
	NFTName           string             `json:"nftname" bson:"nftname" validate:"required"`
	NftContentURL     string             `json:"nftcontenturl" bson:"nftcontenturl"`
	Description       string             `json:"description" bson:"description"`
	Timestamp         string             `json:"timestamp" bson:"timestamp"`
	Collection        string             `json:"collection" bson:"collection" validate:"required"`
	Category          string             `json:"categories" bson:"categories" validate:"required"`
	Tags              []string           `json:"tags" bson:"tags"`
	Imagebase64       string             `json:"imagebase64" bson:"imagebase64" validate:"required"`
	AttachmentType    string             `json:"attachmenttype" bson:"attachmenttype" validate:"required"`
	CurrentPrice      string             `json:"currentprice" bson:"currentprice" `
	CurrentOwnerPK    string             `json:"currentownerpk" bson:"currentownerpk" validate:"required"`
	IssuerPK          string             `json:"nftissuerpk" bson:"nftissuerpk" validate:"required"` // minter pK for POLYGON.ETH.Solana , stellar ==>unioque created account
	ArtistName        string             `json:"artistname" bson:"artistname"`                       // spacific artisct
	ArtistProfileLink string             `json:"artistprofilelink" bson:"artistprofilelink"`
	SellingStatus     string             `json:"sellingstatus" bson:"sellingstatus" validate:"required"` // ONSALE,NOTSALE,NOTLISTED
	SellingType       string             `json:"sellingtype" bson:"sellingtype" `                        // NOTLISTED
	DistributorPK     string             `json:"distributorpk" bson:"distributorpk"`                     // specific for stellar (loged user PK when start to mint)
	MraketContract    string             `json:"marketcontract" bson:"marketcontract"`
	MintedContract    string             `json:"mintcontract" bson:"mintcontract"`
	TokenType         string             `json:"tokentype" bson:"tokentype"`
	Status            string             `json:"status" bson:"status" validate:"required"`
	NFTTxnHash        string             `json:"nfttxnhash" bson:"nfttxnhash" `
	Trending          bool               `json:"trending" bson:"trending" `
	HotPicks          bool               `json:"hotpicks" bson:"hotpicks" `
	Royalty           string             `json:"royalty" bson:"royalty"`
	Thumbnail         string             `json:"thumbnail" bson:"thumbnail"`
	Commission        string             `json:"commission" bson:"commission"`
}
type NFTContentforMatrix struct {
	Id             primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	NFTIdentifier  string             `json:"nftidentifier" bson:"nftidentifier" validate:"required"`
	CreatorUserId  string             `json:"creatoruserid" bson:"creatoruserid" validate:"required"`
	Blockchain     string             `json:"blockchain" bson:"blockchain" validate:"required"`
	NFTName        string             `json:"nftname" bson:"nftname" validate:"required"`
	Imagebase64    string             `json:"imagebase64" bson:"imagebase64" validate:"required"`
	AttachmentType string             `json:"attachmenttype" bson:"attachmenttype" validate:"required"`
	SellingStatus  string             `json:"sellingstatus" bson:"sellingstatus" validate:"required"`
	Trending       bool               `json:"trending" bson:"trending" `
	HotPicks       bool               `json:"hotpicks" bson:"hotpicks" `
	CurrentOwnerPK string             `json:"currentownerpk" bson:"currentownerpk" validate:"required"`
	Thumbnail      string             `json:"thumbnail" bson:"thumbnail"`
}

type Paginateresponse struct {
	Content        []NFTContentforMatrix `json:"content" bson:"content" validate:"required"`
	PaginationInfo PaginationTemplate
}

type NFTStory struct {
	Id            primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	NFTIdentifier string             `json:"NFTIdentifier" bson:"nftidentifier"`
	Blockchain    string             `json:"Blockchain" bson:"blockchain"`
	NFTStory      string             `json:"NFTStory" bson:"nftstory"`
}

type CreatorInfo struct {
	ArtistName    string  `json:"name" bson:"name"`
	Email         string  `json:"Email" bson:"email"`
	PublicKey     string  `json:"PublicKey" bson:"publickey"`
	Averagerating float32 `json:"avgrating" bson:"avgrating"`
}
type PaginatedCreatorInfo struct {
	ArtistInfo     []CreatorInfo `json:"artistinfo" bson:"artistinfo" validate:"required"`
	PaginationInfo PaginationTemplate
}

type ContractInfo struct {
	ContractAddress string `json:"contractaddress" bson:"contractaddress"`
	User            string `json:"user" bson:"user"`
	Blockchain      string `json:"blockchain" bson:"blockchain"`
}

type WalletNFT struct {
	IssuerPublicKey string `json:"issuerpublickey" bson:"issuerpublickey"`
	NFTCreator      string `json:"nftcreator" bson:"nftcreator"`
	NFTName         string `json:"nftname" bson:"nftname"`
	NFTContent      string `json:"nftcontent" bson:"nftcontent"`
	Description     string `json:"description" bson:"description"`
	Collection      string `json:"collection" bson:"collection"`
	BlockChain      string `json:"blockChain" bson:"blockChain"`
	NFTStatus       string `json:"nftstatus" bson:"nftstatus"`
	OTP             string `json:"otp" bson:"otp"`
	Email           string `json:"email" bson:"email"`
	Timestamp       string `json:"timestamp" bson:"timestamp"`
	TXNHash         string `json:"txnhash" bson:"txnhash"`
	ShopID          string `json:"shopid" bson:"shopid"`
}

type ResponseWalletNFT struct {
	IssuerPublicKey string `json:"issuerpublickey" bson:"issuerpublickey"`
	NFTCreator      string `json:"nftcreator" bson:"nftcreator"`
	NFTName         string `json:"nftname" bson:"nftname"`
	NFTContent      string `json:"nftcontent" bson:"nftcontent"`
	Description     string `json:"description" bson:"description"`
	Collection      string `json:"collection" bson:"collection"`
	BlockChain      string `json:"blockChain" bson:"blockChain"`
	NFTStatus       string `json:"nftstatus" bson:"nftstatus"`
	Email           string `json:"email" bson:"email"`
	Timestamp       string `json:"timestamp" bson:"timestamp"`
	TXNHash         string `json:"txnhash" bson:"txnhash"`
	ShopID          string `json:"shopid" bson:"shopid"`
	Thumbnail       string
}

type WalletNFTTenantUser struct {
	ID         primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	Logo       string             `json:"logo" bson:"logo"`
	EmailTopic string             `json:"emailTopic" bson:"emailTopic"`
}

type NFTWalletState struct {
	Id              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	IssuerPublicKey string             `json:"issuerpublickey" bson:"issuerpublickey"`
	Blockchain      string             `json:"blockchain" bson:"blockchain"`
	NFTCreator      string             `json:"nftcreator" bson:"nftcreator"`
	NFTName         string             `json:"nftname" bson:"nftname"`
	NFTRequested    string             `json:"nftrequested" bson:"nftrequested"`
	CurrentOwner    string             `json:"currentowner" bson:"currentowner"`
	NFTStatus       string             `json:"nftstatus" bson:"nftstatus"`
	Timestamp       string             `json:"timestamp" bson:"timestamp"`
	ShopID          string             `json:"shopid" bson:"shopid"`
	Thumbnail       string             `json:"thumbnail" bson:"thumbnail"`
}

type NFTWalletStateTXN struct {
	Id              primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	IssuerPublicKey string             `json:"issuerpublickey" bson:"issuerpublickey"`
	Blockchain      string             `json:"blockchain" bson:"blockchain"`
	NFTCreator      string             `json:"nftcreator" bson:"nftcreator"`
	NFTName         string             `json:"nftname" bson:"nftname"`
	CurrentOwner    string             `json:"currentowner" bson:"currentowner"`
	NFTStatus       string             `json:"nftstatus" bson:"nftstatus"`
	Timestamp       string             `json:"timestamp" bson:"timestamp"`
	ShopID          string             `json:"shopid" bson:"shopid"`
	Thumbnail       string             `json:"thumbnail" bson:"thumbnail"`
	TXNHash         string             `json:"txnhash" bson:"txnhash"`
}

type WalletNFTContentforMatrix struct {
	Id              primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	IssuerPublicKey string             `json:"issuerpublickey" bson:"issuerpublickey"`
	Blockchain      string             `json:"blockchain" bson:"blockchain"`
	NFTCreator      string             `json:"nftcreator" bson:"nftcreator"`
	NFTName         string             `json:"nftname" bson:"nftname"`
	NFTRequested    string             `json:"nftrequested" bson:"nftrequested"`
	CurrentOwner    string             `json:"currentowner" bson:"currentowner"`
	NFTStatus       string             `json:"nftstatus" bson:"nftstatus"`
	Timestamp       string             `json:"timestamp" bson:"timestamp"`
	ShopID          string             `json:"shopid" bson:"shopid"`
	Thumbnail       string             `json:"thumbnail" bson:"thumbnail"`
}

type PaginateWalletNFTResponse struct {
	Content        []WalletNFTContentforMatrix `json:"walletcontent" bson:"walletcontent" validate:"required"`
	PaginationInfo PaginationTemplate
}
