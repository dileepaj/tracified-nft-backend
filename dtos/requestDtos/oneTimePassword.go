package requestDtos

type GenOTP struct {
	Email     string `json:"email" bson:"email,omitempty"`
	ProductID string `json:"productID" bson:"productID,omitempty"`
}
type ValidateOTP struct {
	Email   string `json:"email" bson:"email,omitempty"`
	OTPCode string `json:"otp" bson:"otp,omitempty"`
}

type UpdateOTPStatus struct {
	Email     string `json:"email" bson:"email"`
	OTP       string `json:"otp" bson:"otp"`
	NFTStatus string `json:"nftstatus" bson:"nftstatus"`
}
