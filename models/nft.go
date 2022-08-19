package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NFT struct {
	Id                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NFTIdentifier     string             `json:"nftidentifier" bson:"nftidentifier" validate:"required"` //polygon eth ==> "tokenId - contractAddress ,stellar==> assetName - NFT issuer PK
	CreatorUserId     string             `json:"creatoruserid" bson:"creatoruserid" validate:"required"` //minter user id //stellar==> distributor user Id assect creator user id
	Blockchain        string             `json:"blockchain" bson:"blockchain" validate:"required"`
	NFTName           string             `json:"nftname" bson:"nftname" validate:"required"`
	NftContentURL     string             `json:"nftcontenturl" bson:"nftcontenturl" validate:"required"`
	Description       string             `json:"description" bson:"description"`
	Timestamp         string             `json:"timestamp" bson:"timestamp"`
	Collection        string             `json:"collection" bson:"collection" validate:"required"`
	Category          string             `json:"categories" bson:"categories" validate:"required"`
	Tags              string             `json:"tags" bson:"tags"`
	Imagebase64       string             `json:"imagebase64" bson:"imagebase64" validate:"required"`
	CurrentPrice      string             `json:"currentprice" bson:"currentprice" `
	CurrentOwnerPK    string             `json:"currentownerpk" bson:"currentownerpk" validate:"required"`
	IssuerPK          string             `json:"nftissuerpk" bson:"nftissuerpk" validate:"required"` //minter pK for POLYGON.ETH.Solana , stellar ==>unioque created account
	ArtistName        string             `json:"artistname" bson:"artistname"`                       //spacific artisct
	ArtistProfileLink string             `json:"artistprofilelink" bson:"artistprofilelink"`
	SellingStatus     string             `json:"sellingstatus" bson:"sellingstatus" validate:"required"` //ONSALE,NOTSALE,NOTLISTED
	SellingType       string             `json:"sellingtype" bson:"sellingtype" `                        //NOTLISTED
	DistributorPK     string             `json:"distributorpk" bson:"distributorpk"`                     //specific for stellar (loged user PK when start to mint)
	MraketContract    string             `json:"marketcontract" bson:"marketcontract"`
	MintedContract    string             `json:"mintcontract" bson:"mintcontract"`
	TokenType         string             `json:"tokentype" bson:"tokentype"`
	Status            string             `json:"status" bson:"status" validate:"required"`
	NFTTxnHash        string             `json:"nfttxnhash" bson:"nfttxnhash" `
	Trending          bool               `json:"trending" bson:"trending" `
	HotPicks          bool               `json:"hotpicks" bson:"hotpicks" `
}

//TenentName  	  string			 `json:"tenentname" bson:"ntenentname" validate:"required"` //com[pany Name
type NFTStory struct {
	Id            primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	NFTIdentifier string             `json:"NFTIdentifier" bson:"nftidentifier"`
	Blockchain    string             `json:"Blockchain" bson:"blockchain"`
	NFTStory      string             `json:"NFTStory" bson:"nftstory"`
}
