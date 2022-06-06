package marketplaceBusinessFacade

import (
	"github.com/dileepaj/tracified-nft-backend/models"
)

func CreateUser(user models.User) (string, error) {
	return userRepository.SaveUser(user)
}

func GetBCAccountPKByUserId(userId string) ([]string, error) {
	return userRepository.FindBCAccountPKByUserId(userId)
}

func GetBCAccountPKByTenetName(tenentName string) ([]string, error) {
	return userRepository.FindBCAccountPKByTenentName(tenentName)
}
