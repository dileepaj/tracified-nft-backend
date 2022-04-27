package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
)

/*
Functions mentioned below passes data to the reviewrepository to execute respective queries.
*/
func CreateReview(review models.Review) (string, error) {
	return reviewRepository.CreateReview(review)
}
func GetReviewByNFT(nftidentifier string) ([]models.Review, error) {
	return reviewRepository.FindReviewByNFT(nftidentifier)
}
func GetAllReviews() ([]models.Review, error) {
	return reviewRepository.GetAllReviews()
}
func UpdateReviewStatus(review requestDtos.UpdateReviewStatus) (models.Review, error) {
	return reviewRepository.UpdateReviewStatus(review)
}
func DeleteReview(review requestDtos.DeleteReview) error {
	return reviewRepository.DeleteReview(review)
}
