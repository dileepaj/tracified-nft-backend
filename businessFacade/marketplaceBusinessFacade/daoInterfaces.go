package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/database/repository/marketplaceRepository"
)

var (
	nftRepository        marketplaceRepository.NFTRepository
	ownershipRepository  marketplaceRepository.OwnershipRepository
	watchListRepository  marketplaceRepository.WatchListRepository
	userRepository       marketplaceRepository.UserRepository
	offerRepository      marketplaceRepository.OfferRepository
	newsletterRepository marketplaceRepository.NewsLetterRepository
	faqRepository        marketplaceRepository.FaqRepository
	reviewRepository     marketplaceRepository.ReviewRepository
	CollectionRepository marketplaceRepository.CollectionRepository
	FavouriteRepository  marketplaceRepository.FavouriteRepository
)
