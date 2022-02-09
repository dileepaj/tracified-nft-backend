package nftMarketplace

type mintNFTStellar struct {
	Blockchain           string
	NFTTXNhash           string
	createdAr            string
	NftAssetName         string
	NftContentName       string
	NftContent           string
	CuurentIssuerPK      string
	PreviousOwnerPK      string
	MainAccountPK        string
	InitialDistributorPK string
	InitialIssuerPK      string
	NFTcollection        string
	Category             string
	UerId                string
	Imagebase64          string
	SellingStatus        string
	SellingType          string
	Price                string
}

type mintNFTPloygonAndEth struct {
	Blockchain      string
	NFTTXNhash      string
	createdAr       string
	NftAssetName    string
	NftContentName  string
	NftContent      string
	CuurentOwnerPK  string
	PreviousOwnerPK string
	MainAccountPK   string
	InitialOwner    string
	NFTcollection   string
	Category        string
	UerId           string
	Imagebase64     string
	SellingStatus   string
	SellingType     string
	Price           string
	SmartContract   string
	TokenType       string
	mintedContract  string
}