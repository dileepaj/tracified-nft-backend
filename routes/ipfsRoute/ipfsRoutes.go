package ipfsRoutes

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var IpfsRoutes = models.Routers{
	models.Router{
		Name:    "Upload File to IPFS",
		Method:  "POST",
		Path:    "/api/ipfs",
		Handler: apiHandler.UploadFilesToIpfs,
	},
}
