package controllers

import (
	"github.com/dileepaj/tracified-nft-backend/database/repository/userRepository"
	"github.com/dileepaj/tracified-nft-backend/models"
)

type UserRepository struct {
	NftRepo userRepository.Repository
}

func (repository *UserRepository) CreateUser(user *models.User) {
	repository.NftRepo.Save(user)
}

func (repository *UserRepository) UpdateUser(id string, user *models.User) {
	repository.NftRepo.Update(id, user)
}

func (repository *UserRepository) UpdateAccount(id string, account , blockchain string) {
	repository.NftRepo.UpdateOne(id, account, blockchain)
}