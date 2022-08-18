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
		Path:    "/watchlists/save",
		Handler: apiHandler.CreateWatchList,
	},
	models.Router{
		Name:    "Get All Watchlists",
		Method:  "GET",
		Path:    "/watchlists",
		Handler: apiHandler.GetAllWatchLists,
	},
	models.Router{
		Name:    "Get WatchList By UserPK",
		Method:  "GET",
		Path:    "/watchlists/{user}",
		Handler: apiHandler.GetWatchListByUserPK,
	},
	models.Router{
		Name:    "Get WatchList By UserID",
		Method:  "GET",
		Path:    "/watchList/{currentownerpk}",
		Handler: apiHandler.GetWatchListNFT,
	},
	models.Router{
		Name:    "GET WatchLists By Blockchain",
		Method:  "Get",
		Path:    "/watchlists/{blockchain}/{nftidentifier}",
		Handler: apiHandler.FindWatchListsByBlockchainAndIdentifier,
	},
}
