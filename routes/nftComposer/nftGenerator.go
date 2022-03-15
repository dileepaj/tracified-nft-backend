package nftComposer

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var ComposerRoutes = models.Routers{
	models.Router{
		Name:    "Generate Html NFT",
		Method:  "POST",
		Path:    "/api/composer/generate/nft",
		Handler: apiHandler.HTMLFileGenerator,
	},
	models.Router{
		Name:    "Generated Html of NFT",
		Method:  "POST",
		Path:    "/api/composer/save/nft",
		Handler: apiHandler.SaveHTMLData,
	},
	models.Router{
		Name:    "Save the widget with otp and query",
		Method:  "POST",
		Path:    "/api/composer/save/widget",
		Handler: apiHandler.SaveWidget,
	},
	models.Router{
		Name:    "Query get OTP base on id and execute query",
		Method:  "POST",
		Path:    "/api/composer/query/execute",
		Handler: apiHandler.FindOtpAndExecuteQuery,
	},
}