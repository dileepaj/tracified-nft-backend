package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/businessFacades"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var UserRoutes = models.Routers{

	models.Router{
		Name:    "Get eds by id",
		Method:  "POST",
		Path:    "/api/user/create",
		Handler: businessFacades.CreateUser,
	},
}
