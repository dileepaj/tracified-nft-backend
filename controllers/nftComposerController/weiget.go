package nftcomposercontroller

import (
	nftcomposerrepository "github.com/dileepaj/tracified-nft-backend/database/repository/nftComposerRepository"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var weigetRepository nftcomposerrepository.WeigetRepository

func SaveWeiget(weigets []models.Weight) (string, error) {
	return weigetRepository.SaveWeigetList(weigets)
}