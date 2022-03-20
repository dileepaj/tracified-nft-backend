package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

//This routes handle the all user related routes in the marketplace
var UserRoutes = models.Routers{

	models.Router{
		Name:    "Get eds by id",
		Method:  "POST",
		Path:    "/api/user/create",
		Handler: apiHandler.CreateUser,
	},
}
