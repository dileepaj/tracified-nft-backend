package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var EndorsementRoutes = models.Routers{
	models.Router{
		Name:    "Save Endorsement",
		Method:  "POST",
		Path:    "/api/endorser/save",
		Handler: apiHandler.CreateEndorsement,
	},
	models.Router{
		Name:    "GET Endorsement Status By PublicKey",
		Method:  "Get",
		Path:    "/api/endorsement/{publickey}",
		Handler: apiHandler.GetEndorsedStatus,
	},
	models.Router{
		Name:    "GET Endorsement By Status",
		Method:  "Get",
		Path:    "/api/endorsement/{status}",
		Handler: apiHandler.GetEndorsementbyStatus,
	},
	models.Router{
		Name:    "Update Endorsement Status",
		Method:  "PUT",
		Path:    "/api/endorsement/{publickey}",
		Handler: apiHandler.UpdateEndorsedStatus,
	},
	models.Router{
		Name:    "Update Endorsement",
		Method:  "PUT",
		Path:    "/api/endorsement",
		Handler: apiHandler.UpdateEndorsement,
	},
}
