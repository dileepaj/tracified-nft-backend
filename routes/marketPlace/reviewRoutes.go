package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var ReviewRoutes = models.Routers{

	models.Router{
		Name:    "Create Review",
		Method:  "POST",
		Path:    "/api/review/save",
		Handler: apiHandler.CreateReview,
	},
	models.Router{
		Name:    "GET Review by nftidentifier",
		Method:  "GET",
		Path:    "/api/review/{nftidentifier}",
		Handler: apiHandler.GetNFTReviewByNFT,
	},
	models.Router{
		Name:    "GET All Reviews",
		Method:  "GET",
		Path:    "/api/review",
		Handler: apiHandler.GetAllReviews,
	},
	models.Router{
		Name:    "Update Review Status",
		Method:  "PUT",
		Path:    "/api/review/updateStatus",
		Handler: apiHandler.UpdateReviewStatus,
	},
	models.Router{
		Name:    "Delete Review",
		Method:  "DELETE",
		Path:    "/api/review/delete",
		Handler: apiHandler.DeleteReview,
	},
}
