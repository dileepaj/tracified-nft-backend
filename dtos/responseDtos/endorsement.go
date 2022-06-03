package responseDtos

type ResponseEndorsementUpdate struct {
	Name      string `json:"name" bson:"name" `
	PublicKey string `json:"publickey" bson:"publickey"`
	Status    string `json:"status" bson:"status" `
}
