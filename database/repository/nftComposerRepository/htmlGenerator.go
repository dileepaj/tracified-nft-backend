package nftcomposerrepository

import (
	"context"
	"time"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type HTMLNFTRepository struct{}

/**
Save the Json tha used to create HTML file of NFT
**/
func (r *HTMLNFTRepository)SaveHtmlData(htmlData models.HtmlGenerator) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rst, err := connections.Connect().Collection("htmlnft").InsertOne(ctx, htmlData)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return htmlData.Id.String(), err
	}
	var id = rst.InsertedID.(primitive.ObjectID)
	return id.String(), nil
}