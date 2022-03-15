package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var NftRoutes = models.Routers{

	models.Router{
		Name:    "Save NFT",
		Method:  "POST",
		Path:    "/api/nft/save",
		Handler: apiHandler.CreateNFT,
	},
	models.Router{
		Name:    "GET NFTS By Selling status",
		Method:  "Get",
		Path:    "/api/nfts/get/selling/{status}",
		Handler: apiHandler.GetAllONSaleNFT,
	},
	models.Router{
		Name:    "GET NFTS By Selling status and filter by UserPK",
		Method:  "Get",
		Path:    "/api/nfts/get/selling/{status}/{userpk}",
		Handler: apiHandler.GetAllONSaleNFT,
	},
	models.Router{
		Name:    "GET NFTS By Tags",
		Method:  "Get",
		Path:    "/api/nfts/get/tags/{tags}",
		Handler: apiHandler.GetNFTbyTags,
	},
	models.Router{
		Name:    "GET NFTS By Blockchain",
		Method:  "Get",
		Path:    "/api/nfts/get/blockchain/{blockchain}",
		Handler: apiHandler.GetNFTbyBlockChain,
	},
	models.Router{
		Name:    "GET Watch list NFTS By userId",
		Method:  "Get",
		Path:    "/api/nfts/get/watchlist/{userId}",
		Handler: apiHandler.GetNFTFromWatchList,
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
		Path:    "/api/nfts/get/tenetname/{tenentname}",
		Handler: apiHandler.GetNFTByTenentName,
	},
}