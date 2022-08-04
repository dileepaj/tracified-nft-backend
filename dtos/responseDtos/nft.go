package responseDtos

type ResponseNFTMakeSale struct {
	NFTIdentifier  string   `json:"nftidentifier" bson:"nftidentifier" validate:"required"`
	Blockchain     string   `json:"blockchain" bson:"blockchain" validate:"required"`
	NFTName        string   `json:"nftname" bson:"nftname" validate:"required"`
	NftContentURL  string   `json:"nftcontenturl" bson:"nftcontenturl" validate:"required,url"`
	Timestamp      string   `json:"timestamp" bson:"timestamp"`
	Collection     string   `json:"collection" bson:"collection" validate:"required"`
	Category       string   `json:"category" bson:"category" validate:"required"`
	Tags           []string `json:"tags" bson:"tags" validate:"required"`
	Imagebase64    string   `json:"imagebase64" bson:"nimagebase64" validate:"required,base64"`
	CurrentPrice   string   `json:"currentprice" bson:"currentprice" `
	CreatorPK      string   `json:"creatorpk" bson:"creatorpk" validate:"required"`
	CurrentOwnerPK string   `json:"currentownerpk" bson:"currentownerpk" validate:"required"`
	SellingStatus  string   `json:"sellingstatus" bson:"sellingstatus" validate:"required"` //ONSALE,NOTSALE,NOTLISTED
	SellingType    string   `json:"sellingtype" bson:"sellingtype" `                        //NOTLISTED
	DistributorPK  string   `json:"distributorpk" bson:"distributorpk"`
	Smartcontract  string   `json:"smartcontract" bson:"smartcontract"`
	MintedContract string   `json:"mintsmartcontract" bson:"mintsmartcontract"`
	TokenType      string   `json:"tokentype" bson:"tokentype"`
}

type ResponseNFTMinter struct {
	Imagebase64   string `json:"imagebase64" bson:"imagebase64" `
	NFTIssuerPK   string `json:"nftissuerpk" bson:"nftissuerpk" `
	Blockchain    string `json:"blockchain" bson:"blockchain" `
	NFTName       string `json:"nftname" bson:"nftname"`
	NftContentURL string `json:"nftcontenturl" bson:"nftcontenturl"`
	NFTIdentifier string `json:"nftidentifier" bson:"nftidentifier" `
	NFTTxnHash    string `json:"nfttxnhash" bson:"nfttxnhash" `
}

type ResponseNFTMintTXN struct {
	Imagebase64   string `json:"imagebase64" bson:"imagebase64" `
	NFTIssuerPK   string `json:"nftissuerpk" bson:"nftissuerpk" `
	Blockchain    string `json:"blockchain" bson:"blockchain" `
	NFTName       string `json:"nftname" bson:"nftname"`
	NftContentURL string `json:"nftcontenturl" bson:"nftcontenturl"`
	NFTTxnHash    string `json:"nfttxnhash" bson:"nfttxnhash" `
}
