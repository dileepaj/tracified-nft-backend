package nftComposerBusinessFacade

import (
	"encoding/base64"

	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services/htmlGeneretorService/htmlGenerator"
)

func GenerateHTMLFile(htmlJson models.HtmlGenerator) (string, error) {
	result, err := htmlGenerator.GenerateHTMLTemplate(htmlJson)
	if err != nil {
		return "", err
	} else {
		byteResult := base64.StdEncoding.EncodeToString([]byte(result))
		return byteResult, nil
	}
}
