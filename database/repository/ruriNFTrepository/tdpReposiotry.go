package ruriNFTrepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Rurirepository struct{}

var tdpData = "tdpData"
var usernftmap = "userNFTMapping"

//!NOT IN USE
func (r *Rurirepository) SaveTDPbyBatchID(tdp models.TDP) (string, error) {
	return repository.Save(tdp, tdpData)
}

//!NOT IN USE
func (r *Rurirepository) GetTDPDatabyBatchID(batchID string) ([]models.TDP, error) {
	var tdpResponse []models.TDP
	findOptions := options.Find()
	findOptions.SetLimit(10)
	result, err := repository.FindById("identifier", batchID, tdpData)
	if err != nil {
		logs.ErrorLogger.Println("Error retreiving TDP data :", err.Error())
		return tdpResponse, err
	} else {
		for result.Next(context.TODO()) {
			var tdp models.TDP
			err = result.Decode(&tdp)
			if err != nil {
				logs.ErrorLogger.Println("Error while decoding data from DB: ", err.Error())
				return tdpResponse, err
			}
			tdpResponse = append(tdpResponse, tdp)
		}
		return tdpResponse, nil
	}
}
func (r *Rurirepository) SaveUserMapping(userNftMapping models.UserNFTMapping) (string, error) {
	return repository.Save(userNftMapping, usernftmap)
}
