package nftcomposercontroller

import (
	nftcomposerrepository "github.com/dileepaj/tracified-nft-backend/database/repository/nftComposerRepository"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var repository nftcomposerrepository.HTMLNFTRepository

func SaveHtmlContentData(htmlData models.HtmlGenerator) (string, error) {
	return repository.SaveHtmlData(htmlData)
}