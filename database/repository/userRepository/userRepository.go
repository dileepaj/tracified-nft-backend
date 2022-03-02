package userRepository

import (
	"context"
	"fmt"
	"time"

	"github.com/dileepaj/tracified-nft-backend/connections"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct{}

func (repository *UserRepository) FindById() {}

func (repository *UserRepository) FindBCAccountPKByUserId(id string) ([]string, error) {
	var accounts []string
	// convert id string to ObjectId
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logs.ErrorLogger.Println("Invalid id")
		return accounts, nil
	}
	rst := connections.Connect().Collection("user").FindOne(context.TODO(), bson.D{{"_id", objectId}})
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

func (repository *UserRepository) FindBCAccountPKByTenentName(tenetName string) ([]string, error) {
	var accounts []string
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	findOptions.SetProjection(bson.M{"bcaccounts": 1, "_id": 0})
	rst, err := connections.Connect().Collection("user").Find(context.TODO(), bson.D{{"tenentname", tenetName}}, findOptions)
	fmt.Println("aaaaaaaaaaq--------------------", rst.Next(context.TODO()), "+++++++++++++", err)
	if err != nil{
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

func (repository *UserRepository) Save(user models.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rst, err := connections.Connect().Collection("user").InsertOne(ctx, user)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return user.Email, err
	}
	var id = rst.InsertedID.(primitive.ObjectID)
	return id.String(), nil
}

func (repository *UserRepository) Update() {}

func (repository *UserRepository) UpdateOne() {}
