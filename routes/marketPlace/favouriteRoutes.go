package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var FavouritesRoutes = models.Routers{
	models.Router{
		Name:    "Create Favourites",
		Method:  "POST",
		Path:    "/favourites/save",
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
		Name:    "GETand PUT Favourites By Favourites",
		Method:  "GET",
		Path:    "/favourites/{blockchain}/{nftidentifier}",
		Handler: apiHandler.GetFavouritesByBlockchainAndIdentifier,
	},
	models.Router{
		Name:    "GET Favourites By Favourites",
		Method:  "GET",
		Path:    "/favourite/{blockchain}/{nftidentifier}",
		Handler: apiHandler.FavouritesByBlockchainAndIdentifier,
	},
	models.Router{
		Name:    "verify favourite count By UserPK",
		Method:  "GET",
		Path:    "/verify/favouriteCount/{blockchain}/{user}/{nftidentifer}",
		Handler: apiHandler.VerifyFavouriteTogglebUserPK,
	},
	models.Router{
		Name:    "Remove user from favourite",
		Method:  "DELETE",
		Path:    "/favourite/{favouriteID}",
		Handler: apiHandler.RemoveUserfromFavourite,
	},
}
