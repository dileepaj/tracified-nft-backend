package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/api"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var WatchListRoutes = models.Routers{

	models.Router{
		Name:    "Create WatchList",
		Method:  "POST",
		Path:    "/api/watchList/save",
		Handler: api.CreateWatchList,
	},

}