package responseDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type QueryResult struct {
	Result   string             `json:"result"`
	OTPType  string             `json:"otptype"`
	WidgetId primitive.ObjectID `json:"widgetid"`
}
