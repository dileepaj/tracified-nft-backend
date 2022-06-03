package requestDtos

type UpdateEndorsementByPublicKey struct {
	PublicKey string `json:"publickey" bson:"publickey"`
	Status    string `json:"status" bson:"status" `
}
