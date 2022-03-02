package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/businessFacades"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var NftRoutes = models.Routers{

	models.Router{
		Name:    "Save NFT",
		Method:  "POST",
		Path:    "/api/nft/save",
		Handler: businessFacades.CreateNFT,
	},
	models.Router{
		Name:    "GET NFTS By Selling status",
		Method:  "Get",
		Path:    "/api/nfts/get/selling/{status}",
		Handler: businessFacades.GetAllONSaleNFT,
	},
	models.Router{
		Name:    "GET NFTS By Selling status and filter by UserPK",
		Method:  "Get",
		Path:    "/api/nfts/get/selling/{status}/{userpk}",
		Handler: businessFacades.GetAllONSaleNFT,
	},
	models.Router{
		Name:    "GET NFTS By Tags",
		Method:  "Get",
		Path:    "/api/nfts/get/tags/{tags}",
		Handler: businessFacades.GetNFTbyTags,
	},
	models.Router{
		Name:    "GET NFTS By Blockchain",
		Method:  "Get",
		Path:    "/api/nfts/get/blockchain/{blockchain}",
		Handler: businessFacades.GetNFTbyBlockChain,
	},
	models.Router{
		Name:    "GET Watch list NFTS By userId",
		Method:  "Get",
		Path:    "/api/nfts/get/watchlist/{userId}",
		Handler: businessFacades.GetNFTFromWatchList,
	},
	models.Router{
		Name:    "GET NFTS By userId",
		Method:  "Get",
		Path:    "/api/nfts/get/userid/{userId}",
		Handler: businessFacades.GetNFTByUserId,
	},
	models.Router{
		Name:    "GET NFTS By tenent Name",
		Method:  "Get",
		Path:    "/api/nfts/get/tenetname/{tenentname}",
		Handler: businessFacades.GetNFTByTenentName,
	},
}