package ruriNFTrepository

import (
	"context"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SvgRepository struct{}

var usernftmap = "userSVGMapping"

func (r *SvgRepository) SaveUserMapping(userNftMapping models.UserNFTMapping) (responseDtos.SVGforNFTResponse, error) {
	var response responseDtos.SVGforNFTResponse
	response.SVG = userNftMapping.SVG
	rst, err := repository.Save(userNftMapping, usernftmap)
	if err != nil {
		return response, err
	}
	response.SvgID = rst
	return response, nil
}

func (r *SvgRepository) UpdateUserMappingbySha256(fidby string, svgID primitive.ObjectID, update primitive.M) (responseDtos.SVGforNFTResponse, error) {
	var response responseDtos.SVGforNFTResponse
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := session.Client().Database(connections.DbName).Collection(usernftmap).FindOneAndUpdate(context.TODO(), bson.M{"_id": svgID}, update, &opt)
	if rst != nil {
		err := rst.Decode(&response)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreiving data from DB : ", err.Error())
			return response, err
		}
		return response, nil
	} else {
		return response, nil
	}
}

func (r *SvgRepository) GetSVGbySha256(hash string) (string, error) {
	var response responseDtos.SVGforNFTResponse
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	rst, err := session.Client().Database(connections.DbName).Collection(usernftmap).Find(context.TODO(), bson.M{"hash": hash})
	if err != nil {
		logs.ErrorLogger.Println("Error while retrevinf fata froom DB:", err.Error())
		return response.SVG, err
	}
	for rst.Next((context.TODO())) {
		err = rst.Decode(&response)
		if err != nil {
			logs.ErrorLogger.Println("Error while decoding retreived Data :", err.Error())
			return response.SVG, err
		}
	}
	return response.SVG, nil
}

//! REMOVE AFTER Final imply
//var tdpData = "tdpData"
// func (r *Rurirepository) SaveTDPbyBatchID(tdp models.TDP) (string, error) {
// 	return repository.Save(tdp, tdpData)
// }

// func (r *Rurirepository) GetGeoImageTDPData(batchID string) ([]models.TDP, error) {
// 	var dbResponse []models.TDP
// 	findOptions := options.Find()
// 	findOptions.SetLimit(10)
// 	logs.InfoLogger.Println("id passed:", batchID)
// 	result, err := repository.FindById("identifier", batchID, tdpData)
// 	if err != nil {
// 		logs.ErrorLogger.Println("Error retreiving TDP data :", err.Error())
// 		return dbResponse, err
// 	} else {
// 		for result.Next(context.TODO()) {
// 			var tracabilitydata models.TDP
// 			//var geodata models.GeoImage
// 			err = result.Decode(&tracabilitydata)
// 			logs.InfoLogger.Println("TDP data from DB :", dbResponse)
// 			if err != nil {
// 				logs.ErrorLogger.Println("Error while getting TDP data:", err.Error())
// 				return dbResponse, err
// 			} else {
// 				dbResponse = append(dbResponse, tracabilitydata)
// 			}
// 		}
// 		return dbResponse, nil
// 	}
// }

// //!NOT IN USE
// func (r *Rurirepository) GetTDPDatabyBatchID(batchID string) ([]models.TDP, error) {
// 	var tdpResponse []models.TDP
// 	findOptions := options.Find()
// 	findOptions.SetLimit(10)
// 	result, err := repository.FindById("identifier", batchID, tdpData)
// 	if err != nil {
// 		logs.ErrorLogger.Println("Error retreiving TDP data :", err.Error())
// 		return tdpResponse, err
// 	} else {
// 		for result.Next(context.TODO()) {
// 			var tdp models.TDP
// 			err = result.Decode(&tdp)
// 			if err != nil {
// 				logs.ErrorLogger.Println("Error while decoding data from DB: ", err.Error())
// 				return tdpResponse, err
// 			}
// 			tdpResponse = append(tdpResponse, tdp)
// 		}
// 		return tdpResponse, nil
// 	}
// }
