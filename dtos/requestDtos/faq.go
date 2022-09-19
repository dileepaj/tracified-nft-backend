package requestDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdateFaq struct {
	QuestionID primitive.ObjectID `json:"questionID" bson:"_id,omitempty"`
	Question   string             `json:"question" bson:"question"`
	Answers    []string           `json:"answers" bson:"answers"`
}

type UpdateUserFAQ struct {
	UserQuestionID primitive.ObjectID `json:"userquestionID" bson:"_id,omitempty"`
	Status         string             `json:"status" bson:"status" `
	Answer         string             `json:"answer" bson:"answer" `
}
