package responseDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type QueryResult struct {
	Result   string             `json:"result"`
	OTPType  string             `json:"otptype"`
	WeigetId primitive.ObjectID `json:"weigetid"`
}
