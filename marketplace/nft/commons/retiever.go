package commons

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

func (bc *BlockChainNFT) GetAllSellingNFT() {
	switch bc.BlackChainName {
	case "POLYGON":
		polygon.GetAllSaleNFTDB()
	case "STELLAR":
		stellar.GetAllSaleNFTDB()
	case "ETHeruem":
		etheruem.GetAllNotSaleNFTDB()
	}
}

func (bc *UserOwnNFT) GetUserOwnNFT() {
	switch bc.BlackChainName {
	case "POLYGON":
		polygon.GetAllUserOwnNFTDB()
	case "STELLAR":
		stellar.GetAllUserOwnNFTDB()
	case "ETHeruem":
		etheruem.GetAllUserOwnNFTDB()
	}
}

func (bc *UserOwnNFT) GetUserCreatedNFT() {
	switch bc.BlackChainName {
	case "POLYGON":
		polygon.GetAllUserCreatedNFTDB()
	case "STELLAR":
		stellar.GetAllUserCreatedNFTDB()
	case "ETHeruem":
		etheruem.GetAllUserCreatedNFTDB()
	}
}

func (bc *UserOwnNFT) GetAllNotSaleNFTDB() {
	switch bc.BlackChainName {
	case "POLYGON":
		polygon.GetAllNotSaleNFTDB()
	case "STELLAR":
		stellar.GetAllNotSaleNFTDB()
	case "ETHeruem":
		etheruem.GetAllNotSaleNFTDB()
	}
}

func (bc *UserOwnNFT) TraceNFT() {
	switch bc.BlackChainName {
	case "POLYGON":
		polygon.TraceNFT()
	case "STELLAR":
		stellar.TraceNFT()
	case "ETHeruem":
		etheruem.TraceNFT()
	}
}

//get all NFT without considering the DB
func GetAllNFTForSale() {}
