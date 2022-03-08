package composer

import (
	"github.com/dileepaj/tracified-nft-backend/api"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var ComposerRoutes = models.Routers{

	models.Router{
		Name:    "Generate Html",
		Method:  "POST",
		Path:    "/api/composer/generate/nft",
		Handler: api.HTMLFileGenerator,
	},
}