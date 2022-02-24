package ownershipRepository

import (
	"context"
	"time"

	"github.com/dileepaj/tracified-nft-backend/connections"
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Save(ownership models.Ownership) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rst, err := connections.Connect().Collection("ownership").InsertOne(ctx, ownership)
	if err != nil {
		return ownership.NFTIdentifier, err
	}
	var id =rst.InsertedID.(primitive.ObjectID)
	return id.String(), nil
}
