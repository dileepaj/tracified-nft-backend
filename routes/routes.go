package routes

import (
	"github.com/dileepaj/tracified-nft-backend/models"
)

var ApplicationRoutes models.Routers

func init() {
	routes := []models.Routers{
		testRoutes,
		nftRoutes,
		userRoutes,
	}

	for _, r := range routes {
		ApplicationRoutes = append(ApplicationRoutes, r...)
	}
}
