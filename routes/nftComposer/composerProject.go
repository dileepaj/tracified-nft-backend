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
		Name:    "Save Project",
		Method:  "POST",
		Path:    "/api/project",
		Handler: apiHandler.SaveProject,
	},
	models.Router{
		Name:    "Update Project",
		Method:  "PUT",
		Path:    "/api/project",
		Handler: apiHandler.UpdateProject,
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
		Path:    "/api/projects/{userId}",
		Handler: apiHandler.GetRecentProjects,
	},
	models.Router{
		Name:    "Get All Project base on user",
		Method:  "GET",
		Path:    "/api/project/{projectId}",
		Handler: apiHandler.GetRecentProjectDetails,
	},
	models.Router{
		Name:    "Update Widget",
		Method:  "PUT",
		Path:    "/api/widget",
		Handler: apiHandler.UpdateWidget,
	},
	models.Router{
		Name:    "Save Chart",
		Method:  "POST",
		Path:    "/api/html/chart",
		Handler: apiHandler.SaveChart,
	},
	models.Router{
		Name:    "Save Table",
		Method:  "POST",
		Path:    "/api/html/table",
		Handler: apiHandler.SaveTable,
	},
	models.Router{
		Name:    "Save Proofbot",
		Method:  "POST",
		Path:    "/api/html/proofbot",
		Handler: apiHandler.SaveProofBot,
	},
	models.Router{
		Name:    "Save Stat",
		Method:  "POST",
		Path:    "/api/html/stat",
		Handler: apiHandler.SaveStat,
	},
	models.Router{
		Name:    "Save Image",
		Method:  "POST",
		Path:    "/api/html/image",
		Handler: apiHandler.SaveImage,
	},
	// update
	models.Router{
		Name:    "Update Chart",
		Method:  "PUT",
		Path:    "/api/html/chart",
		Handler: apiHandler.UpdateChart,
	},
	models.Router{
		Name:    "Update Table",
		Method:  "PUT",
		Path:    "/api/html/table",
		Handler: apiHandler.UpdateTable,
	},
	models.Router{
		Name:    "Update Proofbot",
		Method:  "PUT",
		Path:    "/api/html/proofbot",
		Handler: apiHandler.UpdateProofBot,
	},
	models.Router{
		Name:    "Update Stats",
		Method:  "PUT",
		Path:    "/api/html/stats",
		Handler: apiHandler.UpdateStats,
	},
	models.Router{
		Name:    "Update Image",
		Method:  "PUT",
		Path:    "/api/html/image",
		Handler: apiHandler.UpdateImage,
	},
	// remove
	models.Router{
		Name:    "Remove Project",
		Method:  "DELETE",
		Path:    "/api/html/project/{projectId}",
		Handler: apiHandler.RemoveProjet,
	},
	models.Router{
		Name:    "Remove Chart",
		Method:  "DELETE",
		Path:    "/api/html/chart/{widgetId}",
		Handler: apiHandler.RemoveChart,
	},
	models.Router{
		Name:    "Remove Table",
		Method:  "DELETE",
		Path:    "/api/html/table/{widgetId",
		Handler: apiHandler.RemoveTable,
	},
	models.Router{
		Name:    "Remove Proofbot",
		Method:  "DELETE",
		Path:    "/api/html/proofbot/{widgetId",
		Handler: apiHandler.RemoveProofBot,
	},
	models.Router{
		Name:    "Remove Stats",
		Method:  "DELETE",
		Path:    "/api/html/stats/{widgetId",
		Handler: apiHandler.RemoveStats,
	},
	models.Router{
		Name:    "Remove Image",
		Method:  "DELETE",
		Path:    "/api/html/image/{widgetId",
		Handler: apiHandler.RemoveImage,
	},
}
