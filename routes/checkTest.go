package routes

import (
	"github.com/dileepaj/tracified-nft-backend/businessFacades"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var testRoutes = models.Routers{

	models.Router{
		Name:"Connection test API",
		Method:"GET",
		Path:"/api/health",
		Handler:businessFacades.HealthCheck,
	},
}