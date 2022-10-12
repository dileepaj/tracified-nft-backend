package responseDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetPendingUserFAQ struct {
	UserQuestionID primitive.ObjectID `json:"userquestionID" bson:"_id,omitempty"`
	UserMail       string             `json:"usermail" bson:"usermail"`
	Category       string             `json:"category" bson:"category"`
	Subject        string             `json:"subject" bson:"subject"`
	Description    string             `json:"desc" bson:"desc"`
	Status         string             `json:"status" bson:"status"`
	Answer         string             `json:"answer" bson:"answer" `
}

type GetAttachmentbyID struct {
	Attachment string `json:"attached" bson:"attached"`
}
