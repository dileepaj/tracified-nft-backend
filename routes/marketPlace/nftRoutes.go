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
		Path:    "/api/marketplace/save",
		Handler: apiHandler.CreateNFT,
	},
	models.Router{
		Name:    "GET NFTS By Selling status and filter by UserPK",
		Method:  "Get",
		Path:    "/api/selling/{sellingstatus}/{currentownerpk}",
		Handler: apiHandler.GetAllONSaleNFT,
	},
	models.Router{
		Name:    "GET NFTS By Tag names",
		Method:  "Get",
		Path:    "/api/tags/{tags}",
		Handler: apiHandler.GetNFTbyTags,
	},
	models.Router{
		Name:    "GET NFTS By Status",
		Method:  "Get",
		Path:    "/api/nft/{sellingstatus}",
		Handler: apiHandler.GetNFTbyStatus,
	},
	models.Router{
		Name:    "GET NFTS By Blockchain",
		Method:  "Get",
		Path:    "/api/blockchain/{blockchain}",
		Handler: apiHandler.GetBlockchainSpecificNFT,
	},
	models.Router{
		Name:    "GET Watch list NFTS By userId",
		Method:  "Get",
		Path:    "/api/watchlist/{userId}",
		Handler: apiHandler.GetWatchListNFT,
	},
	models.Router{
		Name:    "GET NFTS By userId",
		Method:  "Get",
		Path:    "/api/userid/{userId}",
		Handler: apiHandler.GetNFTByUserId,
	},
	models.Router{
		Name:    "GET Last NFT By userId",
		Method:  "Get",
		Path:    "/api/userid/{creatoruserid}",
		Handler: apiHandler.GetLastNFTByUserId,
	},
	models.Router{
		Name:    "GET SVG By Hash",
		Method:  "Get",
		Path:    "/api/svg/{hash}",
		Handler: apiHandler.GetSVGBySHA256,
	},
	models.Router{
		Name:    "GET NFTS By tenent Name",
		Method:  "Get",
		Path:    "/api/tenentname/{tenentname}",
		Handler: apiHandler.GetNFTByTenentName,
	},
	models.Router{
		Name:    "GET Tags by NftIdentifier",
		Method:  "Get",
		Path:    "/api/tags/nft/{nftidentifier}",
		Handler: apiHandler.GetTagsByNFTIdentifier,
	},
	models.Router{
		Name:    "GET All Collections",
		Method:  "Get",
		Path:    "/api/tags",
		Handler: apiHandler.GetAllTags,
	},
	models.Router{
		Name:    "Update Sale Status",
		Method:  "PUT",
		Path:    "/api/nft/sale",
		Handler: apiHandler.MakeSale,
	},
	models.Router{
		Name:    "Update SVG",
		Method:  "PUT",
		Path:    "/api/svg",
		Handler: apiHandler.UpdateSvgBlockChain,
	},
	models.Router{
		Name:    "Create SVG",
		Method:  "POST",
		Path:    "/api/svg/save",
		Handler: apiHandler.CreateSVG,
	},
	models.Router{
		Name:    "Create TXN",
		Method:  "POST",
		Path:    "/api/txn/save",
		Handler: apiHandler.SaveTXN,
	},
	models.Router{
		Name:    "Get NFT By Blockchain And UserPK",
		Method:  "Get",
		Path:    "/api/nft/{userpk}/{blockchain}",
		Handler: apiHandler.GetNFTByBlockchainAndUserPK,
	},
	models.Router{
		Name:    "Save Tags",
		Method:  "POST",
		Path:    "/api/tags/save",
		Handler: apiHandler.CreateTags,
	},
	models.Router{
		Name:    "Update Minter",
		Method:  "PUT",
		Path:    "/api/marketplace/nft",
		Handler: apiHandler.UpdateMinter,
	},
	models.Router{
		Name:    "Update TXN",
		Method:  "PUT",
		Path:    "/api/marketplace/txn",
		Handler: apiHandler.UpdateTXN,
	},
	models.Router{
		Name:    "Save Ownership",
		Method:  "POST",
		Path:    "/api/marketplace/owner",
		Handler: apiHandler.CreateOwner,
	},
	models.Router{
		Name:    "GET NFTS By Selling status and filter by NFTIdentifier",
		Method:  "Get",
		Path:    "/api/buying/{sellingstatus}/{nftidentifier}/{blockchain}",
		Handler: apiHandler.GetOneONSaleNFT,
	},
}
