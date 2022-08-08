package marketplaceRepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
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
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session in getAllNFT : NFTRepository.go : ", err.Error())
	}
	defer session.EndSession(context.TODO())

	var nft []models.NFT
	findOptions := options.Find()
	findOptions.SetLimit(10)
	result, err := session.Client().Database(connections.DbName).Collection(NFT).Find(context.TODO(), bson.D{{}}, findOptions)
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
		logs.InfoLogger.Println("item info : ", nft)
		nfts = append(nfts, nft)
	}
	logs.InfoLogger.Println("Data retreived from DB : ", nfts)
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
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())

	rst, err := session.Client().Database(connections.DbName).Collection("svg").Find(context.TODO(), bson.M{"hash": hash})
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

func (r *NFTRepository) FindByFieldInMultipleValusWatchList(fields string, watchList []string) ([]models.WatchList, error) {
	var nfts []models.WatchList
	rst, err := repository.FindByFieldInMultipleValus(fields, watchList, WatchList)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return nfts, err
	}
	for rst.Next(context.TODO()) {
		var nft models.WatchList
		err = rst.Decode(&nft)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return nfts, err
		}
		nfts = append(nfts, nft)
	}
	return nfts, nil
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

func (r *NFTRepository) UpdateNFTSALE(findBy string, id string, update primitive.M) (models.NFT, error) {
	var nftResponse models.NFT
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := session.Client().Database(connections.DbName).Collection("nft").FindOneAndUpdate(context.TODO(), bson.M{"nftidentifier": id}, update, &opt)
	if rst != nil {
		err := rst.Decode((&nftResponse))
		logs.InfoLogger.Println("data retreived from DB: ", rst)
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

	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}

	defer session.EndSession(context.TODO())
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := session.Client().Database(connections.DbName).Collection("nft").FindOneAndUpdate(context.TODO(), bson.D{{"imagebase64", id1}, {findby2, id2}}, update, &opt)
	if rst != nil {
		err := rst.Decode((&nftResponse))
		logs.InfoLogger.Println("Minter update result: ", nftResponse)
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
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}

	defer session.EndSession(context.TODO())
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := session.Client().Database(connections.DbName).Collection("nft").FindOneAndUpdate(context.TODO(), bson.D{{"imagebase64", id}, {findbyid2, id2}}, update, &opt)
	if rst != nil {
		err := rst.Decode((&txnResponse))
		logs.InfoLogger.Println("DB Response for update Stellar: ", txnResponse)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from nft nft in UpdateNFTTXN:nftRepository.go: ", err.Error())
			return txnResponse, err
		}
		return txnResponse, nil
	} else {
		return txnResponse, nil

	}
}

func (r *NFTRepository) FindTagsByNFTIdentifier(idName string, id string) ([]models.Tags, error) {
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
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session in getAllNFT : NFTRepository.go : ", err.Error())
	}
	defer session.EndSession(context.TODO())

	var tag []models.Tags
	findOptions := options.Find()
	findOptions.SetLimit(10)
	result, err := session.Client().Database(connections.DbName).Collection(Tags).Find(context.TODO(), bson.D{{}}, findOptions)
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
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := session.Client().Database(connections.DbName).Collection("nft").FindOneAndUpdate(context.TODO(), bson.M{"nftidentifier": id}, update, &opt)
	if rst != nil {
		err := rst.Decode((&nftResponse))
		logs.InfoLogger.Println("data retreived from DB: ", rst)
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
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := session.Client().Database(connections.DbName).Collection("nft").FindOneAndUpdate(context.TODO(), bson.M{"nftidentifier": id}, update, &opt)
	if rst != nil {
		err := rst.Decode((&nftResponse))
		logs.InfoLogger.Println("data retreived from DB: ", rst)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from nft nft in UpdateNFTSALE:nftRepository.go: ", err.Error())
			return nftResponse, err
		}
		return nftResponse, nil
	} else {
		return nftResponse, nil

	}
}
