package controllers

import (
	"github.com/dileepaj/tracified-nft-backend/database/repository/userRepository"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var Repostitory userRepository.UserRepository

func  CreateUser(user models.User)(string,error){
	return Repostitory.Save(user)
}

func GetBCAccountPKByUserId(userId string)([]string,error){
	return Repostitory.FindBCAccountPKByUserId(userId)
}

func GetBCAccountPKByTenetName(tenentName string)([]string,error){
	return Repostitory.FindBCAccountPKByTenentName(tenentName)
}

func  UpdateUser(id string, user *models.User) {
	//NftRepo.Update(id, user)
}

func  UpdateAccount(id string, account , blockchain string) {
	
}