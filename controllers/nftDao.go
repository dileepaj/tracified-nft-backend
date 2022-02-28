package controllers

import (
	"fmt"

	"github.com/dileepaj/tracified-nft-backend/database/repository/nftRepository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/wrappers/requestWrappers"
	"github.com/dileepaj/tracified-nft-backend/wrappers/responseWrappers"
)

func  GetNFTBySellingStatusANDWithoutUserCreatedNFT(id string,userPK string)([]models.NFT,error) {
	fmt.Println(id,userPK)
	return nftRepository.FindById1AndNotId2("sellingstatus",id,"currentownerpk",userPK)
}

func MakeSaleNFT(update requestWrappers.UpdateNFTSALERequest)(responseWrappers.ResponseNFTMakeSale,error){
	return nftRepository.UpdateNFTSALE(update)
}

func  GetNFTbyBlockChain(blockchain string)([]models.NFT,error) {
	return nftRepository.FindById("blockchain",blockchain)
}

func  GetNFTBySellingStatus(status string)([]models.NFT,error) {
	return nftRepository.FindById("sellingstatus",status)
}

func  GetNFTbyTennet(name string) {
	nftRepository.FindById("tenetname", name)
}

func  GetNFTbyTags(tags []string)([]models.NFT,error) {
	return nftRepository.FindByFieldInMultipleValus("tags", tags)
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