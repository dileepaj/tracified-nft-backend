package marketplaceBusinessFacade

import (
	"encoding/base64"

	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services/svgGeneratorService/svgGenerator"
)

func GenerateSVGFileForRURI(svgJson models.SvgCreator) (string, error) {
	result, err := svgGenerator.GenerateSVGRuriTemplate(svgJson)
	if err != nil {
		return "", err
	} else {
		byteResult := base64.StdEncoding.EncodeToString([]byte(result))
		return byteResult, nil
	}
}

func GetRURISVGByBatchID(itemID string) ([]models.NFT, error) {
	return nftRepository.GetRURISVGByBatchID("itemID", itemID)
}
