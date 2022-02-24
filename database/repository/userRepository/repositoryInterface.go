package userRepository

import "go.mongodb.org/mongo-driver/mongo"

type connection struct {
	Connection *mongo.Database
}

func (dbconnection *connection) FindById() {}

func FindByID() {}

func Create() {}

func Update() {}

func UpdateOne() {}