package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/models"
)

func SaveOffer(offer models.Offer) (string, error) {
	return offerRepository.SaveOffers(offer)
}