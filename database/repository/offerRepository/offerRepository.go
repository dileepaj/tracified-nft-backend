package offerRepository

import (
	"context"
	"time"

	"github.com/dileepaj/tracified-nft-backend/connections"
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Save(offer models.Offer) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rst, err := connections.Connect().Collection("offer").InsertOne(ctx, offer)
	if err != nil {
		return offer.NFTIdentifier, err
	}
	var id =rst.InsertedID.(primitive.ObjectID)
	return id.String(), nil
}
