package userRepository

import "github.com/dileepaj/tracified-nft-backend/models"

type Repository interface {
	Count() (int64, error)
	FindById(idName string, id string) (models.User, error)
	Save(payload *models.User) error
	Update(id string, payload *models.User) error
	UpdateOne(id string, account string, blockchain string) error
	Delete(id string) error
}