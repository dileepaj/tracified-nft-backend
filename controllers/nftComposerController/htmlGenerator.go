package nftcomposercontroller

import (
	nftcomposerrepository "github.com/dileepaj/tracified-nft-backend/database/repository/nftComposerRepository"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var repository nftcomposerrepository.HTMLNFTRepository

func SaveHtmlContentData(htmlData models.HtmlGenerator) (string, error) {
	return repository.SaveHtmlData(htmlData)
}

func GetRecntProjectByUser(userid string)([]responseDtos.ResponseProject,error){
	return repository.GetRecentProjectsByUserId(userid)
}