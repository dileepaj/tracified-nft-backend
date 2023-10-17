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

type IpfsInsertObject struct {
	FileType int    `json:"filetype" bson:"filetype"`
	FileName string `json:"filename" bson:"filename"`
	TdpId    string `json:"tdpid" bson:"tdpid"`
	Cid      string `json:"cid" bson:"cid"`
}

type IpfsResponse struct {
	Message string
	Url     string
}
