package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/models"
)

func SaveOwnership(ownership models.Ownership) (string, error) {
	return ownershipRepository.SaveOwnership(ownership)
}
