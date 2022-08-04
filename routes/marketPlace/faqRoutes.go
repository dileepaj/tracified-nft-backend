package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var FaqRoutes = models.Routers{
	//Route will be used to add a new FAQ to DB
	models.Router{
		Name:    "Create FAQ",
		Method:  "POST",
		Path:    "/faq/",
		Handler: apiHandler.CreateFaq,
	},
	//Will return all the Faq in collection
	models.Router{
		Name:    "Get All FAQ",
		Method:  "GET",
		Path:    "/faq/",
		Handler: apiHandler.GetAllFaq,
	},
	//Will return FAQ based on Question ID provided
	models.Router{
		Name:    "Get FAQ by QuestionID",
		Method:  "GET",
		Path:    "/faq/{_id}",
		Handler: apiHandler.GetFaqByID,
	},
	// Will be used to update FAQ
	models.Router{
		Name:    "Update FAQ",
		Method:  "PUT",
		Path:    "/faq/",
		Handler: apiHandler.UpdateFaqbyID,
	},
}
