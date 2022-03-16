package responseDtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResponseProject struct {
	ProjectId   string
	ProjectName string
	Timestamp   primitive.DateTime
}
