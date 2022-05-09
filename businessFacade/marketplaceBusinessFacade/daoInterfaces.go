package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/database/repository/marketplaceRepository"
)

var (
	nftRepository       marketplaceRepository.NFTRepository
	ownershipRepository marketplaceRepository.OwnershipRepository
	watchListRepository marketplaceRepository.WatchListRepository
	userRepository      marketplaceRepository.UserRepository
	offerRepository     marketplaceRepository.OfferRepository
	reviewRepository    marketplaceRepository.ReviewRepository
	faqRepository       marketplaceRepository.FaqRepository
)
