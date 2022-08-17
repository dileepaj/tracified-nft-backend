package requestDtos

type UpdateEndorsementByPublicKey struct {
	PublicKey string `json:"publickey" bson:"publickey"`
	Status    string `json:"status" bson:"status" `
	Review    string `json:"Review" bson:"review"`
	Rating    string `json:"Rating" bson:"rating"`
}

type UpdateEndorsement struct {
	PublicKey string `json:"publickey" bson:"publickey"`
	Name      string `json:"name" bson:"name" `
	Email     string `json:"email" bson:"email" `
	Contact   string `json:"contact" bson:"contact" `
}

type UpdateEndorsement struct {
	PublicKey   string `json:"publickey" bson:"publickey"`
	Status      string `json:"status" bson:"status" `
	Name        string `json:"name" bson:"name" `
	Email       string `json:"email" bson:"email" `
	Contact     string `json:"contact" bson:"contact" `
	Description string `json:"description" bson:"description" `
}
