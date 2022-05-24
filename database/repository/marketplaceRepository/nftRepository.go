package marketplaceRepository

import (
	"context"
	"time"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"

	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	rst, err := repository.FindById(idName, id, NFT)
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

func (r *NFTRepository) GetSVGByHash(idName string, id string) ([]models.SVG, error) {
	if err != nil {
		return svgs, err
	}
	for rst.Next(context.TODO()) {
		var svg models.SVG
		err = rst.Decode(&svg)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return svgs, err
		}
		svgs = append(svgs, svg)
	}
	return svgs, nil
}

func (r *NFTRepository) FindLastNFTById(idName string, id string) ([]models.NFT, error) {
	var nfts []models.NFT
	rst := repository.FindOne(idName, id, "nft")
	if rst != nil {
		err := rst.Decode(&nfts)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nfts, err
		}
		return nfts, nil
	} else {
		return nfts, nil
	}
}

func (r *NFTRepository) FindByFieldInMultipleValus(fields string, tags []string) ([]models.NFT, error) {
	var nfts []models.NFT
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
	projection := bson.M{}
	rst := repository.FindOneAndUpdate("nftidentifier", nft.NFTIdentifier, projection, update, NFT)
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

func (repository *NFTRepository) UpdateNFTMinter(nft requestDtos.UpdateMint) (responseDtos.ResponseNFTMinter, error) {
	var responseNFT responseDtos.ResponseNFTMinter
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	update := bson.M{
		"$set": bson.M{"nftissuerpk": nft.NFTIssuerPK, "nftidentifier": nft.NFTIdentifier, "nfttxnhash": nft.NFTTxnHash},
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := connections.Connect().Collection("nft").FindOneAndUpdate(ctx, bson.M{"imagebase64": nft.Imagebase64}, update, &opt).Decode(&responseNFT)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	return responseNFT, err
}

func (repository *NFTRepository) UpdateNFTTXN(nft requestDtos.UpdateMintTXN) (responseDtos.ResponseNFTMintTXN, error) {
	var responseNFT responseDtos.ResponseNFTMintTXN
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	update := bson.M{
		"$set": bson.M{"nfttxnhash": nft.NFTTxnHash},
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := connections.Connect().Collection("nft").FindOneAndUpdate(ctx, bson.M{"imagebase64": nft.Imagebase64}, update, &opt).Decode(&responseNFT)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	return responseNFT, err
}

func (repository *NFTRepository) FindTagsbyNFTIdentifier(idName1 string, id1 string) ([]models.Tags, error) {
	var tags []models.Tags
	if idName1 != "" {
		findOptions := options.Find()
		rst, err := connections.Connect().Collection("tags").Find(context.TODO(), bson.D{{idName1, id1}}, findOptions)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return tags, err
		}
		for rst.Next(context.TODO()) {
			var tag models.Tags
			err = rst.Decode((&tag))
			if err != nil {
				logs.ErrorLogger.Println(err.Error())
				return tags, err
			}
			tags = append(tags, tag)
		}
		return tags, nil
	} else {
		return tags, nil
	}
}

func (repository *NFTRepository) GetAllTags() ([]models.Tags, error) {
	var tag []models.Tags
	findOptions := options.Find()
	findOptions.SetLimit(10)
	rst, err := connections.Connect().Collection("tags").Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return tag, err
	}
	for rst.Next(context.TODO()) {
		var tags models.Tags
		err = rst.Decode((&tags))
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return tag, err
		}
		tag = append(tag, tags)
	}
	return tag, nil
}
