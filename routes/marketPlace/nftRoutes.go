package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/api"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var NftRoutes = models.Routers{

	models.Router{
		Name:    "Save NFT",
		Method:  "POST",
		Path:    "/api/nft/save",
		Handler: api.CreateNFT,
	},
	models.Router{
		Name:    "GET NFTS By Selling status",
		Method:  "Get",
		Path:    "/api/nfts/get/selling/{status}",
		Handler: api.GetAllONSaleNFT,
	},
	models.Router{
		Name:    "GET NFTS By Selling status and filter by UserPK",
		Method:  "Get",
		Path:    "/api/nfts/get/selling/{status}/{userpk}",
		Handler: api.GetAllONSaleNFT,
	},
	models.Router{
		Name:    "GET NFTS By Tags",
		Method:  "Get",
		Path:    "/api/nfts/get/tags/{tags}",
		Handler: api.GetNFTbyTags,
	},
	models.Router{
		Name:    "GET NFTS By Blockchain",
		Method:  "Get",
		Path:    "/api/nfts/get/blockchain/{blockchain}",
		Handler: api.GetNFTbyBlockChain,
	},
	models.Router{
		Name:    "GET Watch list NFTS By userId",
		Method:  "Get",
		Path:    "/api/nfts/get/watchlist/{userId}",
		Handler: api.GetNFTFromWatchList,
	},
	models.Router{
		Name:    "GET NFTS By userId",
		Method:  "Get",
		Path:    "/api/nfts/get/userid/{userId}",
		Handler: api.GetNFTByUserId,
	},
	models.Router{
		Name:    "GET NFTS By tenent Name",
		Method:  "Get",
		Path:    "/api/nfts/get/tenetname/{tenentname}",
		Handler: api.GetNFTByTenentName,
	},
}