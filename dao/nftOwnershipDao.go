package dao

import (
	"github.com/dileepaj/tracified-nft-backend/database/repository/ownershipRepository"
	"github.com/dileepaj/tracified-nft-backend/models"
)

func SaveOwnership(ownership models.Ownership) (string, error) {
	return ownershipRepository.Save(ownership)
}