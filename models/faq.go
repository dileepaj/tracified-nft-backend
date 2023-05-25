package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Faq represents frequently asked questions.
type Faq struct {
	QuestionID primitive.ObjectID `json:"questionID" bson:"_id,omitempty"`
	Question   string             `json:"question" bson:"question"`
	Answers    []string           `json:"answers" bson:"answers"`
}

// UserQuestions represents user-submitted questions.
type UserQuestions struct {
	UserQuestionID primitive.ObjectID `json:"userquestionID" bson:"_id,omitempty"`
	UserMail       string             `json:"usermail" bson:"usermail"`
	Category       string             `json:"category" bson:"category"`
	Subject        string             `json:"subject" bson:"subject"`
	Description    string             `json:"desc" bson:"desc"`
	Attachment     string             `json:"attached" bson:"attached"`
	Status         string             `json:"status" bson:"status"`
	Answer         string             `json:"answer" bson:"answer" `
}
