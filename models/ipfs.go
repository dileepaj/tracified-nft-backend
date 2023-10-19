package models

type IpfsObjectForTDP struct {
	TDPDetails  TDPDetails
	FileDetails FileDetails
}

type TDPDetails struct {
	TenetID string `json:"tenetid" bson:"tenetid"`
	ItemID  string `json:"itemid" bson:"itemid"`
	BatchID string `json:"batchid" bson:"batchid"`
	TdpID   string `json:"tdpid" bson:"tdpid"`
}
type FileDetails struct {
	FileType    int    `json:"filetype" bson:"filetype"`
	FileContent string `json:"filecontent" bson:"filecontent"`
	FileName    string `json:"filename" bson:"filename"`
}

type InsertTdpDetails struct {
	TenetId string
	Items   []ItemDetails
}

type ItemDetails struct {
	ItemId  string
	Batches []BatchDetails
}

type BatchDetails struct {
	BatchId string
	Tdps    []TDPObject
}

type TDPObject struct {
	TdpId  string
	TdpCid string
	Images []ImageObject
}

type ImageObject struct {
	ImageName string
	ImageCid  string
}

type IpfsResponse struct {
	Message string
	Url     string
}
