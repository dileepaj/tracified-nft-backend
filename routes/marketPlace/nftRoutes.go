package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

//This routes manage the all NFT related rotes in the marketpalce
var NftRoutes = models.Routers{

	models.Router{
		Name:    "Save NFT",
		Method:  "POST",
		Path:    "/api/nft/save",
		Handler: apiHandler.CreateNFT,
	},
	models.Router{
		Name:    "GET NFTS By Selling status and filter by UserPK",
		Method:  "Get",
		Path:    "/api/nfts/get/selling/{status}/{userpk}",
		Handler: apiHandler.GetAllONSaleNFT,
	},
	models.Router{
		Name:    "GET NFTS By Tag names",
		Method:  "Get",
		Path:    "/api/nfts/get/tags/{tags}",
		Handler: apiHandler.GetNFTbyTags,
	},
	models.Router{
		Name:    "GET NFTS By Blockchain",
		Method:  "Get",
		Path:    "/api/nfts/get/blockchain/{blockchain}",
		Handler: apiHandler.GetBlockchainSpecificNFT,
	},
	models.Router{
		Name:    "GET Watch list NFTS By userId",
		Method:  "Get",
		Path:    "/api/nfts/get/watchlist/{userId}",
		Handler: apiHandler.GetWatchListNFT,
	},
	models.Router{
		Name:    "GET NFTS By userId",
		Method:  "Get",
		Path:    "/api/nfts/get/userid/{userId}",
		Handler: apiHandler.GetNFTByUserId,
	},
	models.Router{
		Name:    "GET NFTS By tenent Name",
		Method:  "Get",
		Path:    "/api/nfts/get/tenentname/{tenentname}",
		Handler: apiHandler.GetNFTByTenentName,
	},
}