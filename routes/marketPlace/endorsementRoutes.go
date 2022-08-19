package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var EndorsementRoutes = models.Routers{
	models.Router{
		Name:    "Save Endorsement",
		Method:  "POST",
		Path:    "/endorser/save",
		Handler: apiHandler.CreateEndorsement,
	},
	models.Router{
		Name:    "GET Endorsement Status By PublicKey",
		Method:  "Get",
		Path:    "/endorsement/{publickey}",
		Handler: apiHandler.GetEndorsedStatus,
	},
	models.Router{
		Name:    "GET Endorsement By Status",
		Method:  "Get",
		Path:    "/endorsement/status/{status}",
		Handler: apiHandler.GetEndorsementbyStatus,
	},
	models.Router{
		Name:    "Update Endorsement Status",
		Method:  "PUT",
		Path:    "/endorsement/status",
		Handler: apiHandler.UpdateEndorsedStatus,
	},
	models.Router{
		Name:    "Update Endorsement",
		Method:  "PUT",
		Path:    "/endorsement",
		Handler: apiHandler.UpdateEndorsement,
	},
}
