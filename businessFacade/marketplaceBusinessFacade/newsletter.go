package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/configs"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"gopkg.in/gomail.v2"
)

/*
Function below Calls the respective newsletterRepository methods and parses data accordingly
*/
func CreateNewsLetter(newsletter models.NewsLetter) (string, error) {
	return newsletterRepository.SaveNewsLetter(newsletter)
}

func AddSubscription(subscription models.Subscription) (string, error) {
	return newsletterRepository.AddSubscription(subscription)
}

func GetAllNewsLetters() ([]models.NewsLetter, error) {
	return newsletterRepository.GetAllNewsLetters()
}
func GetNewsLetterByAuthor(authorname string) ([]models.NewsLetter, error) {
	return newsletterRepository.GetNewsLetterByAuthor(authorname)
}
func GetNewsletterByID(newsletterID string) (models.NewsLetter, error) {
	return newsletterRepository.GetNewsletterbyID(newsletterID)
}

func SubscribeToNewsLetter(mail string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", configs.GetSubscribeSenderEmailAddres())
	msg.SetHeader("To", mail)
	msg.SetHeader("Subject", "Tracified Marketplace Subscription")
	msg.SetBody("text/html", configs.GetAcceptedSubscription())
	subscriptionEmail := gomail.NewDialer(
		configs.GetEmailHost(),
		configs.GetEmailPort(),
		configs.GetSubscribeSenderEmailAddres(),
		configs.GetSubscribeSenderEmailKey())
	if err := subscriptionEmail.DialAndSend(msg); err != nil {
		logs.ErrorLogger.Println("Failed to send Subscription email: ", err.Error())
		return err
	}
	logs.InfoLogger.Println("subscription email sent to :", mail)
	return nil
}

func CheckIfSubscribed(mail string) (string, error) {
	return newsletterRepository.CheckIfSubscribed(mail)
}
