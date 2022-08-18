package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var CollectionRoutes = models.Routers{
	models.Router{
		Name:    "Create Collection",
		Method:  "POST",
		Path:    "/collection/save",
		Handler: apiHandler.CreateCollection,
	},
	models.Router{
		Name:    "GET Collection by UserID",
		Method:  "Get",
		Path:    "/collection/userpk/{userid}",
		Handler: apiHandler.GetCollectionByUserPK,
	},
	models.Router{
		Name:    "GET All Collections",
		Method:  "Get",
		Path:    "/collection",
		Handler: apiHandler.GetAllCollections,
	},
	models.Router{
		Name:    "Update Collections",
		Method:  "PUT",
		Path:    "/collections",
		Handler: apiHandler.UpdateCollection,
	},
	models.Router{
		Name:    "Delete Collection by userPK",
		Method:  "DELETE",
		Path:    "/collection/userpk/{userid}",
		Handler: apiHandler.DeleteCollectionByUserPK,
	},
}
