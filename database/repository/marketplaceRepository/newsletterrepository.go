package marketplaceRepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NewsLetterRepository struct{}

var NewsLetter = "newsletter"

func (r *NewsLetterRepository) SaveNewsLetter(newsLetter models.NewsLetter) (string, error) {
	//Calling the common repository save method and passing model class and collection name.
	return repository.Save[models.NewsLetter](newsLetter, NewsLetter)
}

//retreving all nesletters from DB
func (r *NewsLetterRepository) GetAllNewsLetters() ([]models.NewsLetter, error) {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())

	var allNewsLetters []models.NewsLetter
	findOptions := options.Find()
	findOptions.SetLimit(10)
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

//Retreving news letters by author name
func (r *NewsLetterRepository) GetNewsLetterByAuthor(authorname string) ([]models.NewsLetter, error) {
	var newsLetters []models.NewsLetter
	findOptions := options.Find()
	findOptions.SetLimit(10)
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
