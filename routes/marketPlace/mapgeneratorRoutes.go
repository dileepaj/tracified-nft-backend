package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var MapGenerateRoutes = models.Routers{
	models.Router{
		Name:    "Generate Map",
		Method:  "POST",
		Path:    "/generatemap",
		Handler: apiHandler.GenerateMap,
	},
}
