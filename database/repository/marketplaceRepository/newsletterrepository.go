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

type NewsLetterRepository struct{}

var NewsLetter = "newsletter"
var Subscription = "subscription"

func (r *NewsLetterRepository) SaveNewsLetter(newsLetter models.NewsLetter) (string, error) {
	//Calling the common repository save method and passing model class and collection name.
	return repository.Save(newsLetter, NewsLetter)
}

func (r *NewsLetterRepository) AddSubscription(subscription models.Subscription) (string, error) {
	//Calling the common repository save method and passing model class and collection name.
	return repository.Save(subscription, Subscription)
}

// retreving all nesletters from DB
func (r *NewsLetterRepository) GetAllNewsLetters() ([]models.NewsLetter, error) {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())

	var allNewsLetters []models.NewsLetter
	findOptions := options.Find()
	result, err := session.Client().Database(connections.DbName).Collection(NewsLetter).Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println("Error while retreving data in GetAllNewsLetters : newsletterrepository: ", err.Error())
	}
	for result.Next(context.TODO()) {
		var newsLetter models.NewsLetter
		err = result.Decode(&newsLetter)
		if err != nil {
			logs.ErrorLogger.Println("Error while decoding news letter data recived from DB in GetAllNewsLetters:newsletterrepository: ", err.Error())
			return allNewsLetters, err
		}
		allNewsLetters = append(allNewsLetters, newsLetter)
	}
	return allNewsLetters, err
}

// Retreving news letters by author name
func (r *NewsLetterRepository) GetNewsLetterByAuthor(authorname string) ([]models.NewsLetter, error) {
	var newsLetters []models.NewsLetter
	result, err := repository.FindById("author", authorname, NewsLetter)
	if err != nil {
		logs.ErrorLogger.Println("Error while getting data from DB in GetNewsLetterByAuthor:newsletterrepository: ", err.Error())
		return newsLetters, err
	} else {
		for result.Next(context.TODO()) {
			var newsletter models.NewsLetter
			err = result.Decode(&newsletter)
			if err != nil {
				logs.ErrorLogger.Println("Error while decoding news letter data recived from DB in GetNewsLetterByAuthor:newsletterrepository: ", err.Error())
				return newsLetters, err
			}
			newsLetters = append(newsLetters, newsletter)
		}
		return newsLetters, nil
	}
}
func (r *NewsLetterRepository) GetNewsletterbyID(newsLetterID string) (models.NewsLetter, error) {
	var newsletter models.NewsLetter

	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())

	objectID, err := primitive.ObjectIDFromHex(newsLetterID)
	if err != nil {
		logs.WarningLogger.Println("Error Occured when trying to convert hex string in to Object(ID) in getNewsletterbyID : NewsletterRepository: ", err.Error())
	}
	rst, err := session.Client().Database(connections.DbName).Collection("newsletter").Find(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return newsletter, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&newsletter)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection newsletter in getNewsletterbyID : NewsletterRepository:: ", err.Error())
			return newsletter, err
		}
	}
	return newsletter, err

}
