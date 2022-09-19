package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Structure for newsletter model class and collection
type NewsLetter struct {
	NewsId      primitive.ObjectID `json:"newsid" bson:"_id,omitempty"`
	Topic       string             `json:"topic" bson:"topic"`
	Author      string             `json:"author" bson:"author"`
	Date        string             `json:"date" bson:"date"`
	Publisher   string             `json:"publisher" bson:"publisher"`
	WebLink     string             `json:"weblink" bson:"weblink"`
	Description string             `json:"description" bson:"descriptipon"`
	Image       string             `json:"image" bson:"iamge"`
}

type Subscription struct {
	SubscriptionID primitive.ObjectID `json:"subscriptionID" bson:"_id,omitempty"`
	UserMail       string             `json:"mail" bson:"mail"`
}
