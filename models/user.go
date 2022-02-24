package models

type User struct {
	UserId    string `json:"userid" bson:"userid" validate:"required,string"`
	Email     string `json:"email" bson:"email" validate:"required,email"`
	Accounts  string `json:"accounts" bson:"accounts" validate:"required"`
	Company   string `json:"company" bson:"companay" validate:"required"`
	CreatedAt string `json:"createdat" bson:"createdat" validate:"required,timestamp"`
}