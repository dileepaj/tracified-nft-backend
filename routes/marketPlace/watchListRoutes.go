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
	models.Router{
		Name:    "Get All Watchlists",
		Method:  "GET",
		Path:    "/api/watchList",
		Handler: apiHandler.GetAllWatchLists,
	},
	models.Router{
		Name:    "Get WatchList By UserPK",
		Method:  "GET",
		Path:    "/api/watchList/{userpk}",
		Handler: apiHandler.GetWatchListByUserPK,
	},
	models.Router{
		Name:    "Get WatchList By UserID",
		Method:  "GET",
		Path:    "/api/watchList/{userid}",
		Handler: apiHandler.GetWatchListNFT,
	},
}
