package marketplaceRepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"

	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
)

type NFTRepository struct{}

var NFT = "nft"

func (r *NFTRepository) FindNFTById1AndNotId2(idName1 string, id1 string, idName2 string, id2 string) ([]models.NFT, error) {
	var nfts []models.NFT
	rst, err := repository.FindById1AndNotId2(idName1, id1, idName2, id2, NFT)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nfts, err
	}
	for rst.Next(context.TODO()) {
		var nft models.NFT
		err = rst.Decode(&nft)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nfts, err
		}
		nfts = append(nfts, nft)
	}
	return nfts, nil
}

func (r *NFTRepository) FindNFTsById(idName string, id string) ([]models.NFT, error) {
	var nfts []models.NFT
	rst, err := repository.FindById(idName, id, "nft")
	if err != nil {
		return nfts, err
	}
	for rst.Next(context.TODO()) {
		var nft models.NFT
		err = rst.Decode(&nft)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nfts, err
		}
		nfts = append(nfts, nft)
	}
	return nfts, nil
}

func (r *NFTRepository) FindByFieldInMultipleValus(fields string, tags []string) ([]models.NFT, error) {
	var nfts []models.NFT
	rst, err := repository.FindByFieldInMultipleValus(fields, tags, NFT)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nfts, err
	}
	for rst.Next(context.TODO()) {
		var nft models.NFT
		err = rst.Decode(&nft)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nfts, err
		}
		nfts = append(nfts, nft)
	}
	return nfts, nil
}

func (r *NFTRepository) SaveNFT(nft models.NFT) (string, error) {
	return repository.Save[models.NFT](nft, NFT)
}

func (r *NFTRepository) UpdateNFTSALE(nft requestDtos.UpdateNFTSALERequest) (responseDtos.ResponseNFTMakeSale, error) {
	var responseMakeSaleNFT responseDtos.ResponseNFTMakeSale
	update := bson.M{
		"$set": bson.M{"timestamp": nft.Timestamp, "sellingstatus": nft.SellingStatus, "sellingtype": nft.SellingType, "marketcontract": nft.MarketContract},
	}
	rst:= repository.FindOneAndUpdate("nftidentifier", nft.NFTIdentifier, update, NFT)
	if rst != nil {
		err := rst.Decode(&responseMakeSaleNFT)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return responseMakeSaleNFT, err
		}
		return responseMakeSaleNFT, nil
	} else {
		return responseMakeSaleNFT, nil
	}
}
