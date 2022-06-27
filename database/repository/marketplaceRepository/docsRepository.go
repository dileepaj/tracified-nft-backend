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

type DocumentRepository struct{}

var Document = "document"

func (r *DocumentRepository) CreateDocs(docs models.Document) (string, error) {
	return repository.Save(docs, Document)
}

func (r *DocumentRepository) GetAllDocs() ([]models.Document, error) {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session in getAllPartner : partnerRepository.go : ", err.Error())
	}
	defer session.EndSession(context.TODO())

	var allDocs []models.Document
	findOptions := options.Find()
	findOptions.SetLimit(10)
	result, err := session.Client().Database(connections.DbName).Collection(Document).Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println("Error occured when trying to connect to DB and excute Find query in GetAllDocs:documentRepository.go: ", err.Error())
		return allDocs, err
	}
	for result.Next(context.TODO()) {
		var docs models.Document
		err = result.Decode(&docs)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection partner in GetAllDocs:documentRepository.go: ", err.Error())
			return allDocs, err
		}
		allDocs = append(allDocs, docs)
	}
	return allDocs, nil
}

func (r *DocumentRepository) GetDocsByID(docsID string) (models.Document, error) {
	var docs models.Document

	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	objectId, err := primitive.ObjectIDFromHex(docsID)
	if err != nil {
		logs.WarningLogger.Println("Error Occured when trying to convert hex string in to Object(ID) in GetDocsByID : documentRepository: ", err.Error())
	}
	rst, err := session.Client().Database(connections.DbName).Collection("document").Find(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return docs, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&docs)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection document in GetDocumentByID:documentRepository.go: ", err.Error())
			return docs, err
		}
	}
	return docs, err
}

func (r *DocumentRepository) UpdateDocsbyID(findBy string, id primitive.ObjectID, update primitive.M) (models.Document, error) {
	var docsResponse models.Document

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
	rst := session.Client().Database(connections.DbName).Collection("document").FindOneAndUpdate(context.TODO(), bson.M{"_id": id}, update, &opt)
	if rst != nil {
		err := rst.Decode((&docsResponse))
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection document in UpdateDoumentbyID:documentRepository.go: ", err.Error())
			return docsResponse, err
		}
		return docsResponse, nil
	} else {
		return docsResponse, nil

	}
}
