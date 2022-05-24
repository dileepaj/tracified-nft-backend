package marketplaceBusinessFacade

import (
	// "fmt"

	// "github.com/dileepaj/tracified-nft-backend/database/repository/marketplaceRepository"
	// "github.com/dileepaj/tracified-nft-backend/dtos/requestDtos"
	// "github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
)

func CreateFavourites(favs models.Favourite) (string, error) {
	return FavouriteRepository.SaveFavourite(favs)
}

func GetFavouritesByUserPK(userid string) ([]models.Favourite, error) {
	return FavouriteRepository.FindFavouritesbyUserPK("userid", userid)

}

func GetAllFavourites() ([]models.Favourite, error) {
	return FavouriteRepository.GetAllFavourites()
}
