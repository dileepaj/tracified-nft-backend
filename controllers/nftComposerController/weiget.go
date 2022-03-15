package nftcomposercontroller

import (
	nftcomposerrepository "github.com/dileepaj/tracified-nft-backend/database/repository/nftComposerRepository"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var weigetRepository nftcomposerrepository.WeigetRepository

func SaveWeigetList(weigets []models.Weiget) (string, error) {
	return weigetRepository.SaveWeigetList(weigets)
}

func SaveWeiget(weiget models.Weiget) (string, error) {
	return weigetRepository.SaveWeiget(weiget)
}

func FindWeigetAndUpdateQuery(weiget requestDtos.RequestWeiget)(models.Weiget,error){
	return weigetRepository.FindWeigetAndUpdate(weiget)
}

func FindWeigetById(id primitive.ObjectID)(models.Weiget,error){
	return weigetRepository.FindWeigetById(id)
}