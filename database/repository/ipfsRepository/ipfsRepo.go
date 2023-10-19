package ipfsRepository

import (
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

type IpfsRepository struct{}

var Collection = "ipfsfiles"

func (r *IpfsRepository) SaveFileDetails(ipfsObj models.InsertTdpDetails) (string, error) {
	return repository.Save[models.InsertTdpDetails](ipfsObj, Collection)
}

func (r *IpfsRepository) GetTdpDetails(tenetId string) (models.InsertTdpDetails, error) {
	var currentTdpDetails models.InsertTdpDetails
	rst := repository.FindOne(tenetId, "", Collection)
	if rst != nil {
		err := rst.Decode(&currentTdpDetails)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return currentTdpDetails, err
		}
		return currentTdpDetails, nil
	} else {
		return currentTdpDetails, nil
	}
}
