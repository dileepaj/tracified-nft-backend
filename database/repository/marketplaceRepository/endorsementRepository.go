package marketplaceRepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EndorsementRepository struct{}

var Endorsement = "endorsement"

func (r *EndorsementRepository) UpdateEndorsement(findBy string, id string, findBy2 string, id2 string, update primitive.M) (responseDtos.ResponseEndorsementUpdate, error) {
	var endorseResponse responseDtos.ResponseEndorsementUpdate
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("endorsement").FindOneAndUpdate(context.TODO(), bson.D{{findBy, id}, {findBy2, id2}}, update, &opt)
	if rst != nil {
		err := rst.Decode((&endorseResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection endorsement in UpdateEndorsement:EndorsementRepository.go: ", err.Error())
			return endorseResponse, err
		}
		return endorseResponse, nil
	} else {
		return endorseResponse, nil

	}
}

func (r *EndorsementRepository) SaveEndorsement(endorse models.Endorse) (string, error) {
	endorse.IsBestCreator = false
	return repository.Save[models.Endorse](endorse, Endorsement)
}

func (r *EndorsementRepository) FindEndorsermentbyPK(publickey string) (models.Endorse, error) {
	var endorse models.Endorse
	rst, err := connections.GetSessionClient("endorsement").Find(context.TODO(), bson.M{"publickey": publickey})
	if err != nil {
		return endorse, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&endorse)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection endorse in GetEndorseByID:EndorsementRepository.go: ", err.Error())
			return endorse, err
		}
	}
	return endorse, err
}

func (r *EndorsementRepository) GetEndorsementByStatus(idName string, id string) ([]models.Endorse, error) {
	var endorses []models.Endorse
	rst, err := repository.FindById(idName, id, Endorsement)
	if err != nil {
		return endorses, err
	}
	for rst.Next(context.TODO()) {
		var endorse models.Endorse
		err = rst.Decode(&endorse)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return endorses, err
		}
		endorses = append(endorses, endorse)
	}
	return endorses, nil
}

func (r *EndorsementRepository) UpdateSetEndorsement(findBy string, id string, update primitive.M) (models.Endorse, error) {
	var endorseResponse models.Endorse
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("endorsement").FindOneAndUpdate(context.TODO(), bson.M{"publickey": id}, update, &opt)
	if rst != nil {
		err := rst.Decode((&endorseResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection endorsement in UpdateEndorsement:EndorsementRepository.go: ", err.Error())
			return endorseResponse, err
		}
		return endorseResponse, nil
	} else {
		return endorseResponse, nil

	}
}

func (r *EndorsementRepository) UpDateBestCreators(userID string, update primitive.M) (models.Endorse, error) {
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("endorsement").FindOneAndUpdate(context.TODO(), bson.D{{Key: "publickey", Value: userID}}, update, &opt)
	var creatorDetails models.Endorse
	if rst != nil {
		err := rst.Decode((&creatorDetails))
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection endorsement in UpdateEndorsement:EndorsementRepository.go: ", err.Error())
			return creatorDetails, err
		}
		return creatorDetails, nil
	} else {
		return creatorDetails, nil

	}
}
func (r *EndorsementRepository) GetPaginatedBestCreators(filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, creators []models.CreatorInfo) (models.PaginatedCreatorInfo, error) {
	contentResponse, paginationResponse, err := repository.PaginateResponse[[]models.CreatorInfo](
		filterConfig,
		projectionData,
		pagesize,
		pageNo,
		collectionName,
		sortingFeildName,
		creators,
	)
	logs.InfoLogger.Println("content response: ", contentResponse)
	var response models.PaginatedCreatorInfo
	if err != nil {
		logs.InfoLogger.Println("Pagination failure:", err.Error())
		return response, err
	}
	response.ArtistInfo = contentResponse
	response.PaginationInfo = paginationResponse
	return response, nil
}
