package marketplaceBusinessFacade

import (
	"fmt"

	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StoreNFT(createNFTObject models.NFT) (string, error) {
	rst, err1 := nftRepository.SaveNFT(createNFTObject)
	if err1 != nil {
		return "NFT not saved", err1
	}
	return rst, nil

}

func GetImageBase(nftcontent string) (models.NFT, error) {
	return nftRepository.FindImageBase(nftcontent)
}

func StoreNFTStory(createNFTObject models.NFTStory) (string, error) {
	rst, err1 := nftRepository.SaveNFTStory(createNFTObject)
	if err1 != nil {
		return "NFT not saved", err1
	}
	return rst, nil

}

func GetAllNFTs() ([]models.NFT, error) {
	return nftRepository.GetAllNFTs()
}

func StoreTXN(createTXNObject models.TXN) (string, error) {
	rst, err1 := nftRepository.SaveTXN(createTXNObject)
	if err1 != nil {
		return "TXNs not saved", err1
	}
	return rst, nil

}

func StoreOwner(createOwner models.Ownership) (string, error) {
	rst, err1 := nftRepository.SaveOwner(createOwner)
	if err1 != nil {
		return "Owner not saved", err1
	}
	return rst, nil
}

func GetAllONSaleNFT(id string, userPK string) ([]models.NFT, error) {
	return nftRepository.FindNFTById1AndNotId2("sellingstatus", id, "currentownerpk", userPK)
}

func GetNFTStory(id string, blockchain string) ([]models.NFTStory, error) {
	return nftRepository.FindNFTStory("nftidentifier", id, "blockchain", blockchain)
}

func GetNFTByCollection(paginationData requestDtos.NFTsForMatrixView, collectiontoSearch string, pubkey string) (models.Paginateresponse, error) {
	filter := bson.M{
		"blockchain":    paginationData.Blockchain,
		"collection":    collectiontoSearch,
		"creatoruserid": pubkey,
	}
	projection := GetProjectionDataNFTMatrixView()
	var nfts []models.NFTContentforMatrix
	response, err := nftRepository.GetNFTPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "nft", "_id", nfts)
	if err != nil {
		logs.ErrorLogger.Println("Error occurred :", err.Error())
		return models.Paginateresponse(response), err
	}
	return models.Paginateresponse(response), err
}

func GetProjectionDataNFTMatrixView() bson.D {
	projection := bson.D{
		{Key: "nftidentifier", Value: 1},
		{Key: "creatoruserid", Value: 1},
		{Key: "blockchain", Value: 1},
		{Key: "nftname", Value: 1},
		{Key: "imagebase64", Value: 1},
		{Key: "sellingstatus", Value: 1},
		{Key: "trending", Value: 1},
		{Key: "hotpicks", Value: 1},
		{Key: "currentownerpk", Value: 1},
		{Key: "attachmenttype", Value: 1},
		// {Key: "thumbnail", Value: 0},
	}
	return projection
}

func GetNFTPagination(paginationData requestDtos.NFTsForMatrixView) (models.Paginateresponse, error) {
	filter := bson.M{
		"blockchain": paginationData.Blockchain,
	}
	projection := GetProjectionDataNFTMatrixView()
	var nfts []models.NFTContentforMatrix
	response, err := nftRepository.GetNFTPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "nft", "_id", nfts)
	if err != nil {
		logs.ErrorLogger.Println("Error occurred :", err.Error())
		return models.Paginateresponse(response), err
	}
	return models.Paginateresponse(response), err
}

func GetPaginatedNFTbySellingStatus(paginationData requestDtos.NFTsForMatrixView) (models.Paginateresponse, error) {
	filter := bson.M{
		"blockchain":    paginationData.Blockchain,
		"sellingstatus": paginationData.SortbyFeild,
	}
	projection := GetProjectionDataNFTMatrixView()
	var nfts []models.NFTContentforMatrix
	response, err := nftRepository.GetNFTPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "nft", "_id", nfts)
	if err != nil {
		logs.ErrorLogger.Println("Error occurred :", err.Error())
		return models.Paginateresponse(response), err
	}
	return models.Paginateresponse(response), err
}

func GetPaginatedNFTbyStatusFilter(paginationData requestDtos.NFTsForMatrixView) (models.Paginateresponse, error) {
	var filter = bson.M{}
	if paginationData.SortbyFeild == "hotpicks" {
		filter = bson.M{
			"blockchain": paginationData.Blockchain,
			"hotpicks":   true,
		}
	} else if paginationData.SortbyFeild == "trending" {
		filter = bson.M{
			"blockchain": paginationData.Blockchain,
			"trending":   true,
		}
	}
	projection := GetProjectionDataNFTMatrixView()

	var nfts []models.NFTContentforMatrix
	response, err := nftRepository.GetNFTPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "nft", "_id", nfts)
	if err != nil {
		logs.ErrorLogger.Println("Error occured :", err.Error())
		return models.Paginateresponse(response), err
	}
	return models.Paginateresponse(response), err
}

func GetPaginatedOnSaleNFTbyStatusFilter(paginationData requestDtos.NFTsForMatrixView) (models.Paginateresponse, error) {
	var filter = bson.M{}
	if paginationData.SortbyFeild == "hotpicks" {
		filter = bson.M{
			"sellingstatus": "ON SALE",
			"hotpicks":      true,
		}
	} else if paginationData.SortbyFeild == "trending" {
		filter = bson.M{
			"sellingstatus": "ON SALE",
			"trending":      true,
		}
	}
	projection := GetProjectionDataNFTMatrixView()

	var nfts []models.NFTContentforMatrix
	response, err := nftRepository.GetNFTPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "nft", "_id", nfts)
	if err != nil {
		logs.ErrorLogger.Println("Error occured :", err.Error())
		return models.Paginateresponse(response), err
	}
	return models.Paginateresponse(response), err
}

func GetPaginatedResponseforBestCreations(paginationData requestDtos.NFTsForMatrixView) (models.Paginateresponse, error) {
	filter := bson.M{
		"blockchain": paginationData.Blockchain,
		"hotpicks":   true,
		"trending":   true,
	}
	var nfts []models.NFTContentforMatrix
	projection := GetProjectionDataNFTMatrixView()
	response, err := nftRepository.GetNFTPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "nft", "_id", nfts)
	if err != nil {
		logs.ErrorLogger.Println("Error occured :", err.Error())
		return models.Paginateresponse(response), err
	}
	return models.Paginateresponse(response), err
}

func GetOneONSaleNFT(id string, identifier string, blockchain string) ([]models.NFT, error) {
	return nftRepository.FindNFTByIdId2Id3("sellingstatus", id, "nftidentifier", identifier, "blockchain", blockchain)
}

func GetNFTByBlockchainAndUserPK(id string, blockchain string) ([]models.NFT, error) {
	return nftRepository.FindNFTById1AndNotId2("currentownerpk", id, "blockchain", blockchain)
}

func GetTXNByBlockchainAndIdentifier(id string, blockchain string) ([]models.TXN, error) {
	return nftRepository.FindTXNById1AndNotId2("nftidentifier", id, "blockchain", blockchain)
}

func MakeSaleNFT(nft requestDtos.UpdateNFTSALERequest) (models.NFT, error) {
	update := bson.M{
		"$set": bson.M{"timestamp": nft.Timestamp, "currentprice": nft.CurrentPrice, "sellingstatus": nft.SellingStatus, "sellingtype": nft.SellingType, "marketcontract": nft.MarketContract, "currentownerpk": nft.CurrentOwnerPK, "royalty": nft.Royalty, "commission": nft.Commission},
	}
	return nftRepository.UpdateNFTSALE("nftidentifier", nft.NFTIdentifier, "blockchain", nft.Blockchain, update)
}

func GetBlockchainSpecificNFT(blockchain string) ([]models.NFT, error) {
	return nftRepository.FindNFTsById("blockchain", blockchain)
}

func GetNFTBySellingStatus(status string) ([]models.NFT, error) {
	return nftRepository.FindNFTsById("sellingstatus", status)
}

func GEtNFTbyTagsName(paginationData requestDtos.NFTsForMatrixView, tagToSearch string) (models.Paginateresponse, error) {
	regexPattern := primitive.Regex{Pattern: tagToSearch, Options: "i"}
	filter := bson.M{
		"sellingstatus": "ON SALE",
		"tags":          regexPattern,
	}
	projection := GetProjectionDataNFTMatrixView()
	var nfts []models.NFTContentforMatrix
	response, err := nftRepository.GetNFTPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "nft", "_id", nfts)
	if err != nil {
		logs.ErrorLogger.Println("Error occurred :", err.Error())
		return models.Paginateresponse(response), err
	}
	return models.Paginateresponse(response), err
}

func GetNFTbyAccount(userId string) ([]models.NFT, error) {
	return nftRepository.FindNFTsById("currentownerpk", userId)
}

func GetLastNFTbyUserId(userId string) ([]models.NFT, error) {
	return nftRepository.FindLastNFTById("creatoruserid", userId)
}

func GetSVGByHash(hash string) (models.SVG, error) {
	return nftRepository.GetSVGByHash(hash)
}

func GetNFTbyTenentName(tenentName string) ([]models.NFT, error) {
	return nftRepository.FindNFTsById("creatoruserid", tenentName)
}

func GetNFTbyBlockchain(blockchain string) ([]models.NFT, error) {
	return nftRepository.FindNFTByBlockchain("blockchain", blockchain)
}

func CreateTags(tags models.Tags) (string, error) {
	return nftRepository.SaveTags(tags)
}

func GetAllTags() ([]models.Tags, error) {
	return nftRepository.GetAllTags()
}
func GetTagsByNFTName(nftid string) ([]models.Tags, error) {
	return nftRepository.FindTagsByNFTName("nftName", nftid)

}

func UpdateNFTTXN(txn requestDtos.UpdateMintTXN) (models.NFT, error) {
	update := bson.M{
		"$set": bson.M{"nfttxnhash": txn.NFTTxnHash},
	}
	return nftRepository.UpdateNFTTXN("imagebase64", txn.Imagebase64, "blockchain", txn.Blockchian, update)
}
func UpdateNFT(nft requestDtos.UpdateMint) (models.NFT, error) {
	update := bson.M{
		"$set": bson.M{"nftidentifier": nft.NFTIdentifier, "nftissuerpk": nft.NFTIssuerPK, "nfttxnhash": nft.NFTTxnHash},
	}
	return nftRepository.UpdateMinter("imagebase64", nft.Imagebase64, "blockchain", nft.Blockchain, update)
}

func GetNFTByNFTIdentifier(id string) ([]models.NFT, error) {
	return nftRepository.FindNFTsById("nftidentifier", id)
}

func GetThumbnailbyID(id string) (models.ThumbNail, error) {
	return nftRepository.GetThumbnailbyID(id)
}
func GetUserProfileContent(paginationData requestDtos.NFTsForMatrixView, filterquest string) (models.Paginateresponse, error) {
	var filter bson.M
	var nfts []models.NFTContentforMatrix
	var emptyNft models.Paginateresponse
	if paginationData.Blockchain != "ethereum" &&
		paginationData.Blockchain != "polygon" &&
		paginationData.Blockchain != "stellar" &&
		paginationData.Blockchain != "solana" {
		err := fmt.Errorf("invalid blockchain: %s", paginationData.Blockchain)
		return emptyNft, err
	}
	if filterquest == "minted" {
		filter = bson.M{
			"blockchain":     paginationData.Blockchain,
			"currentownerpk": paginationData.SortbyFeild,
			"sellingstatus":  "Minted",
		}
	} else if filterquest == "hotpicks" {
		filter = bson.M{
			"blockchain":     paginationData.Blockchain,
			"currentownerpk": paginationData.SortbyFeild,
			"hotpicks":       true,
		}
	} else if filterquest == "favorite" {
		filter = bson.M{
			"blockchain":     paginationData.Blockchain,
			"currentownerpk": paginationData.SortbyFeild,
			"trending":       true,
		}
	} else if filterquest == "bought" {
		filter = bson.M{
			"blockchain":     paginationData.Blockchain,
			"currentownerpk": paginationData.SortbyFeild,
			"sellingstatus":  "NOTFORSALE",
		}
	} else if filterquest == "onsale" {
		filter = bson.M{
			"blockchain":     paginationData.Blockchain,
			"currentownerpk": paginationData.SortbyFeild,
			"sellingstatus":  "ON SALE",
		}
	} else {
		err := fmt.Errorf("invalid Filter: %s ", filterquest)
		return emptyNft, err
	}
	projection := GetProjectionDataNFTMatrixView()

	response, err := nftRepository.GetNFTPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "nft", "_id", nfts)
	if err != nil {
		logs.ErrorLogger.Println("Error occurred :", err.Error())
		return models.Paginateresponse(response), err
	}
	return models.Paginateresponse(response), err
}

func StoreContracts(createContractObject models.ContractInfo) (string, error) {
	rst, err1 := nftRepository.SaveContract(createContractObject)
	if err1 != nil {
		return "Contract not saved", err1
	}
	return rst, nil

}

func GetContractbyBlockchainAndUser(blockchain string, user string) ([]models.ContractInfo, error) {
	return nftRepository.FindContractByBCandUser("blockchain", blockchain, "user", user)
}

func StoreWalletNFT(createNFTObject models.WalletNFT) (string, error) {
	rst, err1 := nftRepository.SaveWalletNFT(createNFTObject)
	if err1 != nil {
		return "NFT not saved", err1
	}
	return rst, nil

}

func GetAllWalletNFTs() ([]models.ResponseWalletNFT, error) {
	return nftRepository.GetAllWalletNFTs()
}

func GetAllWalletNFTsbyPK(publickey string) ([]models.ResponseWalletNFT, error) {
	return nftRepository.GetWalletNFTsbyPK(publickey)
}

func GetNFTByBlockchainAndIdentifier(id string, blockchain string) (models.NFT, error) {
	return nftRepository.GetNFTByIDAndBC("nftidentifier", id, "blockchain", blockchain)
}

func StoreNFTState(createNFTStateObject models.NFTWalletState) (string, error) {
	rst, err1 := nftRepository.SaveNFTState(createNFTStateObject)
	if err1 != nil {
		return "NFT not saved", err1
	}
	return rst, nil

}

func StoreNFTStateTXN(createNFTStateTxnObject models.NFTWalletStateTXN) (string, error) {
	rst, err1 := nftRepository.SaveNFTStateTXN(createNFTStateTxnObject)
	if err1 != nil {
		return "NFT not saved", err1
	}
	return rst, nil

}

func UpdateNFTState(state requestDtos.UpdateNFTState) (models.NFTWalletState, error) {
	update := bson.M{
		"$set": bson.M{"nftstatus": state.NFTStatus},
	}
	return nftRepository.UpdateNFTState("issuerpublickey", state.IssuerPublicKey, update)
}

func GetCurrentNFTState(issuerpublickey string) (uint8, error) {
	return nftRepository.GetCurrentNFTStatus(issuerpublickey)
}

func DeleteNFTStateByIssuer(nftstate requestDtos.DeleteNFTState) (int64, error) {
	return nftRepository.DeleteNFTState(nftstate)
}

func GetWalletNFTTxnsByIssuer(issuer string) ([]models.NFTWalletStateTXN, error) {
	return nftRepository.FindWalletNFTTxns("issuerpublickey", issuer)
}

// GetWalletNFTByState retrieves wallet NFTs based on the provided state, blockchain, and current owner public key.
// It takes paginationData, which contains information about the blockchain, requested page, and page size.
// The StatetoSearch parameter represents the state of the NFTs to filter (0 for all states).
// The pubkey parameter is the current owner's public key used for filtering.
// It returns a PaginateWalletNFTResponse containing the paginated NFTs and an error, if any.
func GetWalletNFTByState(paginationData requestDtos.WalletNFTsForMatrixView, StatetoSearch int, pubkey string) (models.PaginateWalletNFTResponse, error) {
	// Define the filter based on the provided state and current owner's public key.
	var filter bson.M
	if StatetoSearch == 0 {
		filter = bson.M{
			"blockchain":   paginationData.Blockchain,
			"currentowner": pubkey,
		}
	} else {
		filter = bson.M{
			"blockchain":   paginationData.Blockchain,
			"nftstatus":    StatetoSearch,
			"currentowner": pubkey,
		}
	}

	// Define the projection to retrieve the desired fields for NFTs.
	projection := getProjectionDataWalletNFTMatrixView()

	// Initialize a slice to store the retrieved NFTs.
	var nfts []models.WalletNFTContentforMatrix

	// Call the nftRepository to get the paginated response based on the provided filter and projection.
	response, err := nftRepository.GetWalletNFTPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "nftstate", "_id", nfts)
	if err != nil {
		// Log the error and return the response along with the error.
		logs.ErrorLogger.Println("Error occurred:", err.Error())
		return models.PaginateWalletNFTResponse(response), err
	}

	// Return the paginated response and any potential error.
	return models.PaginateWalletNFTResponse(response), err
}

// getProjectionDataWalletNFTMatrixView returns the projection document to specify which fields to include in the query result.
// This private function defines the fields that will be selected when querying wallet NFTs for the matrix view.
// The returned projection is a bson.D document containing field specifications.
func getProjectionDataWalletNFTMatrixView() bson.D {
	// Define the projection document to include specific fields in the query result.
	projection := bson.D{
		{Key: "issuerpublickey", Value: 1},
		{Key: "nftcreator", Value: 1},
		{Key: "blockchain", Value: 1},
		{Key: "nftname", Value: 1},
		{Key: "nftrequested", Value: 1},
		{Key: "currentowner", Value: 1},
		{Key: "nftstatus", Value: 1},
		{Key: "timestamp", Value: 1},
		{Key: "shopid", Value: 1},
		{Key: "thumbnail", Value: 1},
		{Key: "nftid", Value: 1},
	}

	// Return the projection document.
	return projection
}

// GetWalletNFTByStateForRequested retrieves wallet NFTs based on the provided state, blockchain, and requested public key.
// It takes paginationData, which contains information about the blockchain, requested page, and page size.
// The StatetoSearch parameter represents the state of the NFTs to filter (0 for all states).
// The pubkey parameter is the requested public key used for filtering.
// It returns a PaginateWalletNFTResponse containing the paginated NFTs and an error, if any.
func GetWalletNFTByStateForRequested(paginationData requestDtos.WalletNFTsForMatrixView, StatetoSearch int, pubkey string) (models.PaginateWalletNFTResponse, error) {
	// Define the filter based on the provided state and requested public key.
	var filter bson.M
	if StatetoSearch == 0 {
		filter = bson.M{
			"blockchain":   paginationData.Blockchain,
			"nftrequested": pubkey,
		}
	} else {
		filter = bson.M{
			"blockchain":   paginationData.Blockchain,
			"nftstatus":    StatetoSearch,
			"nftrequested": pubkey,
		}
	}

	// Define the projection to retrieve the desired fields for NFTs.
	projection := getProjectionDataWalletNFTMatrixView()

	// Initialize a slice to store the retrieved NFTs.
	var nfts []models.WalletNFTContentforMatrix

	// Call the nftRepository to get the paginated response based on the provided filter and projection.
	response, err := nftRepository.GetWalletNFTPaginatedResponse(filter, projection, paginationData.PageSize, paginationData.RequestedPage, "nftstate", "_id", nfts)
	if err != nil {
		// Log the error and return the response along with the error.
		logs.ErrorLogger.Println("Error occurred:", err.Error())
		return models.PaginateWalletNFTResponse(response), err
	}

	// Return the paginated response and any potential error.
	return models.PaginateWalletNFTResponse(response), err
}

func UpdateWalletNFTOwner(nft requestDtos.WalletNFTUpdateOwner) (requestDtos.WalletNFTUpdateOwner, error) {
	update := bson.M{
		"$set": bson.M{"nftowner": nft.NFTOwner},
	}
	return nftRepository.UpdateWalletNFTOwner("_id", nft.ID, update)
}
