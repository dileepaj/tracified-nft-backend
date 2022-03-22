package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var CollectionRoutes = models.Routers{
	models.Router{
		Name:    "Create Collection",
		Method:  "POST",
		Path:    "/api/collection/save",
		Handler: apiHandler.CreateCollection,
	},
	models.Router{
		Name:    "GET Collection by ID",
		Method:  "Get",
		Path:    "/api/collection/id/{_id}",
		Handler: apiHandler.GetCollectionById,
	},
	models.Router{
		Name:    "GET Collection by UserPK",
		Method:  "Get",
		Path:    "/api/collection/userpk/{userid}",
		Handler: apiHandler.GetCollectionByUserPK,
	},
	models.Router{
		Name:    "GET All Collections",
		Method:  "Get",
		Path:    "/api/collection",
		Handler: apiHandler.GetAllCollections,
	},
	models.Router{
		Name:    "Update Review Status",
		Method:  "PUT",
		Path:    "/api/collection",
		Handler: apiHandler.UpdateCollection,
	},
	models.Router{
		Name:    "Delete Collection by Id",
		Method:  "DELETE",
		Path:    "/api/collection/id/{_id}",
		Handler: apiHandler.DeleteCollectionById,
	},
	models.Router{
		Name:    "Delete Collection by userPK",
		Method:  "DELETE",
		Path:    "/api/collection/userpk/{userid}",
		Handler: apiHandler.DeleteCollectionByUserPK,
	},
}
