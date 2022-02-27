package dao

import (
	"fmt"

	"github.com/dileepaj/tracified-nft-backend/database/repository/nftRepository"
	"github.com/dileepaj/tracified-nft-backend/models"
)

func  GetNFTBySellingStatusANDWithoutUserCreatedNFT(id string,userPK string)([]models.NFT,error) {
	fmt.Println(id,userPK)
	return nftRepository.FindById1AndNotId2("sellingstatus",id,"currentownerpk",userPK)
}

func  GetAllNFTOwnByUser(idName string, id string) {
	//nftRepository.FindById1AndId2(idName, id)
}

func  GetAllNFTCreatedBYUser(idName string, id string) {
	//nftRepository.FindById(idName, id)
}

func  GetAllNotSaleNFT(idName string, id string) {
	//nftRepository.FindById(idName, id)
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