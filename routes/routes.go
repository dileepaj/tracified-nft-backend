package routes

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/routes/marketPlace"
	"github.com/dileepaj/tracified-nft-backend/routes/nftComposer"
)

var ApplicationRoutes models.Routers

func init() {
	routes := []models.Routers{
		testRoutes,
		marketPlace.NftRoutes,
		marketPlace.UserRoutes,
		marketPlace.WatchListRoutes,
		nftComposer.ComposerRoutes,
	}

	for _, r := range routes {
		ApplicationRoutes = append(ApplicationRoutes, r...)
	}
}
