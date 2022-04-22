package nftComposerBusinessFacade

import (
	"encoding/base64"

	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services/htmlGeneretorService/htmlGenerator"
)

// Generate htmlconvert it to base64 return base 64 encode html string(html file can not send via httmp call sometime may be change)
func GenerateHTMLFile(htmlJson models.HtmlGenerator) (string, error) {
	result, err := htmlGenerator.GenerateHTMLTemplate(htmlJson)
	if err != nil {
		return "", err
	} else {
		//encoded html to base64 
		byteResult := base64.StdEncoding.EncodeToString([]byte(result))
		return byteResult, nil
	}
}