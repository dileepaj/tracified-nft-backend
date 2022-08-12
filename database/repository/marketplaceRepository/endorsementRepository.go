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

func (r *EndorsementRepository) UpdateEndorsement(findBy string, id string, update primitive.M) (responseDtos.ResponseEndorsementUpdate, error) {
	var endorseResponse responseDtos.ResponseEndorsementUpdate

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
	rst := session.Client().Database(connections.DbName).Collection("endorsement").FindOneAndUpdate(context.TODO(), bson.M{"publickey": id}, update, &opt)
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
	return repository.Save[models.Endorse](endorse, Endorsement)
}

func (r *EndorsementRepository) FindEndorsermentbyPK(publickey string) (models.Endorse, error) {
	var endorse models.Endorse
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	rst, err := session.Client().Database(connections.DbName).Collection("endorsement").Find(context.TODO(), bson.M{"publickey": publickey})
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
	rst := session.Client().Database(connections.DbName).Collection("endorsement").FindOneAndUpdate(context.TODO(), bson.M{"publickey": id}, update, &opt)
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
