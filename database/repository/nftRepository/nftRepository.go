package nftRepository

import (
	"context"
	"time"

	connetions "github.com/dileepaj/tracified-nft-backend/connections"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindById1AndNotId2(idName1 string,id1 string,idName2 string, id2 string) ([]models.NFT,error) {
	var nfts []models.NFT
	if (idName1!="" && idName2!=""){
		findOptions := options.Find()
		findOptions.SetSort(bson.D{{"timestamp", -1}})
		rst, err := connetions.Connect().Collection("nft").Find(context.TODO(), bson.D{{idName1, id1},{idName2,bson.D{{"$ne",id2}}}},findOptions)
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
	}else{
		return nfts,nil
	}
}

func FindById1AndId2(idName1 string,id1 string,idName2 string, id2 string) ([]models.NFT,error) {
	var nfts []models.NFT
	if (idName1!="" && idName2!=""){
		findOptions := options.Find()
		findOptions.SetSort(bson.D{{"timestamp", 1}})
		rst, err := connetions.Connect().Collection("nft").Find(context.TODO(), bson.D{{idName1, id1},{idName2,id2}},findOptions)
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
	}else{
		return nfts,nil
	}
}

func FindById(idName1 string,id1 string) ([]models.NFT,error) {
	var nfts []models.NFT
	if (idName1!="" ){
		findOptions := options.Find()
		findOptions.SetSort(bson.D{{"timestamp", -1}})
		rst, err := connetions.Connect().Collection("nft").Find(context.TODO(), bson.D{{idName1, id1}},findOptions)
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
	}else{
		return nfts,nil
	}
}

func Save(nft models.NFT)  (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rst,err := connetions.Connect().Collection("nft").InsertOne(ctx,nft)
	if err != nil{
		logs.ErrorLogger.Println(err.Error())
		return nft.NFTIdentifier, err
	}
	var id =rst.InsertedID.(primitive.ObjectID)
	return id.String(), nil
}

func Update() {}
