package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

// This routes manage the all NFT related rotes in the marketpalce
var NftRoutes = models.Routers{

	models.Router{
		Name:    "Save NFT",
		Method:  "POST",
		Path:    "/marketplace/save",
		Handler: apiHandler.CreateNFT,
	},
	models.Router{
		Name:    "GET NFT",
		Method:  "GET",
		Path:    "/marketplace",
		Handler: apiHandler.GetAllNFTs,
	},
	models.Router{
		Name:    "GET NFTS By Selling status and filter by currentownerpk",
		Method:  "Get",
		Path:    "/selling/{sellingstatus}/{currentownerpk}",
		Handler: apiHandler.GetAllONSaleNFT,
	},
	models.Router{
		Name:    "GET NFTS By Tag names",
		Method:  "Get",
		Path:    "/tags/{tags}",
		Handler: apiHandler.GetNFTbyTags,
	},
	models.Router{
		Name:    "GET NFTS By Status",
		Method:  "Get",
		Path:    "/nft/{sellingstatus}",
		Handler: apiHandler.GetNFTbyStatus,
	},
	models.Router{
		Name:    "GET NFTS By Blockchain",
		Method:  "Get",
		Path:    "/blockchain/{blockchain}",
		Handler: apiHandler.GetBlockchainSpecificNFT,
	},
	models.Router{
		Name:    "GET NFTS By currentownerpk",
		Method:  "Get",
		Path:    "/userid/{currentownerpk}",
		Handler: apiHandler.GetNFTByUserId,
	},
	models.Router{
		Name:    "GET Last NFT By userId",
		Method:  "Get",
		Path:    "/userid/{creatoruserid}",
		Handler: apiHandler.GetLastNFTByUserId,
	},
	models.Router{
		Name:    "GET SVG By Hash",
		Method:  "Get",
		Path:    "/svg/{hash}",
		Handler: apiHandler.GetSVGBySHA256,
	},
	models.Router{
		Name:    "GET NFTS By tenent Name",
		Method:  "Get",
		Path:    "/tenentname/{creatoruserid}",
		Handler: apiHandler.GetNFTByTenentName,
	},
	models.Router{
		Name:    "GET Tags by NftName",
		Method:  "Get",
		Path:    "/tags/nft/{nftName}",
		Handler: apiHandler.GetTagsByNFTName,
	},
	models.Router{
		Name:    "GET All Tags",
		Method:  "Get",
		Path:    "/tags",
		Handler: apiHandler.GetAllTags,
	},
	models.Router{
		Name:    "Update Sale Status",
		Method:  "PUT",
		Path:    "/nft/sale",
		Handler: apiHandler.MakeSale,
	},
	models.Router{
		Name:    "Update SVG",
		Method:  "PUT",
		Path:    "/svg",
		Handler: apiHandler.UpdateSvgBlockChain,
	},
	models.Router{
		Name:    "Create SVG",
		Method:  "POST",
		Path:    "/svg/save",
		Handler: apiHandler.CreateSVG,
	},
	models.Router{
		Name:    "Create TXN",
		Method:  "POST",
		Path:    "/txn/save",
		Handler: apiHandler.SaveTXN,
	},
	models.Router{
		Name:    "Get NFT By Blockchain And UserPK",
		Method:  "Get",
		Path:    "/nft/{currentownerpk}/{blockchain}",
		Handler: apiHandler.GetNFTByBlockchainAndUserPK,
	},
	models.Router{
		Name:    "Get NFT By Blockchain",
		Method:  "Get",
		Path:    "/nft/{blockchain}",
		Handler: apiHandler.GetNFTByBlockchain,
	},
	models.Router{
		Name:    "Get TXN By Blockchain And NftIdentifier",
		Method:  "Get",
		Path:    "/txn/{nftidentifier}/{blockchain}",
		Handler: apiHandler.GetTXNByBlockchainAndIdentifier,
	},
	models.Router{
		Name:    "Save Tags",
		Method:  "POST",
		Path:    "/tags/save",
		Handler: apiHandler.CreateTags,
	},
	models.Router{
		Name:    "Update Minter on Solana",
		Method:  "PUT",
		Path:    "/marketplace/nft",
		Handler: apiHandler.UpdateMinter,
	},
	models.Router{
		Name:    "Update TXN for Stellar",
		Method:  "PUT",
		Path:    "/marketplace/txn",
		Handler: apiHandler.UpdateTXN,
	},
	models.Router{
		Name:    "Save Ownership",
		Method:  "POST",
		Path:    "/marketplace/owner",
		Handler: apiHandler.CreateOwner,
	},
	models.Router{
		Name:    "GET NFTS By Selling status NFTIdentifier and Blockchain",
		Method:  "Get",
		Path:    "/buying/{sellingstatus}/{nftidentifier}/{blockchain}",
		Handler: apiHandler.GetOneONSaleNFT,
	},
	models.Router{
		Name:    "Save NFT Story",
		Method:  "POST",
		Path:    "/story/",
		Handler: apiHandler.SaveNFTStory,
	},
	models.Router{
		Name:    "Get NFT Story by NFTIdentifier and Blockchain",
		Method:  "Get",
		Path:    "/story/{nftidentifier}/{blockchain}",
		Handler: apiHandler.GetNFTStory,
	},
	models.Router{
		Name:    "Get NFT By Collection Name",
		Method:  "Get",
		Path:    "/nftcollection/{collection}",
		Handler: apiHandler.GetNFTByCollection,
	},
	models.Router{
		Name:    "Get NFT pagination",
		Method:  "Get",
		Path:    "/nftspaginate/{blockchain}/{pagesize}/{requestedPage}",
		Handler: apiHandler.GetPaginatedNFTs,
	},
	//another route for getting nfts should be here but there are two functions for it already
}
