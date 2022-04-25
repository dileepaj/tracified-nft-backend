package marketplaceRepository

import (
	"context"
	"time"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
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
	findOptions.SetLimit(10)
	result, err := connections.Connect().Collection("review").Find(context.TODO(), bson.D{{}}, findOptions)
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

func (r *ReviewRepository) UpdateReviewStatus(review requestDtos.UpdateReviewStatus) (responseDtos.ResponseReviewStatusUpdate, error) {
	var responseReviewStatus responseDtos.ResponseReviewStatusUpdate
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	update := bson.M{
		"$set": bson.M{"status": review.Status},
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := connections.Connect().Collection("review").FindOneAndUpdate(ctx, bson.M{"_id": review.Id}, update, &opt).Decode(&responseReviewStatus)
	if err != nil {
		logs.ErrorLogger.Println("Error occured when executing FindOneAndUpdate Query in UpdateReviewStatus(reviewRepository): ", err.Error())
	}
	return responseReviewStatus, err

}

func (r *ReviewRepository) DeleteReview(review requestDtos.DeleteReview) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := connections.Connect().Collection("review").DeleteOne(ctx, bson.M{"_id": review.Id})
	if err != nil {
		logs.ErrorLogger.Println("Error occured when Connecting to DB and executing DeleteOne Query in DeleteReview(reviewRepository): ", err.Error())
	}
	return err
}
