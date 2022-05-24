package marketplaceRepository

import (
	"context"
	"fmt"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FavouriteRepository struct{}

var Favourite = "favourite"

func (r *FavouriteRepository) SaveFavourite(favourite models.Favourite) (string, error) {
	return repository.Save[models.Favourite](favourite, Favourite)
}

func (repository *FavouriteRepository) FindFavouritesbyUserPK(idName1 string, id1 string) ([]models.Favourite, error) {
	var favourites []models.Favourite
	if idName1 != "" {
		findOptions := options.Find()
		rst, err := connections.Connect().Collection("favourites").Find(context.TODO(), bson.D{{idName1, id1}}, findOptions)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return favourites, err
		}
		for rst.Next(context.TODO()) {
			var favourite models.Favourite
			err = rst.Decode((&favourite))
			if err != nil {
				logs.ErrorLogger.Println(err.Error())
				return favourites, err
			}
			favourites = append(favourites, favourite)
		}
		return favourites, nil
	} else {
		return favourites, nil
	}
}
func (repository *FavouriteRepository) GetAllFavourites() ([]models.Favourite, error) {
	fmt.Println("executing repo get all favourites")
	var favourites []models.Favourite
	findOptions := options.Find()
	findOptions.SetLimit(10)
	rst, err := connections.Connect().Collection("favourite").Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return favourites, err
	}
	fmt.Println("outside loop")
	for rst.Next(context.TODO()) {
		var favourite models.Favourite
		err = rst.Decode((&favourite))
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return favourites, err
		}
		fmt.Println("inside loop : ", favourite)
		favourites = append(favourites, favourite)
	}
	return favourites, nil
}
