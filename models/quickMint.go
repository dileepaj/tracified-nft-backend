package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TDP struct {
	TenantID         string             `json:"tenantID" bson:"tenantid,omitempty"`
	StageID          string             `json:"stageID" bson:"stagetid,omitempty"`
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
type TDPforSVG struct {
	TenantID         string                 `json:"tenantID" bson:"tenantid,omitempty"`
	StageID          int                    `json:"stageID" bson:"stagetid,omitempty"`
	UserID           string                 `json:"userID" bson:"usertid,omitempty"`
	TimeStamp        primitive.DateTime     `json:"timestamp" bson:"timestamp,omitempty"`
	TraceabilityData map[string]interface{} `json:"traceabilityData" bson:"tracabilityData,omitempty"`
	TracibilityIDs   []string               `json:"tracibilityIDs" bson:"tracibilityIDs,omitempty"`
	WorkFlowRevision string                 `json:"workFlowRevision" bson:"workFlowRevision,omitempty"`
	SignedBy         []SignedBy             `json:"signedBy" bson:"signedBy,omitempty"`
	ArtifactMetaData map[any][]any          `json:"-" bson:"artifactMetaData,omitempty"`
	Id               string                 `json:"id" bson:"id,omitempty"`
	Identifier       string                 `json:"identifier" bson:"identifier,omitempty"`
}
type TraceabilityData struct {
	Type int    `json:"type" bson:"type,omitempty"`
	Val  any    `json:"val" bson:"val,omitempty"`
	Key  string `json:"key" bson:"key,omitempty"`
}

//	type GeoImage struct {
//		Type int                    `json:"type" bson:"type,omitempty"`
//		Val  map[string]interface{} `json:"val" bson:"val,omitempty"`
//		Key  string                 `json:"key" bson:"key,omitempty"`
//	}
type GeoImageData struct {
	Description string             `json:"description" bson:"description,omitempty"`
	GeoCode     GeoCode            `json:"geoCode" bson:"geoCode,omitempty"`
	Image       string             `json:"image" bson:"image,omitempty"`
	TimeStamp   primitive.DateTime `json:"timestamp" bson:"timestamp,omitempty"`
}
type GeoCode struct {
	Lat  float32 `json:"lat" bson:"lat,omitempty"`
	Long float32 `json:"long" bson:"long,omitempty"`
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
	OtpID      primitive.ObjectID `json:"otpid" bson:"_id,omitempty"`
	Email      string             `json:"email" bson:"email,omitempty"`
	Otp        string             `json:"otp" bson:"otp,omitempty"`
	BatchID    string             `json:"batchid" bson:"batchid,omitempty"`
	Validated  string             `json:"validated" bson:"validated,omitempty"`
	ExpireDate primitive.DateTime `json:"expDate" bson:"expDate,omitempty"`
}

type UserNFTMapping struct {
	SvgID   primitive.ObjectID `json:"svgid" bson:"_id,omitempty"`
	BatchID string             `json:"batchid" bson:"batchid,omitempty"`
	Email   string             `json:"email" bson:"email,omitempty"`
	SVG     string             `json:"svg" bson:"svg,omitempty"`
	Hash    string             `json:"hash" bson:"hash,omitempty"`
}

type CollectorInfo struct {
	Photo         string `json:"photo" bson:"photo,omitempty"`
	Name          string `json:"name" bson:"name,omitempty"`
	Address       string `json:"address" bson:"address,omitempty"`
	ContactNumber string `json:"contactNumber" bson:"contactNumber,omitempty"`
}

type CertificationAuthority struct {
	Name    string `json:"name" bson:"name,omitempty"`
	Address string `json:"address" bson:"address,omitempty"`
}

type ExporterInfo struct {
	LicenseExpirationDate string `json:"licenseExpirationDate" bson:"licenseExpirationDate,omitempty"`
	Name                  string `json:"name" bson:"name,omitempty"`
	Address               string `json:"address" bson:"address,omitempty"`
	LicenseNumber         string `json:"licenseNumber" bson:"licenseNumber,omitempty"`
}

type Appraiser struct {
	Name          string `json:"name" bson:"name,omitempty"`
	Qualification string `json:"qualification" bson:"qualification,omitempty"`
}

type TDPParent struct {
	StageID                 string `json:"stageID" bson:"stagetid,omitempty"`
	TraceabilityDataPackets []TDP  `json:"traceabilityDataPackets" bson:"traceabilityDataPackets,omitempty"`
	Id                      string `json:"id" bson:"id,omitempty"`
	Identifier              string `json:"identifier" bson:"identifier,omitempty"`
}

// new ruri nft development
type DigitalTwin struct {
	Name        string      `json:"name" bson:"name,omitempty"`
	Item        string      `json:"item" bson:"item,omitempty"`
	VerticalTab []Component `json:"verticalTab" bson:"verticalTab,omitempty"`
}

type Component struct {
	Title       string       `json:"title" bson:"title,omitempty"`
	Name        string       `json:"name" bson:"name,omitempty"`
	Item        string       `json:"item" bson:"item,omitempty"`
	VerticalTab []Component  `json:"verticalTab" bson:"verticalTab,omitempty"`
	Tabs        []Component  `json:"tabs" bson:"tabs,omitempty"`
	Subtitle    string       `json:"subtitle" bson:"subtitle,omitempty"`
	Component   string       `json:"component" bson:"component,omitempty"`
	Icon        string       `json:"icon" bson:"icon,omitempty"`
	Images      Images       `json:"images" bson:"images,omitempty"`
	Key         string       `json:"key" bson:"key,omitempty"`
	Value       any          `json:"value" bson:"value,omitempty"`
	Coordinates []Coordinate `json:"coordinates" bson:"coordinates,omitempty"`
	Children    []Component  `json:"children" bson:"name,omitempty"`
}

type Coordinate struct {
	Title       string            `json:"title" bson:"title,omitempty"`
	Description string            `json:"description" bson:"description,omitempty"`
	Values      []CoordinateValue `json:"values" bson:"values,omitempty"`
}

type CoordinateValue struct {
	Lat  float64 `json:"lat" bson:"lat,omitempty"`
	Long float64 `json:"long" bson:"long,omitempty"`
}

type ValueWithProof struct {
	Provable bool     `json:"provable" bson:"provable,omitempty"`
	Proofs   []Proof  `json:"proofs" bson:"proofs,omitempty"`
	Value    any      `json:"value" bson:"value,omitempty"`
	UserId   []string `json:"userid" bson:"userid,omitempty"`
	TdpId    []string `json:"tdpid" bson:"tdpid,omitempty"`
}

type Proof struct {
	Name        string `json:"name" bson:"name,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
}

type Images struct {
	/* Provable bool         `json:"provable" bson:"provable,omitempty"`
	Proofs   []Proof      `json:"proofs" bson:"proofs,omitempty"` */
	Value []ImageValue `json:"value" bson:"value,omitempty"`
	/* 	UserId   []string     `json:"userid" bson:"userid,omitempty"`
	   	TdpId    []string     `json:"tdpid" bson:"tdpid,omitempty"` */
}

type ImageValue struct {
	Img       string `json:"img" bson:"img,omitempty"`
	Comment   string `json:"comment" bson:"comment,omitempty"`
	Time      string `json:"time" bson:"time,omitempty"`
	FieldName string `json:"fieldName" bson:"fieldName,omitempty"`

}

type Response struct {
	SVGID  string
	Status string
}
