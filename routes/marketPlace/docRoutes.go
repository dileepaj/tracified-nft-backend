package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var DocsRoutes = models.Routers{
	//Route will be used to add a new FAQ to DB
	models.Router{
		Name:    "Create Docs",
		Method:  "POST",
		Path:    "/docs/",
		Handler: apiHandler.CreateDocs,
	},
	//Will return all the Faq in collection
	models.Router{
		Name:    "Get All Docs",
		Method:  "GET",
		Path:    "/docs/",
		Handler: apiHandler.GetAllDocs,
	},
	//Will return FAQ based on Question ID provided
	models.Router{
		Name:    "Get Docs by DocsID",
		Method:  "GET",
		Path:    "/docs/{_id}",
		Handler: apiHandler.GetDocsByID,
	},
	// Will be used to update FAQ
	models.Router{
		Name:    "Update Docs",
		Method:  "PUT",
		Path:    "/docs/",
		Handler: apiHandler.UpdateDocsbyID,
	},
}
