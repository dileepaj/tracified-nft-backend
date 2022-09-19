package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/configs"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/gomail.v2"
)

func CreateFaq(faq models.Faq) (string, error) {
	return faqRepository.CreateFaq(faq)
}

func StoreUserFAQ(faq models.UserQuestions) (string, error) {
	return faqRepository.StoreUserFAQ(faq)
}

func GetAllFaq() ([]models.Faq, error) {
	return faqRepository.GetAllFaq()
}

func GetUserFAQByStatus(status string) ([]models.UserQuestions, error) {
	return faqRepository.GetUserFAQByStatus("status", status)
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

func UpdateUserFAQ(faq requestDtos.UpdateUserFAQ) (models.UserQuestions, error) {
	update := bson.M{
		"$set": bson.M{"status": faq.Status, "answer": faq.Answer},
	}
	return faqRepository.UpdateUserFAQ("userquestionID", faq.UserQuestionID, update)
}

func GetUserFAQByID(id primitive.ObjectID) (models.UserQuestions, error) {
	return faqRepository.FindUserFAQbyID(id)
}

func SendResponseToFAQ(faq models.UserQuestions) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", configs.GetChatSenderEmailAddres())
	msg.SetHeader("To", faq.UserMail)
	msg.SetHeader("Subject", "Tracified Marketplace FAQ Response")
	msg.SetBody("text/html", configs.GetAcceptedFAQEmail(faq.Category, faq.Subject, faq.Description, faq.Answer))
	faqEmail := gomail.NewDialer(
		configs.GetEmailHost(),
		configs.GetEmailPort(),
		configs.GetChatSenderEmailAddres(),
		configs.GetChatSenderEmailKey())
	if err := faqEmail.DialAndSend(msg); err != nil {
		logs.ErrorLogger.Println("Failed to send FAQ email: ", err.Error())
		return err
	}
	logs.InfoLogger.Println("endorsment email sent to :", faq.UserMail)
	return nil
}
