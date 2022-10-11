package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/configs"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/gomail.v2"
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

func GetEndorsmentByUserPK(pk string) (models.Endorse, error) {
	return EndorsementRepository.FindEndorsermentbyPK(pk)
}

func UpdateEndorsement(endorse requestDtos.UpdateEndorsementByPublicKey) (responseDtos.ResponseEndorsementUpdate, error) {
	update := bson.M{
		"$set": bson.M{"rating": endorse.Rating, "review": endorse.Review, "status": endorse.Status},
	}
	return EndorsementRepository.UpdateEndorsement("publickey", endorse.PublicKey, "email", endorse.Email, update)
}

func UpdateSetEndorsement(endorse requestDtos.UpdateEndorsement) (models.Endorse, error) {
	update := bson.M{
		"$set": bson.M{"name": endorse.Name, "email": endorse.Email, "contact": endorse.Contact},
	}
	return EndorsementRepository.UpdateSetEndorsement("publickey", endorse.PublicKey, update)
}

func SendEndorsmentEmail(endorsment models.Endorse) error {
	if endorsment.Status == "Accepted" {
		msg := gomail.NewMessage()
		msg.SetHeader("From", configs.GetEndrosmentSenderEmailAddres())
		msg.SetHeader("To", endorsment.Email)
		msg.SetHeader("Subject", "Tracified Marketplace Endorsment Response")
		msg.SetBody("text/html", configs.GetAcceptedEndorsmentEmail(endorsment.Name, endorsment.Rating, endorsment.Review))
		endorsmentEmail := gomail.NewDialer(
			configs.GetEmailHost(),
			configs.GetEmailPort(),
			configs.GetEndrosmentSenderEmailAddres(),
			configs.GetEndorsmentSenderEmailKey())
		if err := endorsmentEmail.DialAndSend(msg); err != nil {
			logs.ErrorLogger.Println("Failed to send Endrosment email: ", err.Error())
			return err
		}
		logs.InfoLogger.Println("endorsment email sent to :", endorsment.Email)
		return nil
	} else if endorsment.Status == "Declined" {
		msg := gomail.NewMessage()
		msg.SetHeader("From", configs.GetEndrosmentSenderEmailAddres())
		msg.SetHeader("To", endorsment.Email)
		msg.SetHeader("Subject", "Tracified Marketplace Endorsment Response")
		msg.SetBody("text/html", configs.GetDeclinedEndorsmentEmail(endorsment.Name, endorsment.Rating, endorsment.Review))
		endorsmentEmail := gomail.NewDialer(
			configs.GetEmailHost(),
			configs.GetEmailPort(),
			configs.GetEndrosmentSenderEmailAddres(),
			configs.GetEndorsmentSenderEmailKey())
		if err := endorsmentEmail.DialAndSend(msg); err != nil {
			logs.ErrorLogger.Println("Failed to send Endrosment email: ", err.Error())
			return err
		}
	}
	return nil
}
