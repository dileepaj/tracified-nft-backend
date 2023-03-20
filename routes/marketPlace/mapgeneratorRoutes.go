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
	models.Router{
		Name:    "Get Map",
		Method:  "GET",
		Path:    "/GetMap/{mapid}",
		Handler: apiHandler.GetMapByID,
	},
}
