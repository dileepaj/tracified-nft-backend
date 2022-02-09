package common

import (
	"github.com/dileepaj/tracified-nft-backend/marketplace/nft/etheruem"
	"github.com/dileepaj/tracified-nft-backend/marketplace/nft/polygon"
	"github.com/dileepaj/tracified-nft-backend/marketplace/nft/stellar"
)

type BlockChainNFT struct {
	BlackChainName string
}

type UserOwnNFT struct {
	BlackChainName string
	OwnerPK        string
	SellingStatus  string
}

type UserCreatedNFT struct {
	BlackChainName string
	OwnerPK        string
}

type NFTPriceHistory struct {
	BlackChainName string
	NFTTxnHash     string
}

func (bc *BlockChainNFT) GetAllSellingNFT() {
	GetAllSaleNFTFromDB()
}

func (bc *UserOwnNFT) GetUserOwnNFT() {
	GetAllUserOwnNFTFromDB()
}

func (bc *UserOwnNFT) GetUserCreatedNFT() {
	GetAllUserCreatedNFTFromDB()
}

func (bc *UserOwnNFT) GetAllNotSaleNFTFromDB() {
	GetAllNotSaleNFTFromDB()
}

func (bc *UserOwnNFT) TraceHistoryNFT() {
	switch bc.BlackChainName {
	case "POLYGON":
		polygon.TraceNFT()
	case "STELLAR":
		stellar.TraceNFT()
	case "ETHeruem":
		etheruem.TraceNFT()
	}
}

func (bc *UserOwnNFT) GetNFTPriceHistory() {
	switch bc.BlackChainName {
	case "POLYGON":
		polygon.GetNFTPriceHistory()
	case "STELLAR":
		stellar.GetNFTPriceHistory()
	case "ETHeruem":
		etheruem.GetNFTPriceHistory()
	}
}

func (bc *UserOwnNFT) ViewOneNFTProfile() {
	switch bc.BlackChainName {
	case "POLYGON":
		polygon.ViewOneNFTProfile()
	case "STELLAR":
		stellar.ViewOneNFTProfile()
	case "ETHeruem":
		etheruem.ViewOneNFTProfile()
	}
}

//get all NFT without considering the BC
func GetAllNFTForSale() {}
