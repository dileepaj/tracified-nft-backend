package marketplaceRepository

import (
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
)

type OfferRepository struct{}

var Offer = "offer"

func (r *OfferRepository) SaveOffers(offer models.Offer) (string, error) {
	return repository.Save[models.Offer](offer, NFT)
}
