package nftController

import (
	"github.com/dileepaj/tracified-nft-backend/database/repository/nftRepository"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var repository nftRepository.NFTRepository

func GetNFTBySellingStatusAndNotUserCreated(id string, userPK string) ([]models.NFT, error) {
	return repository.FindById1AndNotId2("sellingstatus", id, "currentownerpk", userPK)
}

func MakeSaleNFT(update requestDtos.UpdateNFTSALERequest) (responseDtos.ResponseNFTMakeSale, error) {
	return repository.UpdateNFTSALE(update)
}

func GetNFTbyBlockChain(blockchain string) ([]models.NFT, error) {
	return repository.FindById("blockchain", blockchain)
}

func GetNFTBySellingStatus(status string) ([]models.NFT, error) {
	return repository.FindById("sellingstatus", status)
}

func GetNFTbyTags(tags []string) ([]models.NFT, error) {
	return repository.FindByFieldInMultipleValus("tags", tags)
}

func GetNFTbyNFTIdentifier(tags []string) ([]models.NFT, error) {
	return repository.FindByFieldInMultipleValus("nftidentifier", tags)
}

func GetNFTbyAccount(accounts []string) ([]models.NFT, error) {
	return repository.FindByFieldInMultipleValus("currentownerpk", accounts)
}

func CreateNFT(nft models.NFT) (string, error) {
	return repository.Save(nft)
}

// func  MakeSale(id string, nft *models.NFT){
// 	nftRepository.Update(id,nft)
// }

// func  MakeBuy(id string, nft *models.NFT){
// 	nftRepository.Update(id,nft)
// }
