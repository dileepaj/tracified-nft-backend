package models

type IpfsObject struct {
	FileType    int    `json:"filetype" bson:"filetype"`
	FileContent string `json:"filecontent" bson:"filecontent"`
	FileName    string `json:"filename" bson:"filename"`
	TdpId       string `json:"tdpid" bson:"tdpid"`
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
