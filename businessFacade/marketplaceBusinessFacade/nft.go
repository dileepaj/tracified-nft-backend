package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func StoreNFT(createNFTObject models.NFT) (string, error) {
	rst, err1 := nftRepository.SaveNFT(createNFTObject)
	if err1 != nil {
		return "NFT not saved", err1
	}
	return rst, nil

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
		"$set": bson.M{"timestamp": nft.Timestamp, "currentprice": nft.CurrentPrice, "sellingstatus": nft.SellingStatus, "sellingtype": nft.SellingType, "marketcontract": nft.MarketContract, "currentownerpk": nft.CurrentOwnerPK},
	}
	return nftRepository.UpdateNFTSALE("nftidentifier", nft.NFTIdentifier, update)
}

func GetBlockchainSpecificNFT(blockchain string) ([]models.NFT, error) {
	return nftRepository.FindNFTsById("blockchain", blockchain)
}

func GetNFTBySellingStatus(status string) ([]models.NFT, error) {
	return nftRepository.FindNFTsById("sellingstatus", status)
}

func GetNFTbyTagsName(tags string) ([]models.NFT, error) {
	return nftRepository.FindNFTsById("tags", tags)
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
