package routes

import (
	"github.com/dileepaj/tracified-nft-backend/businessFacades"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var nftRoutes = models.Routers{

	models.Router{
		Name:    "Save NFT",
		Method:  "POST",
		Path:    "/api/nft/save",
		Handler: businessFacades.SaveNFT,
	},
}