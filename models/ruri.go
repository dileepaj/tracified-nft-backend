package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TDP struct {
	TenantID         string             `json:"tenantid" bson:"tenantid,omitempty"`
	StageID          int                `json:"stagetid" bson:"stagetid,omitempty"`
	UserID           string             `json:"usertid" bson:"usertid,omitempty"`
	TimeStamp        primitive.DateTime `json:"timestamp" bson:"timestamp,omitempty"`
	TracabilityData  []TracibilityData  `json:"tracabilityData" bson:"tracabilityData,omitempty"`
	TracibilityIDs   []string           `json:"tracibilityIDs" bson:"tracibilityIDs,omitempty"`
	WorkFlowRevision string             `json:"workFlowRevision" bson:"workFlowRevision,omitempty"`
	SignedBy         []SignedBy         `json:"signedBy" bson:"signedBy,omitempty"`
	ArtifactMetaData string             `json:"artifactMetaData" bson:"artifactMetaData,omitempty"`
	Id               string             `json:"id" bson:"id,omitempty"`
}

type TracibilityData struct {
	Type int    `json:"type" bson:"type,omitempty"`
	Val  any    `json:"val" bson:"val,omitempty"`
	Key  string `json:"key" bson:"key,omitempty"`
}

type SignedBy struct {
	PublicKey string `json:"type" bson:"type,omitempty"`
	Role      string `json:"role" bson:"role,omitempty"`
}

type RuriItemData struct {
	ItemID         string `json:"itemid" bson:"itemid,omitempty"`
	HasTracability string `json:"hastracability" bson:"hastracability,omitempty"`
	BatchID        string `json:"batchid" bson:"batchid,omitempty"`
}

type OTPData struct {
	OtpID   primitive.ObjectID `json:"otpid" bson:"_id,omitempty"`
	Email   string             `json:"email" bson:"email,omitempty"`
	Otp     string             `json:"otp" bson:"otp,omitempty"`
	BatchID string             `json:"batchid" bson:"batchid,omitempty"`
}
