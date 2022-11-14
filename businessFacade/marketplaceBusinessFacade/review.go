package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"go.mongodb.org/mongo-driver/bson"
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

func calcRating(list models.CreatorsList, Rating float32) models.CreatorsList {
	switch Rating {
	case 1:
		list.Star1Ratings += 1
		list.TotalStars += 1
	case 1.5:
		list.Star_1_5_Ratings += 1
		list.TotalStars += 1
	case 2:
		list.Star_2_Ratings += 1
		list.TotalStars += 1
	case 2.5:
		list.Star_2_5Ratings += 1
		list.TotalStars += 1
	case 3:
		list.Star_3_Ratings += 1
		list.TotalStars += 1
	case 3.5:
		list.Star_3_5_Ratings += 1
		list.TotalStars += 1
	case 4:
		list.Star_4_Ratings += 1
		list.TotalStars += 1
	case 4.5:
		list.Star_4_5_Ratings += 1
		list.TotalStars += 1
	case 5:
		list.Star_5_Ratings += 1
		list.TotalStars += 1
	}
	return list
}

func CalcAvgStarRating(creatorlist models.CreatorsList) models.CreatorsList {
	var onestar = float32(creatorlist.Star1Ratings) * 1
	var oneHalfstar = float32(creatorlist.Star_1_5_Ratings) * 1.5
	var twoStar = float32(creatorlist.Star_2_Ratings) * 2
	var twoHalfStar = float32(creatorlist.Star_2_5Ratings) * 2.5
	var threeStar = float32(creatorlist.Star_3_Ratings) * 3
	var threeHalfStar = float32(creatorlist.Star_3_5_Ratings) * 3.5
	var fourStar = float32(creatorlist.Star_4_Ratings) * 4
	var fourHalfStar = float32(creatorlist.Star_4_5_Ratings) * 4.5
	var fiveStar = float32(creatorlist.Star_5_Ratings) * 5
	var ans = onestar + oneHalfstar + twoStar + twoHalfStar + threeStar + threeHalfStar + fourStar + fourHalfStar + fiveStar
	creatorlist.AvgRating = ans / float32(creatorlist.TotalStars)
	return creatorlist
}

func GetBestCreators() ([]models.CreatorsList, error) {
	var bestCreators []models.CreatorsList
	var allCreators []models.CreatorsList
	result, err := GetAllReviews()
	if err != nil {
		return bestCreators, err
	}
	for i, element := range result {
		if len(allCreators) == 0 {
			var newCreator models.CreatorsList
			newCreator.NftIdentifier = result[i].NFTIdentifier
			newCreator.UserID = element.UserID
			res := calcRating(newCreator, element.Rating)
			allCreators = append(allCreators, res)
			continue
		}
		foundflag := false
		foundat := 0
		_ = foundat
		_ = foundflag
		for index := range allCreators {
			if element.NFTIdentifier == allCreators[index].NftIdentifier {
				foundflag = true
				foundat = index
				continue
			}

		}
		if foundflag {
			res := calcRating(allCreators[foundat], element.Rating)
			allCreators[foundat] = res
		} else {
			var newCreator models.CreatorsList
			newCreator.NftIdentifier = element.NFTIdentifier
			newCreator.UserID = element.UserID
			res := calcRating(newCreator, element.Rating)
			allCreators = append(allCreators, res)
		}

	}
	for _, element := range allCreators {
		res := CalcAvgStarRating(element)
		if res.AvgRating >= 1 {
			bestCreators = append(bestCreators, res)
		}
	}

	return bestCreators, nil
}

func GetReviewsbyFilter(reviewData requestDtos.ReviewFiltering) (models.ReviewPaginatedResponse, error) {
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "nftidentifier", Value: 1},
		{Key: "userid", Value: 1},
		{Key: "rating", Value: 1},
		{Key: "description", Value: 1},
		{Key: "timestamp", Value: 1},
	}
	var reviewPaginate []models.ReviewsforPagination
	var returnData models.ReviewPaginatedResponse
	filter := bson.M{
		"status":        "Pending",
		"nftidentifier": reviewData.NFTIdentifier,
	}
	if reviewData.Filterby == "high" {
		response, err := reviewRepository.GetReviewsbyFilter(filter, projection, reviewData.PageSize, reviewData.RequestedPage, "review", reviewData.Filterby, reviewData.FilterType, reviewPaginate)
		if err != nil {
			return returnData, err
		}
		returnData = response
		return returnData, nil
	} else if reviewData.Filterby == "low" {
		response, err := reviewRepository.GetReviewsbyFilter(filter, projection, reviewData.PageSize, reviewData.RequestedPage, "review", reviewData.Filterby, reviewData.FilterType, reviewPaginate)
		if err != nil {
			return returnData, err
		}
		returnData = response
		return returnData, nil
	}
	response, err := reviewRepository.GetReviewsbyFilter(filter, projection, reviewData.PageSize, reviewData.RequestedPage, "review", reviewData.Filterby, reviewData.FilterType, reviewPaginate)
	if err != nil {
		return returnData, err
	}
	returnData = response
	return returnData, nil

}
