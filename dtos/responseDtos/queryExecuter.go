package responseDtos

type QueryResult struct {
	Result   string `json:"result"`
	OTPType  string `json:"otptype"`
	WidgetId string `json:"widgetid"`
}

type OTP struct {
	Url string `json:"Url" bson:"url"`
}

type OTPContent struct {
	OtpContent string `json:"OtpConetent" bson:"otpcontent"`
}