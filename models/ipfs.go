package models

type IpfsObjectForTDP struct {
	TDPDetails  TDPDetails  `json:"tdpdetails" bson:"tdpdetails"`
	FileDetails FileDetails `json:"filedetails" bson:"filedetails"`
	FileType    int         `json:"filetype" bson:"filetype"`
}

type TDPDetails struct {
	TenetID string `json:"tenetid" bson:"tenetid"`
	ItemID  string `json:"itemid" bson:"itemid"`
	BatchID string `json:"batchid" bson:"batchid"`
	TdpID   string `json:"tdpid" bson:"tdpid"`
}
type FileDetails struct {
	FileContent string `json:"filecontent" bson:"filecontent"`
	FileName    string `json:"filename" bson:"filename"`
}
type TracifiedDataPacket struct {
	TenetId string
	ItemId  string
	BatchId string
	TdpId   string
	TdpCid  string
	Images  []ImageObject
}

type ImageObject struct {
	ImageName string
	ImageCid  string
}

type IpfsResponse struct {
	Message string
	Url     string
}

type IpfsObjectForCollections struct {
	CollectionDetails NFTCollection `json:"collectiondetails" bson:"collectiondetails"`
	FileDetails       FileDetails   `json:"filedetails" bson:"filedetails"`
	FileType          int           `json:"filetype" bson:"filetype"`
}
