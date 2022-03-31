package routes

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

//This routes use to check the API status
var testRoutes = models.Routers{

	models.Router{
		Name:"Connection test API",
		Method:"GET",
		Path:"/api/health",
		Handler:apiHandler.HealthCheck,
	},
}