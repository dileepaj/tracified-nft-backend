package models

type Endorse struct {
	Name        string `json:"name" bson:"name"`
	PublicKey   string `json:"publickey" bson:"publickey"`
	Email       string `json:"email" bson:"email"`
	Contact     string `json:"contact" bson:"contact"`
	Description string `json:"desc" bson:"desc"`
}
