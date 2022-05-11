package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateFaq(faq models.Faq) (string, error) {
	return faqRepository.CreateFaq(faq)
}
func GetAllFaq() ([]models.Faq, error) {
	return faqRepository.GetAllFaq()
}
func GetFaqByID(questionID string) (models.Faq, error) {

	return faqRepository.GetFaqByID(questionID)
}
func UpdateFaqbyID(faq requestDtos.UpdateFaq) (models.Faq, error) {
	update := bson.M{
		"$set": bson.M{"question": faq.Question, "answers": faq.Answers},
	}
	return faqRepository.UpdateFaqbyID("_id", faq.QuestionID, update)
}
