package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/api"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var UserRoutes = models.Routers{

	models.Router{
		Name:    "Get eds by id",
		Method:  "POST",
		Path:    "/api/user/create",
		Handler: api.CreateUser,
	},
}
