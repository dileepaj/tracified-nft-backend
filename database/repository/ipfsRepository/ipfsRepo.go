package ipfsRepository

import (
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
)

type IpfsRepository struct{}

var Collection = "ipfsfiles"

func (r *IpfsRepository) SaveFileDetails(ipfsObj models.IpfsInsertObject) (string, error) {
	return repository.Save[models.IpfsInsertObject](ipfsObj, Collection)
}
