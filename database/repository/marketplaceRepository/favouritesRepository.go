package marketplaceRepository

import (
	"context"

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

func (r *FavouriteRepository) FindFavouritesbyUserPK(userpk string) (models.Favourite, error) {
	var favourite models.Favourite

	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	rst, err := session.Client().Database(connections.DbName).Collection("favourite").Find(context.TODO(), bson.M{"creatoruserid": userpk})
	if err != nil {
		return favourite, err
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&favourite)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection favourite in GetfavouriteByID:favouriteRepository.go: ", err.Error())
			return favourite, err
		}
	}
	return favourite, err
}
func (r *FavouriteRepository) GetAllFavourites() ([]models.Favourite, error) {
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session in getAllFavourite : FavouriteRepository.go : ", err.Error())
	}
	defer session.EndSession(context.TODO())

	var favourite []models.Favourite
	findOptions := options.Find()
	findOptions.SetLimit(10)
	result, err := session.Client().Database(connections.DbName).Collection(Favourite).Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println("Error occured when trying to connect to DB and excute Find query in GetAllFavourite:FavouriteRepository.go: ", err.Error())
		return favourite, err
	}
	for result.Next(context.TODO()) {
		var favourites models.Favourite
		err = result.Decode(&favourites)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection favourites in GetAllFavourites:favouritesRepository.go: ", err.Error())
			return favourite, err
		}
		favourite = append(favourite, favourites)
	}
	return favourite, nil
}
