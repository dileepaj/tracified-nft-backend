package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateSVG(svg models.SVG) (string, error) {
	return CollectionRepository.SaveSVG(svg)
}

func UpdateSVGBlockchain(svg models.SVG) (models.SVG, error) {
	update := bson.M{
		"$set": bson.M{"blockchain": svg.Blockchain},
	}
	return CollectionRepository.UpdateSVGBlockchain(svg.Id, update)
}
