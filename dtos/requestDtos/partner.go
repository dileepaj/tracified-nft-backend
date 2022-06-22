package requestDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdatePartner struct {
	ID          primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Topic       string             `json:"topic" bson:"topic"`
	WebLink     string             `json:"weblink" bson:"weblink"`
	CompanyName string             `json:"companyname" bson:"companyname"`
	Image       string             `json:"image" bson:"image"`
	Description string             `json:"desc" bson:"desc"`
}
