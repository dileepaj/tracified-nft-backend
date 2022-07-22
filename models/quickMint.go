package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TDP struct {
	TenantID         string             `json:"tenantID" bson:"tenantid,omitempty"`
	StageID          int                `json:"stageID" bson:"stagetid,omitempty"`
	UserID           string             `json:"userID" bson:"usertid,omitempty"`
	TimeStamp        primitive.DateTime `json:"timestamp" bson:"timestamp,omitempty"`
	TraceabilityData []TraceabilityData `json:"traceabilityData" bson:"tracabilityData,omitempty"`
	TracibilityIDs   []string           `json:"tracibilityIDs" bson:"tracibilityIDs,omitempty"`
	WorkFlowRevision string             `json:"workFlowRevision" bson:"workFlowRevision,omitempty"`
	SignedBy         []SignedBy         `json:"signedBy" bson:"signedBy,omitempty"`
	ArtifactMetaData map[any][]any      `json:"-" bson:"artifactMetaData,omitempty"`
	Id               string             `json:"id" bson:"id,omitempty"`
	Identifier       string             `json:"identifier" bson:"identifier,omitempty"`
}

type TraceabilityData struct {
	Type int    `json:"type" bson:"type,omitempty"`
	Val  any    `json:"val" bson:"val,omitempty"`
	Key  string `json:"key" bson:"key,omitempty"`
}

type SignedBy struct {
	PublicKey string `json:"publicKey" bson:"type,omitempty"`
	Role      string `json:"role" bson:"role,omitempty"`
}

type ItemData struct {
	ItemID         string `json:"itemid" bson:"itemid,omitempty"`
	HasTracability string `json:"hastracability" bson:"hastracability,omitempty"`
	BatchID        string `json:"batchid" bson:"batchid,omitempty"`
}

type UserAuth struct {
	OtpID   primitive.ObjectID `json:"otpid" bson:"_id,omitempty"`
	Email   string             `json:"email" bson:"email,omitempty"`
	Otp     string             `json:"otp" bson:"otp,omitempty"`
	BatchID string             `json:"batchid" bson:"batchid,omitempty"`
}

type UserNFTMapping struct {
	BatchID string `json:"batchid" bson:"batchid,omitempty"`
	Email   string `json:"email" bson:"email,omitempty"`
	SVG     string `json:"svg" bson:"svg,omitempty"`
}
