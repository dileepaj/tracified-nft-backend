package marketplaceRepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReviewRepository struct{}

var Review = "review"

func (r *ReviewRepository) CreateReview(review models.Review) (string, error) {
	return repository.Save(review, Review)
}
func (r *ReviewRepository) FindReviewByNFT(nftidentifier string) ([]models.Review, error) {
	var reviews []models.Review
	result, err := repository.FindById("nftidentifier", nftidentifier, Review)
	if err != nil {
		logs.ErrorLogger.Println("Error occured while executing FindById Query in FindReviewByNFT(reviewRepository):  ", err.Error())
		return reviews, err
	} else {
		for result.Next(context.TODO()) {
			var review models.Review
			err = result.Decode(&review)
			if err != nil {
				logs.ErrorLogger.Println("Error occured while retreving data from collection Reveiew in FindReviewByNFT(reviewRepository): ", err.Error())
				return reviews, err
			}
			reviews = append(reviews, review)
		}
		return reviews, nil
	}
}
func (r *ReviewRepository) GetAllReviews() ([]models.Review, error) {
	var allReviews []models.Review
	findOptions := options.Find()
	result, err := connections.GetSessionClient(Review).Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println("Error occured when trying to connect to DB and excute Find Query in GetAllReviews(reviewRepository): ", err.Error())
		return allReviews, err
	}
	for result.Next(context.TODO()) {
		var reviews models.Review
		err = result.Decode((&reviews))
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection Reveiew in FindReviewByNFT(reviewRepository): ", err.Error())
			return allReviews, err
		}
		allReviews = append(allReviews, reviews)
	}
	return allReviews, nil
}

func (r *ReviewRepository) UpdateReviewStatus(review requestDtos.UpdateReviewStatus) (models.Review, error) {
	var responseReviewStatus models.Review
	update := bson.M{
		"$set": bson.M{"status": review.Status},
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := connections.GetSessionClient("review").FindOneAndUpdate(context.TODO(), bson.M{"_id": review.Id}, update, &opt)
	if rst != nil {
		err := rst.Decode((&responseReviewStatus))
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return responseReviewStatus, err
		}
		return responseReviewStatus, nil
	} else {
		return responseReviewStatus, nil

	}
}

func (r *ReviewRepository) DeleteReview(review requestDtos.DeleteReview) error {
	result, err := connections.GetSessionClient("review").DeleteOne(context.TODO(), bson.M{"_id": review.Id})
	if err != nil {
		logs.ErrorLogger.Println("Error occured when Connecting to DB and executing DeleteOne Query in DeleteReview(reviewRepository): ", err.Error())
	}
	logs.InfoLogger.Println("review deleted :", result.DeletedCount)
	return err

}

func (r *ReviewRepository) GetReviewsbyFilter(filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, sortType int, reviews []models.ReviewsforPagination) (models.ReviewPaginatedResponse, error) {
	contentResponse, paginationResponse, err := repository.PaginateWithCustomSort[[]models.ReviewsforPagination](
		filterConfig,
		projectionData,
		pagesize,
		pageNo,
		collectionName,
		sortingFeildName,
		sortType,
		reviews,
	)
	var response models.ReviewPaginatedResponse
	if err != nil {
		logs.InfoLogger.Println("Pagination failure:", err.Error())
		return response, err
	}
	response.ReviewContent = contentResponse
	response.PaginationInfo = paginationResponse
	return response, nil
}
