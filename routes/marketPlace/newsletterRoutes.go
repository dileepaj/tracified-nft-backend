package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

//This routes handle the all watchlist  related routes in the marketplace
var NewsLetterRoutes = models.Routers{

	models.Router{
		Name:    "Create NewsLetter",
		Method:  "POST",
		Path:    "/newsletter/",
		Handler: apiHandler.CreateNewsLetter,
	},
	models.Router{
		Name:    "Get All NewsLetter",
		Method:  "GET",
		Path:    "/newsletter/",
		Handler: apiHandler.GetAllNewsLetters,
	},
	models.Router{
		Name:    "get NewsLetter by Auther",
		Method:  "GET",
		Path:    "/newsletter/{name}",
		Handler: apiHandler.GetNewslettersByAuthor,
	},
}
