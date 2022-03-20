package marketplaceRepository

import (
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
)

type OwnershipRepository struct{}

var Ownership = "ownership"

func (r *OwnershipRepository) SaveOwnership(ownership models.Ownership) (string, error) {
	return repository.Save[models.Ownership](ownership, Ownership)
}
