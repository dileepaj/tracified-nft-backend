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

type FaqRepository struct{}

var Faq = "faq"
var userFaq = "userFAQ"

func (r *FaqRepository) CreateFaq(faq models.Faq) (string, error) {
	return repository.Save(faq, Faq)
}

func (r *FaqRepository) StoreUserFAQ(faq models.UserQuestions) (string, error) {
	return repository.Save(faq, userFaq)
}

func (r *FaqRepository) GetUserFAQByStatus(idName string, id string) ([]models.UserQuestions, error) {
	var faq []models.UserQuestions
	rst, err := repository.FindById(idName, id, userFaq)
	if err != nil {
		return faq, err
	}
	for rst.Next(context.TODO()) {
		var faqs models.UserQuestions
		err = rst.Decode(&faqs)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return faq, err
		}
		faq = append(faq, faqs)
	}
	return faq, nil
}

func (r *FaqRepository) GetAllFaq() ([]models.Faq, error) {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session in getAllFaq : faqRepository.go : ", err.Error())
	}
	defer session.EndSession(context.TODO())

	var allFaq []models.Faq
	findOptions := options.Find()
	result, err := session.Client().Database(connections.DbName).Collection(Faq).Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println("Error occured when trying to connect to DB and excute Find query in GetAllFaq:faqRepository.go: ", err.Error())
		return allFaq, err
	}
	for result.Next(context.TODO()) {
		var faq models.Faq
		err = result.Decode(&faq)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection faq in GetAllFaq:faqRepository.go: ", err.Error())
			return allFaq, err
		}
		allFaq = append(allFaq, faq)
	}
	return allFaq, nil
}

func (r *FaqRepository) GetFaqByID(questionID string) (models.Faq, error) {
	var faq models.Faq

	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	objectId, err := primitive.ObjectIDFromHex(questionID)
	if err != nil {
		logs.WarningLogger.Println("Error Occured when trying to convert hex string in to Object(ID) in GetFaqByID : faqRepository: ", err.Error())
	}
	rst, err := session.Client().Database(connections.DbName).Collection("faq").Find(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return faq, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&faq)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection faq in GetFaqByID:faqRepository.go: ", err.Error())
			return faq, err
		}
	}
	return faq, err
}

func (r *FaqRepository) UpdateFaqbyID(findBy string, id primitive.ObjectID, update primitive.M) (models.Faq, error) {
	var faqResponse models.Faq

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
	rst := session.Client().Database(connections.DbName).Collection("faq").FindOneAndUpdate(context.TODO(), bson.M{"_id": id}, update, &opt)
	if rst != nil {
		err := rst.Decode((&faqResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection faq in UpdateFaqbyID:faqRepository.go: ", err.Error())
			return faqResponse, err
		}
		return faqResponse, nil
	} else {
		return faqResponse, nil

	}
}

func (r *FaqRepository) UpdateUserFAQ(findBy string, id primitive.ObjectID, update primitive.M) (models.UserQuestions, error) {
	var userfaqResponse models.UserQuestions

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
	rst := session.Client().Database(connections.DbName).Collection("userFAQ").FindOneAndUpdate(context.TODO(), bson.M{"_id": id}, update, &opt)
	if rst != nil {
		err := rst.Decode((&userfaqResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection userFAQ in UpdateUserFAQ:FAQRepository.go: ", err.Error())
			return userfaqResponse, err
		}
		return userfaqResponse, nil
	} else {
		return userfaqResponse, nil

	}
}

func (r *FaqRepository) FindUserFAQbyID(id primitive.ObjectID) (models.UserQuestions, error) {
	var faq models.UserQuestions
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	rst, err := session.Client().Database(connections.DbName).Collection("userFAQ").Find(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return faq, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&faq)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection faq in GetUserFAQByID:FAQRepository.go: ", err.Error())
			return faq, err
		}
	}
	return faq, err
}
