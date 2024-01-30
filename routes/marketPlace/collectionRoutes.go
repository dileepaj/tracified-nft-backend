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
		Name:    "GET Collection by Publickey",
		Method:  "Get",
		Path:    "/collection/{pubkey}",
		Handler: apiHandler.GetCollectionByPublicKey,
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
		Path:    "/collection",
		Handler: apiHandler.UpdateCollection,
	},
	models.Router{
		Name:    "Delete Collection by userPK",
		Method:  "DELETE",
		Path:    "/collection/userpk/{userid}",
		Handler: apiHandler.DeleteCollectionByUserPK,
	},
	models.Router{
		Name:    "GET Collection by UserID and Mail",
		Method:  "Get",
		Path:    "/collection/user/{userid}/{publickey}",
		Handler: apiHandler.GetCollectionByUserPKAndMail,
	},
	models.Router{
		Name:    "Update Collection Visibility",
		Method:  "PUT",
		Path:    "/collection-visibility",
		Handler: apiHandler.UpdateCollectionVisibility,
	},
	models.Router{
		Name:    "GET Collection by UserID",
		Method:  "Get",
		Path:    "/collection/owner/{objectid}",
		Handler: apiHandler.GetCollectionByEndorsementId,
	},
}
