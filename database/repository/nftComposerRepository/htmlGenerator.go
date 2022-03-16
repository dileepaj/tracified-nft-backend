package nftcomposerrepository

import (
	"context"
	"fmt"
	"time"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type HTMLNFTRepository struct{}

/**
Save the Json tha used to create HTML file of NFT
**/
func (r *HTMLNFTRepository)SaveHtmlData(htmlData models.HtmlGenerator) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rst, err := connections.Connect().Collection("htmlnft").InsertOne(ctx, htmlData)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return htmlData.Id.String(), err
	}
	var id = rst.InsertedID.(primitive.ObjectID)
	return id.String(), nil
}

func (r *HTMLNFTRepository)GetRecentProjectsByUserId(userId string)([]responseDtos.ResponseProject,error){
	var projects []responseDtos.ResponseProject
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	//findOptions.SetProjection(bson.M{"projectname": 1,"projectid":1 ,"timestamp":1, "_id": 0,"nftcontent":0})
	rst, err := connections.Connect().Collection("htmlnft").Find(context.TODO(), bson.D{{"userid", userId}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return projects, err
	}
	for rst.Next(context.TODO()) {
		var project responseDtos.ResponseProject
		err = rst.Decode(&project)
		if err != nil {
			logs.ErrorLogger.Println(err.Error())
			return projects, err
		}
		fmt.Println("ssss",project)
		projects = append(projects, project)
	}
	return projects, nil
}