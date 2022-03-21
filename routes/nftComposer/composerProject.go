package nftComposer

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

// This routes handle the all html generating rouites in the nft composer
var ComposerRoutes = models.Routers{
	models.Router{
		Name:    "Generate Html NFT",
		Method:  "POST",
		Path:    "/api/generate",
		Handler: apiHandler.HTMLFileGenerator,
	},
	models.Router{
		Name:    "Generated Html of NFT",
		Method:  "POST",
		Path:    "/api/project",
		Handler: apiHandler.SaveProject,
	},
	models.Router{
		Name:    "Save the widget with otp and query",
		Method:  "POST",
		Path:    "/api/widget",
		Handler: apiHandler.SaveWidget,
	},
	models.Router{
		Name:    "Query get OTP base on id and execute query",
		Method:  "POST",
		Path:    "/api/query/execute",
		Handler: apiHandler.QueryExecuter,
	},
	models.Router{
		Name:    "Get All Project base on user",
		Method:  "GET",
		Path:    "/api/projcts/{userId}",
		Handler: apiHandler.GetRecentProjects,
	},
		models.Router{
		Name:    "Get All Project base on user",
		Method:  "GET",
		Path:    "/api/projct/{projectId}",
		Handler: apiHandler.GetRecentProjectDetails,
	},
}
