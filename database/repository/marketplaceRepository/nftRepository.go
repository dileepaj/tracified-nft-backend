package marketplaceRepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NFTRepository struct{}

var NFT = "nft"
var Tags = "tags"
var Owner = "owner"
var Story = "nftstory"
var Contract = "contracts"
var walletnft = "walletnft"
var NFTState = "nftstate"
var walletNFTTxns = "nftstatetxns"

func (r *NFTRepository) FindNFTById1AndNotId2(idName1 string, id1 string, idName2 string, id2 string) ([]models.NFT, error) {
	var nfts []models.NFT
	rst, err := repository.FindById1AndNotId2(idName1, id1, idName2, id2, NFT)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nfts, err
	}
	for rst.Next(context.TODO()) {
		var nft models.NFT
		err = rst.Decode(&nft)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nfts, err
		}
		nfts = append(nfts, nft)
	}
	return nfts, nil
}

func (r *NFTRepository) FindImageBase(idName1 string) (models.NFT, error) {
	var nft models.NFT
	rst, err := connections.GetSessionClient("nft").Find(context.TODO(), bson.M{"imagebase64": idName1})
	if err != nil {
		return nft, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&nft)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection nft in GetImageBase64:NFTRepository.go: ", err.Error())
			return nft, err
		}
	}
	return nft, err
}

func (r *NFTRepository) FindNFTStory(idName1 string, id1 string, idName2 string, id2 string) ([]models.NFTStory, error) {
	var nfts []models.NFTStory
	rst, err := repository.FindById1AndNotId2(idName1, id1, idName2, id2, Story)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nfts, err
	}
	for rst.Next(context.TODO()) {
		var nft models.NFTStory
		err = rst.Decode(&nft)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nfts, err
		}
		nfts = append(nfts, nft)
	}
	return nfts, nil
}

func (r *NFTRepository) FindNFTByCollection(idName1 string, id1 string) ([]models.NFT, error) {
	var nfts []models.NFT
	rst, err := repository.FindById(idName1, id1, NFT)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nfts, err
	}
	for rst.Next(context.TODO()) {
		var nft models.NFT
		err = rst.Decode(&nft)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nfts, err
		}
		nfts = append(nfts, nft)
	}
	return nfts, nil
}

func (r *NFTRepository) FindTXNById1AndNotId2(idName1 string, id1 string, idName2 string, id2 string) ([]models.TXN, error) {
	var nfts []models.TXN
	rst, err := repository.FindById1AndNotId2(idName1, id1, idName2, id2, Txn)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nfts, err
	}
	for rst.Next(context.TODO()) {
		var nft models.TXN
		err = rst.Decode(&nft)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nfts, err
		}
		nfts = append(nfts, nft)
	}
	return nfts, nil
}

func (r *NFTRepository) GetAllNFTs() ([]models.NFT, error) {
	var nft []models.NFT
	findOptions := options.Find()

	result, err := connections.GetSessionClient(NFT).Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println("Error occured when trying to connect to DB and excute Find query in GetAllNFT:NFTRepository.go: ", err.Error())
		return nft, err
	}
	for result.Next(context.TODO()) {
		var nfts models.NFT
		err = result.Decode(&nfts)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection nfts in GetAllNFTs:nftsRepository.go: ", err.Error())
			return nft, err
		}
		nft = append(nft, nfts)
	}
	return nft, nil
}

func (r *NFTRepository) FindNFTByIdId2Id3(idName1 string, id1 string, idName2 string, id2 string, idName3 string, id3 string) ([]models.NFT, error) {
	var nfts []models.NFT
	rst, err := repository.FindById1Id2Id3(idName1, id1, idName2, id2, idName3, id3, NFT)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nfts, err
	}
	for rst.Next(context.TODO()) {
		var nft models.NFT
		err = rst.Decode(&nft)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nfts, err
		}
		nfts = append(nfts, nft)
	}
	return nfts, nil
}

func (r *NFTRepository) FindNFTsById(idName string, id string) ([]models.NFT, error) {
	var nfts []models.NFT
	rst, err := repository.FindById(idName, id, NFT)
	if err != nil {
		return nfts, err
	}
	for rst.Next(context.TODO()) {
		var nft models.NFT
		err = rst.Decode(&nft)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nfts, err
		}
		nfts = append(nfts, nft)
	}
	return nfts, nil
}

func (r *NFTRepository) GetSVGByHash(hash string) (models.SVG, error) {
	var svg models.SVG
	rst, err := connections.GetSessionClient("svg").Find(context.TODO(), bson.M{"hash": hash})
	if err != nil {
		return svg, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&svg)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection svg in GetSVGByHash:NFTRepository.go: ", err.Error())
			return svg, err
		}
	}
	return svg, err
}

func (r *NFTRepository) GetThumbnailbyID(id string) (models.ThumbNail, error) {
	var thumbnail models.ThumbNail
	objectID, objerr := primitive.ObjectIDFromHex(id)
	if objerr != nil {
		logs.WarningLogger.Println("Error Occured when trying to convert hex string in to Object(ID) in GetThumbnailbyID : nftRepo: ", objerr.Error())
		return thumbnail, objerr
	}

	rst, err := connections.GetSessionClient("nft").Find(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return thumbnail, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&thumbnail)
		if err != nil {
			logs.ErrorLogger.Println("Failed to retrive NFT thumbnail: ", err.Error())
			return thumbnail, err
		}
	}
	return thumbnail, err
}

func (r *NFTRepository) FindLastNFTById(idName string, id string) ([]models.NFT, error) {
	var nfts []models.NFT
	rst := repository.FindOne(idName, id, "nft")
	if rst != nil {
		err := rst.Decode(&nfts)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nfts, err
		}
		return nfts, nil
	} else {
		return nfts, nil
	}
}

func (r *NFTRepository) FindByFieldInMultipleValusTags(fields string, tags []string) ([]models.NFT, error) {
	var nfts []models.NFT
	rst, err := repository.FindByFieldInMultipleValus(fields, tags, Tags)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nfts, err
	}
	for rst.Next(context.TODO()) {
		var nft models.NFT
		err = rst.Decode(&nft)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nfts, err
		}
		nfts = append(nfts, nft)
	}
	return nfts, nil
}

func (r *NFTRepository) FindByFieldInMultipleValusAccount(fields string, nft []string) ([]models.NFT, error) {
	var nfts []models.NFT
	rst, err := repository.FindByFieldInMultipleValus(fields, nft, NFT)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nfts, err
	}
	for rst.Next(context.TODO()) {
		var nft models.NFT
		err = rst.Decode(&nft)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nfts, err
		}
		nfts = append(nfts, nft)
	}
	return nfts, nil
}

func (r *NFTRepository) FindByFieldInMultipleValusTennant(fields string, owner []string) ([]models.NFT, error) {
	var nfts []models.NFT
	rst, err := repository.FindByFieldInMultipleValus(fields, owner, Owner)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nfts, err
	}
	for rst.Next(context.TODO()) {
		var nft models.NFT
		err = rst.Decode(&nft)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nfts, err
		}
		nfts = append(nfts, nft)
	}
	return nfts, nil
}

func (r *NFTRepository) SaveNFT(nft models.NFT) (string, error) {
	return repository.Save[models.NFT](nft, NFT)
}

func (r *NFTRepository) SaveNFTStory(nft models.NFTStory) (string, error) {
	return repository.Save[models.NFTStory](nft, Story)
}

func (r *NFTRepository) SaveTXN(txn models.TXN) (string, error) {
	return repository.Save[models.TXN](txn, Txn)
}

func (r *NFTRepository) SaveOwner(owner models.Ownership) (string, error) {
	return repository.Save[models.Ownership](owner, Owner)
}

func (r *NFTRepository) UpdateNFTSALE(findBy string, id string, findby2 string, id2 string, update primitive.M) (models.NFT, error) {
	var nftResponse models.NFT
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("nft").FindOneAndUpdate(context.TODO(), bson.D{{findBy, id}, {findby2, id2}}, update, &opt)
	if rst != nil {
		err := rst.Decode((&nftResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from nft nft in UpdateNFTSALE:nftRepository.go: ", err.Error())
			return nftResponse, err
		}
		return nftResponse, nil
	} else {
		return nftResponse, nil

	}
}

func (r *NFTRepository) UpdateMinter(findBy1 string, id1 string, findby2 string, id2 string, update primitive.M) (models.NFT, error) {
	var nftResponse models.NFT
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("nft").FindOneAndUpdate(context.TODO(), bson.D{{"imagebase64", id1}, {findby2, id2}}, update, &opt)
	if rst != nil {
		err := rst.Decode((&nftResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection nft in UpdateMinter:nftRepository.go: ", err.Error())
			return nftResponse, err
		}
		return nftResponse, nil
	} else {
		return nftResponse, nil

	}
}

func (r *NFTRepository) UpdateNFTTXN(findBy string, id string, findbyid2 string, id2 string, update primitive.M) (models.NFT, error) {
	var txnResponse models.NFT
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("nft").FindOneAndUpdate(context.TODO(), bson.D{{"imagebase64", id}, {findbyid2, id2}}, update, &opt)
	if rst != nil {
		err := rst.Decode((&txnResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from nft nft in UpdateNFTTXN:nftRepository.go: ", err.Error())
			return txnResponse, err
		}
		return txnResponse, nil
	} else {
		return txnResponse, nil

	}
}

func (r *NFTRepository) FindTagsByNFTName(idName string, id string) ([]models.Tags, error) {
	var tags []models.Tags
	rst, err := repository.FindById(idName, id, Tags)
	if err != nil {
		return tags, err
	}
	for rst.Next(context.TODO()) {
		var tag models.Tags
		err = rst.Decode(&tag)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return tags, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (r *NFTRepository) FindNFTByBlockchain(idName string, id string) ([]models.NFT, error) {
	var nfts []models.NFT
	rst, err := repository.FindById(idName, id, NFT)
	if err != nil {
		return nfts, err
	}
	for rst.Next(context.TODO()) {
		var nft models.NFT
		err = rst.Decode(&nft)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nfts, err
		}
		nfts = append(nfts, nft)
	}
	return nfts, nil
}

func (r *NFTRepository) GetAllTags() ([]models.Tags, error) {
	var tag []models.Tags
	findOptions := options.Find()
	result, err := connections.GetSessionClient(Tags).Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println("Error occured when trying to connect to DB and excute Find query in GetAllNFT:NFTRepository.go: ", err.Error())
		return tag, err
	}
	for result.Next(context.TODO()) {
		var tags models.Tags
		err = result.Decode(&tags)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection tags in GetAllTags:tagsRepository.go: ", err.Error())
			return tag, err
		}
		tag = append(tag, tags)
	}
	return tag, nil
}

func (r *NFTRepository) SaveTags(tags models.Tags) (string, error) {
	return repository.Save[models.Tags](tags, Tags)
}

func (r *NFTRepository) UpdateTrending(findBy string, id string, update primitive.M) (models.NFT, error) {
	var nftResponse models.NFT
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("nft").FindOneAndUpdate(context.TODO(), bson.M{"nftidentifier": id}, update, &opt)
	if rst != nil {
		err := rst.Decode((&nftResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from nft nft in UpdateNFTSALE:nftRepository.go: ", err.Error())
			return nftResponse, err
		}
		return nftResponse, nil
	} else {
		return nftResponse, nil

	}
}

func (r *NFTRepository) UpdateHotPicks(findBy string, id string, update primitive.M) (models.NFT, error) {
	var nftResponse models.NFT
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("nft").FindOneAndUpdate(context.TODO(), bson.M{"nftidentifier": id}, update, &opt)
	if rst != nil {
		err := rst.Decode((&nftResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from nft nft in UpdateNFTSALE:nftRepository.go: ", err.Error())
			return nftResponse, err
		}
		return nftResponse, nil
	} else {
		return nftResponse, nil

	}
}

func (r *NFTRepository) GetNFTPaginatedResponse(filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, nfts []models.NFTContentforMatrix) (models.Paginateresponse, error) {
	contentResponse, paginationResponse, err := repository.PaginateResponse[[]models.NFTContentforMatrix](
		filterConfig,
		projectionData,
		pagesize,
		pageNo,
		collectionName,
		sortingFeildName,
		nfts,
	)
	var response models.Paginateresponse
	if err != nil {
		return response, err
	}
	response.Content = contentResponse
	response.PaginationInfo = paginationResponse
	return response, nil
}

func (r *NFTRepository) SaveContract(contract models.ContractInfo) (string, error) {
	return repository.Save[models.ContractInfo](contract, Contract)
}

func (r *NFTRepository) FindContractByBCandUser(idName1 string, id1 string, idName2 string, id2 string) ([]models.ContractInfo, error) {
	var contract []models.ContractInfo
	rst, err := repository.FindById1AndNotId2(idName1, id1, idName2, id2, Contract)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return contract, err
	}
	for rst.Next(context.TODO()) {
		var contracts models.ContractInfo
		err = rst.Decode(&contracts)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return contract, err
		}
		contract = append(contract, contracts)
	}
	return contract, nil
}

func (r *NFTRepository) SaveWalletNFT(nft models.WalletNFT) (string, error) {
	return repository.Save[models.WalletNFT](nft, walletnft)
}

func (r *NFTRepository) GetAllWalletNFTs() ([]models.ResponseWalletNFT, error) {
	var nft []models.ResponseWalletNFT
	findOptions := options.Find().SetProjection(bson.M{"otp": 0})

	result, err := connections.GetSessionClient(walletnft).Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println("Error occured when trying to connect to DB and excute Find query in GetAllWalletNFTs:NFTRepository.go: ", err.Error())
		return nft, err
	}
	for result.Next(context.TODO()) {
		var nfts models.ResponseWalletNFT
		err = result.Decode(&nfts)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection nfts in GetAllWalletNFTs:nftsRepository.go: ", err.Error())
			return nft, err
		}
		nft = append(nft, nfts)
	}
	return nft, nil
}

func (r *NFTRepository) GetWalletNFTsbyPK(publickey string) ([]models.ResponseWalletNFT, error) {
	var nft []models.ResponseWalletNFT
	findOptions := options.Find().SetProjection(bson.M{"otp": 0})

	result, err := connections.GetSessionClient(walletnft).Find(context.TODO(), bson.D{{"nftowner", publickey}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println("Error occured when trying to connect to DB and excute Find query in GetWalletNFTsbyPK:NFTRepository.go: ", err.Error())
		return nft, err
	}
	for result.Next(context.TODO()) {
		var nfts models.ResponseWalletNFT
		err = result.Decode(&nfts)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection nfts in GetWalletNFTsbyPK:nftsRepository.go: ", err.Error())
			return nft, err
		}
		nft = append(nft, nfts)
	}
	return nft, nil
}

func (r *NFTRepository) GetNFTByIDAndBC(idName1 string, id1 string, idName2 string, id2 string) (models.NFT, error) {
	var nft models.NFT
	rst, err := repository.FindById1AndNotId2(idName1, id1, idName2, id2, NFT)
	if err != nil {
		return nft, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&nft)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection nft in GetNFTByIDAndBC:NFTRepository.go: ", err.Error())
			return nft, err
		}
	}
	return nft, err
}

func (r *NFTRepository) SaveNFTState(nftstate models.NFTWalletState) (string, error) {
	return repository.Save[models.NFTWalletState](nftstate, NFTState)
}

func (r *NFTRepository) SaveNFTStateTXN(nftstate models.NFTWalletStateTXN) (string, error) {
	return repository.Save[models.NFTWalletStateTXN](nftstate, walletNFTTxns)
}

func (r *NFTRepository) UpdateNFTState(findBy string, id string, update primitive.M) (models.NFTWalletState, error) {
	var stateResponse models.NFTWalletState
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("nftstate").FindOneAndUpdate(context.TODO(), bson.D{{"issuerpublickey", id}}, update, &opt)
	if rst != nil {
		err := rst.Decode((&stateResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from nft nft in UpdateNFTTXN:nftRepository.go: ", err.Error())
			return stateResponse, err
		}
		return stateResponse, nil
	} else {
		return stateResponse, nil

	}
}

func (r *NFTRepository) GetCurrentNFTStatus(issuerpublickey string) (uint8, error) {
	var nftState models.NFTWalletState
	nftState.NFTStatus = 0
	rst := repository.FindOne("issuerpublickey", issuerpublickey, "nftstate")
	if rst != nil {
		err := rst.Decode(&nftState)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nftState.NFTStatus, err
		}
	}
	return nftState.NFTStatus, nil
}

func (r *NFTRepository) DeleteNFTState(nftstate requestDtos.DeleteNFTState) (int64, error) {
	result, err := connections.GetSessionClient(NFTState).DeleteOne(context.TODO(), bson.M{"issuerpublickey": nftstate.IssuerPublicKey})
	if err != nil {
		logs.ErrorLogger.Println("Error occured when Connecting to DB and executing DeleteOne Query in DeleteNFTState(NFTRepository): ", err.Error())
	}
	return result.DeletedCount, err

}

func (r *NFTRepository) FindWalletNFTTxns(idName1 string, id1 string) ([]models.NFTWalletStateTXN, error) {
	var txns []models.NFTWalletStateTXN
	rst, err := repository.FindById(idName1, id1, walletNFTTxns)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return txns, err
	}
	for rst.Next(context.TODO()) {
		var txn models.NFTWalletStateTXN
		err = rst.Decode(&txn)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return txns, err
		}
		txns = append(txns, txn)
	}
	return txns, nil
}

func (r *NFTRepository) GetWalletNFTPaginatedResponse(filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, nfts []models.WalletNFTContentforMatrix) (models.PaginateWalletNFTResponse, error) {
	contentResponse, paginationResponse, err := repository.PaginateResponse[[]models.WalletNFTContentforMatrix](
		filterConfig,
		projectionData,
		pagesize,
		pageNo,
		collectionName,
		sortingFeildName,
		nfts,
	)
	var response models.PaginateWalletNFTResponse
	if err != nil {
		return response, err
	}
	response.Content = contentResponse
	response.PaginationInfo = paginationResponse
	return response, nil
}

func (r *NFTRepository) UpdateWalletNFTOwner(findBy1 string, id1 primitive.ObjectID, update primitive.M) (requestDtos.WalletNFTUpdateOwner, error) {
	var nftResponse requestDtos.WalletNFTUpdateOwner
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("walletnft").FindOneAndUpdate(context.TODO(), bson.D{{Key: findBy1, Value: id1}}, update, &opt)
	if rst != nil {
		err := rst.Decode((&nftResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from nft nft in UpdateNFTSALE:nftRepository.go: ", err.Error())
			return nftResponse, err
		}
		return nftResponse, nil
	} else {
		return nftResponse, nil

	}
}

func (r *NFTRepository) GetWalletNFTStateInformation(id string) (responseDtos.WalletNFTStateInfo, error) {
	var nftInfo responseDtos.WalletNFTStateInfo
	rst := repository.FindOne("nftid", id, "nftstate")
	if rst != nil {
		err := rst.Decode(&nftInfo)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nftInfo, err
		}
	}
	return nftInfo, nil
}
