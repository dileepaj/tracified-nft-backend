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
		Path:    "/generate",
		Handler: apiHandler.HTMLFileGenerator,
	},
	models.Router{
		Name:    "Generate SVG of NFT",
		Method:  "POST",
		Path:    "/generate/svg",
		Handler: apiHandler.SVGFileGenerator,
	},
	models.Router{
		Name:    "Save Project",
		Method:  "POST",
		Path:    "/project",
		Handler: apiHandler.SaveProject,
	},
	models.Router{
		Name:    "Update Project",
		Method:  "PUT",
		Path:    "/project",
		Handler: apiHandler.UpdateProject,
	},
	models.Router{
		Name:    "Save the widget with otp and query",
		Method:  "POST",
		Path:    "/widget",
		Handler: apiHandler.SaveWidget,
	},
	models.Router{
		Name:    "Query get OTP base on id and execute query",
		Method:  "POST",
		Path:    "/query/execute",
		Handler: apiHandler.QueryExecuter,
	},
	models.Router{
		Name:    "Get All Project base on user",
		Method:  "GET",
		Path:    "/projects/{userId}",
		Handler: apiHandler.GetRecentProjects,
	},
	models.Router{
		Name:    "Get All Project base on user",
		Method:  "GET",
		Path:    "/project/{projectId}",
		Handler: apiHandler.GetRecentProjectDetails,
	},
	models.Router{
		Name:    "Update Widget",
		Method:  "PUT",
		Path:    "/widget",
		Handler: apiHandler.UpdateWidget,
	},
	models.Router{
		Name:    "Save Chart",
		Method:  "POST",
		Path:    "/html/chart",
		Handler: apiHandler.SaveChart,
	},
	models.Router{
		Name:    "Save Table",
		Method:  "POST",
		Path:    "/html/table",
		Handler: apiHandler.SaveTable,
	},
	models.Router{
		Name:    "Save Proofbot",
		Method:  "POST",
		Path:    "/html/proofbot",
		Handler: apiHandler.SaveProofBot,
	},
	models.Router{
		Name:    "Save Stat",
		Method:  "POST",
		Path:    "/html/stat",
		Handler: apiHandler.SaveStat,
	},
	models.Router{
		Name:    "Save Image",
		Method:  "POST",
		Path:    "/html/image",
		Handler: apiHandler.SaveImage,
	},
	models.Router{
		Name:    "Save TimeLine",
		Method:  "POST",
		Path:    "/html/timeline",
		Handler: apiHandler.SaveTimeline,
	},
	// update
	models.Router{
		Name:    "Update Chart",
		Method:  "PUT",
		Path:    "/html/chart",
		Handler: apiHandler.UpdateChart,
	},
	models.Router{
		Name:    "Update Table",
		Method:  "PUT",
		Path:    "/html/table",
		Handler: apiHandler.UpdateTable,
	},
	models.Router{
		Name:    "Update Proofbot",
		Method:  "PUT",
		Path:    "/html/proofbot",
		Handler: apiHandler.UpdateProofBot,
	},
	models.Router{
		Name:    "Update Stats",
		Method:  "PUT",
		Path:    "/html/stats",
		Handler: apiHandler.UpdateStats,
	},
	models.Router{
		Name:    "Update Image",
		Method:  "PUT",
		Path:    "/html/image",
		Handler: apiHandler.UpdateImage,
	},
		models.Router{
		Name:    "Update Timeline",
		Method:  "PUT",
		Path:    "/html/timeline",
		Handler: apiHandler.UpdateTimeline,
	},
	// remove
	models.Router{
		Name:    "Remove Project",
		Method:  "DELETE",
		Path:    "/html/project/{projectId}",
		Handler: apiHandler.RemoveProjet,
	},
	models.Router{
		Name:    "Remove Chart",
		Method:  "DELETE",
		Path:    "/html/chart/{widgetId}",
		Handler: apiHandler.RemoveChart,
	},
	models.Router{
		Name:    "Remove Table",
		Method:  "DELETE",
		Path:    "/html/table/{widgetId}",
		Handler: apiHandler.RemoveTable,
	},
	models.Router{
		Name:    "Remove Proofbot",
		Method:  "DELETE",
		Path:    "/html/proofbot/{widgetId}",
		Handler: apiHandler.RemoveProofBot,
	},
	models.Router{
		Name:    "Remove Stats",
		Method:  "DELETE",
		Path:    "/html/stats/{widgetId}",
		Handler: apiHandler.RemoveStats,
	},
	models.Router{
		Name:    "Remove Image",
		Method:  "DELETE",
		Path:    "/html/image/{widgetId}",
		Handler: apiHandler.RemoveImage,
	},
		models.Router{
		Name:    "Remove Tilemine",
		Method:  "DELETE",
		Path:    "/html/timeline/{widgetId}",
		Handler: apiHandler.RemoveTimeline,
	},
}
