package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreatePartner(partner models.Partner) (string, error) {
	return partnerRepository.CreatePartner(partner)
}
func GetAllPartner() ([]models.Partner, error) {
	return partnerRepository.GetAllPartner()
}
func GetPartnerByID(partnerID string) (models.Partner, error) {

	return partnerRepository.GetPartnerByID(partnerID)
}
func UpdatePartnerbyID(partner requestDtos.UpdatePartner) (models.Partner, error) {
	update := bson.M{
		"$set": bson.M{"companyname": partner.CompanyName, "weblink": partner.WebLink, "desc": partner.Description, "image": partner.Image, "topic": partner.Topic},
	}
	return partnerRepository.UpdatePartnerbyID("_id", partner.ID, update)
}
