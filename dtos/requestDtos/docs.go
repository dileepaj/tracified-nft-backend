package requestDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdateDoc struct {
	TopicID primitive.ObjectID `json:"topicID" bson:"_id,omitempty"`
	Topic   string             `json:"topic" bson:"topic"`
	Answers []string           `json:"answers" bson:"answers"`
}
