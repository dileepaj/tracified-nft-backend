package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var FavouritesRoutes = models.Routers{
	models.Router{
		Name:    "Create Favourites",
		Method:  "POST",
		Path:    "/api/favourites/save",
		Handler: apiHandler.CreateFavourites,
	},
	models.Router{
		Name:    "Get All Favourites",
		Method:  "GET",
		Path:    "/favourites",
		Handler: apiHandler.GetAllFavourites,
	},
	models.Router{
		Name:    "Get Favourites By UserPK",
		Method:  "GET",
		Path:    "/favourites/{user}",
		Handler: apiHandler.GetFavouritesByUserPK,
	},
	models.Router{
		Name:    "GET Favourites By Favourites",
		Method:  "Get",
		Path:    "/favourites/{blockchain}/{nftidentifier}",
		Handler: apiHandler.GetFavouritesByBlockchainAndIdentifier,
	},
}
