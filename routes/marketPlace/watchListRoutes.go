package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/businessFacades"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var WatchListRoutes = models.Routers{

	models.Router{
		Name:    "Create WatchList",
		Method:  "POST",
		Path:    "/api/watchList/save",
		Handler: businessFacades.CreateWatchList,
	},

}