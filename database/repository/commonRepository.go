package repository

import (
	"context"
	"os"

	"github.com/dileepaj/tracified-nft-backend/commons"
	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	paginate "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Save[T models.SaveType](model T, collection string) (string, error) {
	rst, err := connections.GetSessionClient(collection).InsertOne(context.TODO(), model)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return "", err
	}
	id := rst.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func InsertMany[T models.InsertManyType](model T, collection string) (string, error) {
	var docs []interface{}
	for _, t := range model {
		docs = append(docs, t)
	}
	rst, err := connections.GetSessionClient(collection).InsertOne(context.TODO(), model)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return "Error while inserting widgets", err
	}
	id := rst.InsertedID.(primitive.ObjectID)
	return id.String(), nil
}

func FindById(idName string, id string, collection string) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	findOptions.SetProjection(bson.M{"otp": 0})
	rst, err := connections.GetSessionClient(collection).Find(context.TODO(), bson.D{{idName, id}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

func FindOne[T models.FindOneType](idName string, id T, collection string) *mongo.SingleResult {
	rst := connections.GetSessionClient(collection).FindOne(context.TODO(), bson.D{{idName, id}})
	return rst
}

func FindById1AndId2(idName1 string, id1 string, idName2 string, id2 string, collection string) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	rst, err := connections.GetSessionClient(collection).Find(context.TODO(), bson.D{{idName1, id1}, {idName2, id2}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

func FindById1AndNotId2(idName1 string, id1 string, idName2 string, id2 string, collection string) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	rst, err := connections.GetSessionClient(collection).Find(context.TODO(), bson.D{{idName1, id1}, {idName2, id2}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

func FindById1Id2Id3(idName1 string, id1 string, idName2 string, id2 string, idName3 string, id3 string, collection string) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	rst, err := connections.GetSessionClient(collection).Find(context.TODO(), bson.D{{idName1, id1}, {idName2, id2}, {idName3, id3}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

func FindByFieldInMultipleValus(fields string, tags []string, collection string) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	rst, err := connections.GetSessionClient(collection).Find(context.TODO(), bson.D{{fields, bson.D{{"$in", tags}}}}, findOptions)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return rst, err
	} else {
		return rst, nil
	}
}

func FindOneAndUpdate(findBy string, value string, update primitive.M, projectionData primitive.M, collection string) *mongo.SingleResult {
	after := options.After
	projection := projectionData
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Projection:     &projection,
	}
	rst := connections.GetSessionClient(collection).FindOneAndUpdate(context.TODO(), bson.M{findBy: value}, update, &opt)
	return rst
}

func Remove(idName string, id, collection string) (int64, error) {
	result, err := connections.GetSessionClient(collection).DeleteMany(context.TODO(), bson.M{idName: id})
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
		return 0, err
	}
	return result.DeletedCount, nil
}

type paginateResponseType interface {
	[]models.NFTContentforMatrix | []models.CreatorInfo | []models.ReviewsforPagination
}

func PaginateResponse[PaginatedData paginateResponseType](filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, object PaginatedData) (PaginatedData, models.PaginationTemplate, error) {
	var paginationdata models.PaginationTemplate
	ctx := context.Background()
	DbName := commons.GoDotEnvVariable("DATABASE_NAME")
	connectionString := os.Getenv("BE_MONGOLAB_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		logs.ErrorLogger.Println("failed to connect to DB: ", err.Error())
	}
	defer client.Disconnect(ctx);
	dbConnection := client.Database(DbName)
	filter := filterConfig
	limit := int64(pagesize)
	page := int64(pageNo)
	collection := dbConnection.Collection(collectionName)
	projection := projectionData
	// var nfts []models.PaginateResponseMatrix
	paginatedData, paginateerr := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Sort(sortingFeildName, -1).Select(projection).Filter(filter).Decode(&object).Find()
	paginationdata.TotalElements = int32(paginatedData.Pagination.Total)
	paginationdata.TotalPages = int32(paginatedData.Pagination.TotalPage)
	paginationdata.Currentpage = int32(paginatedData.Pagination.Page)
	paginationdata.PageSize = int32(paginatedData.Pagination.PerPage)
	paginationdata.Previouspage = int32(paginatedData.Pagination.Prev)
	paginationdata.NextPage = int32(paginatedData.Pagination.Next)
	if paginateerr != nil {
		logs.ErrorLogger.Println("Pagination failure :", paginateerr.Error())
		return object, paginationdata, err
	}
	return object, paginationdata, nil
}

func PaginateWithCustomSort[PaginatedData paginateResponseType](filterConfig bson.M, projectionData bson.D, pagesize int32, pageNo int32, collectionName string, sortingFeildName string, sortyType int, object PaginatedData) (PaginatedData, models.PaginationTemplate, error) {
	var paginationdata models.PaginationTemplate
	ctx := context.Background()
	DbName := commons.GoDotEnvVariable("DATABASE_NAME")
	connectionString := os.Getenv("BE_MONGOLAB_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		logs.ErrorLogger.Println("failed to connect to DB: ", err.Error())
	}
	defer client.Disconnect(ctx);
	dbConnection := client.Database(DbName)
	filter := filterConfig
	limit := int64(pagesize)
	page := int64(pageNo)
	collection := dbConnection.Collection(collectionName)
	projection := projectionData
	paginatedData, err := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Sort(sortingFeildName, sortyType).Select(projection).Filter(filter).Decode(&object).Find()
	paginationdata.TotalElements = int32(paginatedData.Pagination.Total)
	paginationdata.TotalPages = int32(paginatedData.Pagination.TotalPage)
	paginationdata.Currentpage = int32(paginatedData.Pagination.Page)
	paginationdata.PageSize = int32(paginatedData.Pagination.PerPage)
	paginationdata.Previouspage = int32(paginatedData.Pagination.Prev)
	paginationdata.NextPage = int32(paginatedData.Pagination.Next)
	if err != nil {
		logs.ErrorLogger.Println("Pagination failure :", err.Error())
		return object, paginationdata, err
	}
	return object, paginationdata, nil
}
