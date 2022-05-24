package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/models"
)

func CreateSVG(svg models.SVG) (string, error) {
	return CollectionRepository.SaveSVG(svg)
}
