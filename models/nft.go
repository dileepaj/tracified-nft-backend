package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NFT struct {
	Id                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NFTIdentifier     string             `json:"nftidentifier" bson:"nftidentifier" validate:"required"` //polygon eth ==> "tokenId - contractAddress ,stellar==> assetName - NFT issuer PK 
	CreatorUserId	  string			 `json:"creatoruserid" bson:"creatoruserid" validate:"required"`//minter user id //stellar==> distributor user Id assect creator user id 
	Blockchain        string             `json:"blockchain" bson:"blockchain" validate:"required"`
	NFTName           string             `json:"nftname" bson:"nftname" validate:"required"`
	NftContentURL     string             `json:"nftcontenturl" bson:"nftcontenturl" validate:"required,url"`
	Description       string             `json:"description" bson:"description"`
	Timestamp         primitive.DateTime `json:"timestamp" bson:"timestamp"`
	Collection        string             `json:"collection" bson:"collection" validate:"required"`
	Category          string             `json:"category" bson:"category" validate:"required"`
	Tags              []string           `json:"tags" bson:"tags" validate:"required"`
	Imagebase64       string             `json:"imagebase64" bson:"imagebase64" validate:"required,base64"`
	CurrentPrice      string             `json:"currentprice" bson:"currentprice" `
	CurrentOwnerPK    string             `json:"currentownerpk" bson:"currentownerpk" validate:"required"`
	IssuerPK          string             `json:"issuerpk" bson:"issuerpk" validate:"required"` //minter pK for POLYGON.ETH.Solana , stellar ==>unioque created account
	ArtistName        string             `json:"artistname" bson:"artistname"` //spacific artisct 
	ArtistProfileLink string             `json:"artistprofilelink" bson:"artistprofilelink"`
	SellingStatus     string             `json:"sellingstatus" bson:"sellingstatus" validate:"required"` //ONSALE,NOTSALE,NOTLISTED
	SellingType       string             `json:"sellingtype" bson:"sellingtype" `                        //NOTLISTED
	DistributorPK     string             `json:"distributorpk" bson:"distributorpk"` //specific for stellar (loged user PK when start to mint)
	MraketContract    string             `json:"marketcontract" bson:"marketcontract"`
	MintedContract    string             `json:"mintcontract" bson:"mintcontract"`
	TokenType         string             `json:"tokentype" bson:"tokentype"`
	Status            string             `json:"status" bson:"status" validate:"required"`
}

//TenentName  	  string			 `json:"tenentname" bson:"ntenentname" validate:"required"` //com[pany Name