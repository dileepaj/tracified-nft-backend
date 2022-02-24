package requestWrappers

import "github.com/dileepaj/tracified-nft-backend/models"

type CreateNFTRequest struct {
	NFT models.NFT
	Ownership models.Ownership
}

type UpdateBFTRequest struct {
}