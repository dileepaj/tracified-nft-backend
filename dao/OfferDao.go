package dao

import (
	"github.com/dileepaj/tracified-nft-backend/database/repository/offerRepository"
	"github.com/dileepaj/tracified-nft-backend/models"
)

func SaveOffer(offer models.Offer) (string, error) {
	return offerRepository.Save(offer)
}