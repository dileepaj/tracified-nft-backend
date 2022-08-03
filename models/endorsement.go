package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Endorse struct {
	Id          primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	Name        string             `json:"Name" bson:"name"`
	PublicKey   string             `json:"PublicKey" bson:"publickey"`
	Email       string             `json:"Email" bson:"email"`
	Contact     string             `json:"Contact" bson:"contact"`
	Description string             `json:"Description" bson:"description"`
	Blockchain  string             `json:"Blockchain" bson:"blockchain"`
	Status      string             `json:"Status" bson:"status"`
}
