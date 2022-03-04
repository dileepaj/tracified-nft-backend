package controllers

import (
	"github.com/dileepaj/tracified-nft-backend/database/repository/watchListRepository"
	"github.com/dileepaj/tracified-nft-backend/models"
)

var Repository watchListRepository.WatchListRepository

func CreateWatchList(watchList models.WatchList) (string, error) {
	return Repository.Save(watchList)
}

func FindNFTIdentifieryByUserId(userId string)([]string,error){
	return Repository.FindNFTIdentifieryByUserId(userId)
}