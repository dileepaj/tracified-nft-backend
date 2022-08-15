package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func StoreEndorse(createEndorseObject models.Endorse) (string, error) {
	rst, err1 := EndorsementRepository.SaveEndorsement(createEndorseObject)
	if err1 != nil {
		return "Endorsement not saved", err1
	}
	return rst, nil
}

func GetEndorsementByStatus(status string) ([]models.Endorse, error) {
	return EndorsementRepository.GetEndorsementByStatus("status", status)
}

func GetEndorsedStatus(publickey string) (models.Endorse, error) {
	return EndorsementRepository.FindEndorsermentbyPK(publickey)
}

func UpdateEndorsement(endorse requestDtos.UpdateEndorsementByPublicKey) (responseDtos.ResponseEndorsementUpdate, error) {
	update := bson.M{
		"$set": bson.M{"rating": endorse.Rating, "review": endorse.Review, "status": endorse.Status},
	}
	return EndorsementRepository.UpdateEndorsement("publickey", endorse.PublicKey, update)
}

func UpdateSetEndorsement(endorse requestDtos.UpdateEndorsement) (models.Endorse, error) {
	update := bson.M{
		"$set": bson.M{"name": endorse.Name, "email": endorse.Email, "contact": endorse.Contact},
	}
	return EndorsementRepository.UpdateSetEndorsement("publickey", endorse.PublicKey, update)
}

func UpdateSetEndorsement(endorse requestDtos.UpdateEndorsement) (models.Endorse, error) {
	update := bson.M{
		"$set": bson.M{"name": endorse.Name, "email": endorse.Email, "contact": endorse.Contact, "desc": endorse.Description},
	}
	return EndorsementRepository.UpdateSetEndorsement("publickey", endorse.PublicKey, update)
}
