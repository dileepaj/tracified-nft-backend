package responseDtos

type QueryResult struct {
	Result   string `json:"result"`
	OTPType  string `json:"otptype"`
	WidgetId string `json:"widgetid"`
}
