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
	//Route will be used to store user questions
	models.Router{
		Name:    "Store User FAQ",
		Method:  "POST",
		Path:    "/userfaq/",
		Handler: apiHandler.StoreFAQ,
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
	models.Router{
		Name:    "GET User FAQ By Status",
		Method:  "Get",
		Path:    "/userfaq/status/{status}",
		Handler: apiHandler.GetUserFAQbyStatus,
	},
	models.Router{
		Name:    "Update UserFAQ Status",
		Method:  "PUT",
		Path:    "/userfaq/status",
		Handler: apiHandler.UpdateUserFAQStatus,
	},
	models.Router{
		Name:    "GET FAQ attachment by QID",
		Method:  "GET",
		Path:    "/userfaq/attachment/{qid}",
		Handler: apiHandler.GetUserFAQAttachmentbyID,
	},
}
