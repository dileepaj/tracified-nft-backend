package nftRepository

import (
	"context"
	"time"

	connetions "github.com/dileepaj/tracified-nft-backend/connections"
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindById(IdName string, id string) {}

func Save(nft models.NFT)  (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rst,err := connetions.Connect().Collection("nft").InsertOne(ctx,nft)
	if err != nil{
		return nft.NFTIdentifier, err
	}
	var id =rst.InsertedID.(primitive.ObjectID)
	return id.String(), nil
}

func Update() {}
