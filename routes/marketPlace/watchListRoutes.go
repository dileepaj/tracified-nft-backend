package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

//This routes handle the all watchlist  related routes in the marketplace
var WatchListRoutes = models.Routers{

	models.Router{
		Name:    "Create WatchList",
		Method:  "POST",
		Path:    "/api/watchList/save",
		Handler: apiHandler.CreateWatchList,
	},

}