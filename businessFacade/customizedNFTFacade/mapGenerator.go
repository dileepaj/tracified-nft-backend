package customizedNFTFacade

import (
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services/mapGenerator"
)

func GetMap(mapdata []models.MapInfo) string {
	rst := mapGenerator.GenerateMap(mapdata)
	return rst
}
