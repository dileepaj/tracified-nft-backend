package ipfsRepository

import (
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

type IpfsRepository struct{}

var Collection = "ipfsfiles"

func (r *IpfsRepository) SaveFileDetails(ipfsObj models.TracifiedDataPacket) (string, error) {
	return repository.Save[models.TracifiedDataPacket](ipfsObj, Collection)
}

func (r *IpfsRepository) GetTdpDetails(key string, id string) (models.TracifiedDataPacket, error) {
	var tdpDetailsArray models.TracifiedDataPacket
	rst := repository.FindOne(key, id, Collection)
	if rst != nil {
		err := rst.Decode(&tdpDetailsArray)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return tdpDetailsArray, err
		}
		return tdpDetailsArray, nil
	} else {
		return tdpDetailsArray, nil
	}
}
