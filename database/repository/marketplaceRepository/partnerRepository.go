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

type PartnerRepository struct{}

var Partner = "partner"

func (r *PartnerRepository) CreatePartner(partner models.Partner) (string, error) {
	return repository.Save(partner, Partner)
}

func (r *PartnerRepository) GetAllPartner() ([]models.Partner, error) {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session in getAllPartner : partnerRepository.go : ", err.Error())
	}
	defer session.EndSession(context.TODO())

	var allPartner []models.Partner
	findOptions := options.Find()
	findOptions.SetLimit(10)
	result, err := session.Client().Database(connections.DbName).Collection(Partner).Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println("Error occured when trying to connect to DB and excute Find query in GetAllPartner:partnerRepository.go: ", err.Error())
		return allPartner, err
	}
	for result.Next(context.TODO()) {
		var partner models.Partner
		err = result.Decode(&partner)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection partner in GetAllPartner:partnerRepository.go: ", err.Error())
			return allPartner, err
		}
		allPartner = append(allPartner, partner)
	}
	return allPartner, nil
}

func (r *PartnerRepository) GetPartnerByID(partnerID string) (models.Partner, error) {
	var partner models.Partner

	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	objectId, err := primitive.ObjectIDFromHex(partnerID)
	if err != nil {
		logs.WarningLogger.Println("Error Occured when trying to convert hex string in to Object(ID) in GetPartnerByID : partnerRepository: ", err.Error())
	}
	rst, err := session.Client().Database(connections.DbName).Collection("partner").Find(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return partner, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&partner)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection partner in GetPartnerByID:partnerRepository.go: ", err.Error())
			return partner, err
		}
	}
	return partner, err
}

func (r *PartnerRepository) UpdatePartnerbyID(findBy string, id primitive.ObjectID, update primitive.M) (models.Partner, error) {
	var partnerResponse models.Partner

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
	rst := session.Client().Database(connections.DbName).Collection("partner").FindOneAndUpdate(context.TODO(), bson.M{"_id": id}, update, &opt)
	if rst != nil {
		err := rst.Decode((&partnerResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection partner in UpdatePartnerbyID:partnerRepository.go: ", err.Error())
			return partnerResponse, err
		}
		return partnerResponse, nil
	} else {
		return partnerResponse, nil

	}
}
