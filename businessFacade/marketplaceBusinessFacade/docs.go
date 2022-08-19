package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateDocs(docs models.Document) (string, error) {
	return documentRepository.CreateDocs(docs)
}
func GetAllDocs() ([]models.Document, error) {
	return documentRepository.GetAllDocs()
}
func GetDocsByID(docsID string) (models.Document, error) {

	return documentRepository.GetDocsByID(docsID)
}
func UpdateDocsbyID(docs requestDtos.UpdateDoc) (models.Document, error) {
	update := bson.M{
		"$set": bson.M{"topic": docs.Topic, "answers": docs.Answers},
	}
	return documentRepository.UpdateDocsbyID("_id", docs.TopicID, update)
}
