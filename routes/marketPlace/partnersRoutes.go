package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var PartnerRoutes = models.Routers{
	//Route will be used to add a new FAQ to DB
	models.Router{
		Name:    "Create Partner",
		Method:  "POST",
		Path:    "/partner/",
		Handler: apiHandler.CreatePartner,
	},
	//Will return all the Faq in collection
	models.Router{
		Name:    "Get All Partner",
		Method:  "GET",
		Path:    "/partner/",
		Handler: apiHandler.GetAllPartner,
	},
	//Will return FAQ based on Question ID provided
	models.Router{
		Name:    "Get Partner by PartnerID",
		Method:  "GET",
		Path:    "/partner/{_id}",
		Handler: apiHandler.GetPartnerByID,
	},
	// Will be used to update FAQ
	models.Router{
		Name:    "Update Partner",
		Method:  "PUT",
		Path:    "/partner/",
		Handler: apiHandler.UpdatePartnerbyID,
	},
}
