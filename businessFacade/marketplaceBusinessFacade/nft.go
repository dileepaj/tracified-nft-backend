package marketplaceBusinessFacade

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
)

func StoreNFT(createNFTObject models.NFT) (string, error) {
	rst, err1 := nftRepository.SaveNFT(createNFTObject)
	if err1 != nil {
		return "NFT not saved", err1
	}
	return rst, nil

}

func StoreOwner(createOwner models.Ownership) (string, error) {
	rst, err1 := nftRepository.SaveOwner(createOwner)
	if err1 != nil {
		return "Owner not saved", err1
	}
	return rst, nil
}

func GetAllONSaleNFT(id string, userPK string) ([]models.NFT, error) {
	return nftRepository.FindNFTById1AndNotId2("sellingstatus", id, "currentownerpk", userPK)
}

func MakeSaleNFT(update requestDtos.UpdateNFTSALERequest) (responseDtos.ResponseNFTMakeSale, error) {
	return nftRepository.UpdateNFTSALE(update)
}

func GetBlockchainSpecificNFT(blockchain string) ([]models.NFT, error) {
	return nftRepository.FindNFTsById("blockchain", blockchain)
}

func GetNFTBySellingStatus(status string) ([]models.NFT, error) {
	return nftRepository.FindNFTsById("sellingstatus", status)
}

func GetNFTbyTagsName(tags string) ([]models.NFT, error) {
	var tagsArray []string
	_ = json.Unmarshal([]byte(tags), &tagsArray)
	return nftRepository.FindByFieldInMultipleValus("tags", tagsArray)
}

func GetWatchListNFT(userId string) ([]models.NFT, error) {
	results, err := FindNFTIdentifieryByUserId(userId)
	if err != nil || len(results) == 0 {
		return []models.NFT{}, err
	} else {
		return nftRepository.FindByFieldInMultipleValus("nftidentifier", results)
	}
}

func GetNFTbyAccount(userId string) ([]models.NFT, error) {
	results, err := GetBCAccountPKByUserId(userId)
	if err != nil || len(results) == 0 {
		return []models.NFT{}, err
	} else {
		return nftRepository.FindByFieldInMultipleValus("currentownerpk", results)
	}
}

func GetNFTbyTenentName(tenentName string) ([]models.NFT, error) {
	results, err := GetBCAccountPKByTenetName(tenentName)
	if err != nil || len(results) == 0 {
		return []models.NFT{}, err
	} else {
		return nftRepository.FindByFieldInMultipleValus("currentownerpk", results)
	}
}

func CreateTags(tags models.Tags) (string, error) {
	log.Println("------------------------------------testing 7 ---------------------------------------------------")
	return nftRepository.SaveTags(tags)
}

func GetAllTags() ([]models.Tags, error) {
	fmt.Println("Calling repo...")
	return nftRepository.GetAllTags()
}
func GetTagsByNFTIdentifier(nftid string) ([]models.Tags, error) {
	log.Println("---------------------------------------test 2------------------------------", nftid)
	return nftRepository.FindTagsbyNFTIdentifier("nftidentifier", nftid)

}

func UpdateNFT(update requestDtos.UpdateMint) (responseDtos.ResponseNFTMinter, error) {
	return nftRepository.UpdateNFTMinter(update)
}
