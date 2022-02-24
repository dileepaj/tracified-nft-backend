package dao

import (
	"github.com/dileepaj/tracified-nft-backend/database/repository/nftRepository"
	"github.com/dileepaj/tracified-nft-backend/models"
)

func  GetAllNFTBySellingStatus(idName string, id string) {
	nftRepository.FindById(idName, id)
}

func  GetAllNFTOwnByUser(idName string, id string) {
	nftRepository.FindById(idName, id)
}

func  GetAllNFTCreatedBYUser(idName string, id string) {
	nftRepository.FindById(idName, id)
}

func  GetAllNotSaleNFT(idName string, id string) {
	nftRepository.FindById(idName, id)
}

func  SaveNFT(nft models.NFT)(string, error) {
	return nftRepository.Save(nft)
}

// func  MakeSale(id string, nft *models.NFT){
// 	nftRepository.Update(id,nft)
// }

// func  MakeBuy(id string, nft *models.NFT){
// 	nftRepository.Update(id,nft)
// }