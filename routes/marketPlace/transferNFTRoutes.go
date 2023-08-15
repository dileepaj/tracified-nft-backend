package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var WalletNFTRoutes = models.Routers{
	models.Router{
		Name:    "Update wallet nft state",
		Method:  "PUT",
		Path:    "/walletnft/state",
		Handler: apiHandler.UpdateWalletNFTState,
	},
	models.Router{
		Name:    "Save wallet nft state txns",
		Method:  "POST",
		Path:    "/walletnft/txns",
		Handler: apiHandler.SaveWalletNFTTXNs,
	},
	models.Router{
		Name:    "Save Wallet NFT State",
		Method:  "POST",
		Path:    "/walletnft/state",
		Handler: apiHandler.SaveWalletNFTStates,
	},
	// ?blockchain=stellar&nftstatus=1&currentowner=abc&pagesize=1&requestedPage=1
	models.Router{
		Name:   "Get Wallet Users NFT By State",
		Method: "Get",
		Path:   "/walletnft",
		// Path:    "/walletnft/{blockchain}/{nftstatus}/{currentowner}/{pagesize}/{requestedPage}",
		Handler: apiHandler.GetWalletNFTByStateAndCurrentOwner,
	},
	// ?blockchain=stellar&nftstatus=1&nftrequested=abc&pagesize=1&requestedPage=1
	models.Router{
		Name:   "Get Wallet Users NFT Requestee By State",
		Method: "Get",
		Path:   "/walletnft/requested",
		// Path:    "/walletnft/requested/{blockchain}/{nftstatus}/{nftrequested}/{pagesize}/{requestedPage}",
		Handler: apiHandler.GetWalletNFTByStatusAndRequested,
	},
	models.Router{
		Name:    "Delete NFT State by userPK",
		Method:  "DELETE",
		Path:    "/walletnft/state/{issuerpublickey}",
		Handler: apiHandler.DeleteWalletNFTByIssuer,
	},
	models.Router{
		Name:    "GET NFT State TXN by Issuer",
		Method:  "Get",
		Path:    "/nftstate/txns/{issuerpublickey}",
		Handler: apiHandler.GetWalletTxnsByIssuer,
	},
	models.Router{
		Name:    "GET NFT owner and state infomation",
		Method:  "Get",
		Path:    "/nftstate/info/{nftid}",
		Handler: apiHandler.GetWalletNFTStateInformation,
	},
}
