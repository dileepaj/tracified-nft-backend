package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NFT struct {
	Id                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NFTIdentifier     string             `json:"nftidentifier" bson:"nftidentifier" ` //polygon eth ==> "tokenId - contractAddress ,stellar==> assetName - NFT issuer PK
	CreatorUserId     string             `json:"creatoruserid" bson:"creatoruserid" ` //minter user id //stellar==> distributor user Id assect creator user id
	Blockchain        string             `json:"blockchain" bson:"blockchain" `
	NFTName           string             `json:"nftname" bson:"nftname" `
	NftContentURL     string             `json:"nftcontenturl" bson:"nftcontenturl" `
	Description       string             `json:"description" bson:"description"`
	Timestamp         primitive.DateTime `json:"timestamp" bson:"timestamp,omitempty"`
	Collection        string             `json:"collection" bson:"collection" `
	Category          string             `json:"category" bson:"category" `
	Tags              []string           `json:"tags" bson:"tags" `
	Imagebase64       string             `json:"imagebase64" bson:"imagebase64" `
	CurrentPrice      string             `json:"currentprice" bson:"currentprice" `
	CurrentOwnerPK    string             `json:"currentownerpk" bson:"currentownerpk" `
	NFTIssuerPK       string             `json:"nftissuerpk" bson:"nftissuerpk" ` //minter pK for POLYGON.ETH.Solana , stellar ==>unioque created account
	ArtistName        string             `json:"artistname" bson:"artistname"`    //spacific artisct
	ArtistProfileLink string             `json:"artistprofilelink" bson:"artistprofilelink"`
	SellingStatus     string             `json:"sellingstatus" bson:"sellingstatus" ` //ONSALE,NOTSALE,NOTLISTED
	SellingType       string             `json:"sellingtype" bson:"sellingtype" `     //NOTLISTED
	DistributorPK     string             `json:"distributorpk" bson:"distributorpk"`  //specific for stellar (loged user PK when start to mint)
	MarketContract    string             `json:"marketcontract" bson:"marketcontract"`
	MintedContract    string             `json:"mintcontract" bson:"mintcontract"`
	TokenType         string             `json:"tokentype" bson:"tokentype"`
	Status            string             `json:"status" bson:"status" `
	Copies            string             `json:"copies" bson:"copies" `
}

//TenentName  	  string			 `json:"tenentname" bson:"ntenentname" validate:"required"` //com[pany Name
