package customizedNFTFacade

import (
	"github.com/dileepaj/tracified-nft-backend/database/repository/customizedNFTrepository"
	"github.com/dileepaj/tracified-nft-backend/database/repository/marketplaceRepository"
)

var (
	svgRepository customizedNFTrepository.SvgRepository
	otpRepository customizedNFTrepository.OtpRepository
	mapRepository customizedNFTrepository.MapRepository
	nftRepository marketplaceRepository.NFTRepository
)
