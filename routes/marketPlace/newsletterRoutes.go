package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

// This routes handle the all watchlist  related routes in the marketplace
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
		Path:    "/newsletter/author/{name}/",
		Handler: apiHandler.GetNewslettersByAuthor,
	},
	models.Router{
		Name:    "get NewsLetter by ID",
		Method:  "GET",
		Path:    "/newsletter/{_id}/",
		Handler: apiHandler.GetNewsletterByID,
	},
	models.Router{
		Name:    "Subscription",
		Method:  "POST",
		Path:    "/subscribe/",
		Handler: apiHandler.Subscribe,
	},
	models.Router{
		Name:    "Check Subscription",
		Method:  "GET",
		Path:    "/subscribe/check/{email}",
		Handler: apiHandler.CheckIfSubscribed,
	},
}
