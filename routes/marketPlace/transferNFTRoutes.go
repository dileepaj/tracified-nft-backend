package marketPlace

import (
	"github.com/dileepaj/tracified-nft-backend/apiHandler"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var WalletNFTRoutes = models.Routers{
	models.Router{
		Name:    "Update wallet nft state",
		Method:  "PUT",
		Path:    "/walletnft/state/update",
		Handler: apiHandler.UpdateWalletNFTState,
	},
	models.Router{
		Name:    "Save wallet nft state txns",
		Method:  "POST",
		Path:    "/walletnft/txns/save",
		Handler: apiHandler.SaveWalletNFTTXNs,
	},
	models.Router{
		Name:    "Save Wallet NFT State",
		Method:  "POST",
		Path:    "/walletnft/state/save",
		Handler: apiHandler.SaveWalletNFTStates,
	},
	models.Router{
		Name:    "Get Wallet Users NFT By State",
		Method:  "Get",
		Path:    "/walletnft/{blockchain}/{nftstatus}/{currentowner}/{pagesize}/{requestedPage}",
		Handler: apiHandler.GetWalletNFTByStateAndCurrentOwner,
	},
	models.Router{
		Name:    "Get Wallet Users NFT Requestee By State",
		Method:  "Get",
		Path:    "/walletnft/requested/{blockchain}/{nftstatus}/{nftrequested}/{pagesize}/{requestedPage}",
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
}
