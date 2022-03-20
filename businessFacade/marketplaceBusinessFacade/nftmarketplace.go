package marketplaceBusinessFacade

import (
	"encoding/json"

	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
)

func StoreNFT(createNFTObject requestDtos.CreateNFTRequest) (string, error) {
	rst, err1 := nftRepository.SaveNFT(createNFTObject.NFT)

	if err1 != nil {
		return "NFT not saved", err1
	} else {
		_, err2 := SaveOwnership(createNFTObject.Ownership)
		if err2 != nil {
			return "Ownership not saved", err2
		} else {
			return rst, nil
		}
	}
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
