package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	Timestamp  primitive.DateTime `json:"timestamp" bson:"timestamp"`
	Blockchain string             `json:"blockchain" bson:"blockchain"`
	Address    string             `json:"Address" bson:"Address"`
}
type User struct {
	UserId     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email      string             `json:"email" bson:"email" validate:"required,email"`
	TenentName string             `json:"tenentname" bson:"tenentname" validate:"required"`
	BCAccounts []Account			`json:"bcaccounts" bson:"bcaccounts"`
	Timestamp  primitive.DateTime `json:"timestamp" bson:"timestamp"`
}
