package marketplaceRepository

import (
	"context"

	//"github.com/chebyrash/promise"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct{}

var User = "user"

func (r *UserRepository) FindById() {}

// User id --> find useraccount --> extract BC account from it

//Get Last NFT

func (r *UserRepository) FindBCAccountPKByUserId(id string) ([]string, error) {
	var accounts []string
	// convert id string to ObjectId
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logs.ErrorLogger.Println("Invalid id")
		return accounts, nil
	}
	rst := repository.FindOne[primitive.ObjectID]("_id", objectId, "user")
	if rst != nil {
		var user models.User
		err = rst.Decode(&user)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return accounts, err
		}
		for i := 0; i < len(user.BCAccounts); i++ {
			accounts = append(accounts, user.BCAccounts[i].Address)
		}
		return accounts, nil
	} else {
		return accounts, nil
	}
}

// User tenenetanme --> find many useraccounts   --> extract BC account from it
func (r *UserRepository) FindBCAccountPKByTenentName(tenetName string) ([]string, error) {
	var accounts []string
	rst, err := repository.FindById("tenentname", tenetName, User)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return accounts, err
	} else {
		for rst.Next(context.TODO()) {
			var user models.User
			err = rst.Decode(&user)
			if err != nil {
				logs.ErrorLogger.Println(err.Error())
				return accounts, err
			}
			for i := 0; i < len(user.BCAccounts); i++ {
				accounts = append(accounts, user.BCAccounts[i].Address)
			}
		}
		return accounts, nil
	}
}

func (r *UserRepository) SaveUser(user models.User) (string, error) {
	return repository.Save[models.User](user, User)
}
