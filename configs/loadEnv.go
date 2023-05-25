package configs

import (
	"os"

	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/joho/godotenv"
)

var (
	backendToken   = ""
	port           = ""
	EnvName        = ""
	backendBaseUrl = ""
	ruriShopify    = ""
	digitalTwin    = ""
	nftBackendUrl  = ""
	adminBackend   = ""
	gateway        = ""
)

func LoadEnv() {
	loaderr := godotenv.Load(".env")
	if loaderr != nil {
		logs.ErrorLogger.Println("Failed to load env file : ", loaderr.Error())
		return
	}
	backendToken = os.Getenv("BACKEND_TOKEN")
	EnvName = os.Getenv("BRANCH_NAME")
	port = os.Getenv("BE_PORT")
	backendBaseUrl = os.Getenv("BACKEND_BASEURL")
	nftBackendUrl = os.Getenv("NFT_BACKEND_BASEURL")
	ruriShopify = os.Getenv("RURI_SHOPIFY")
	digitalTwin = os.Getenv("DIGITALTWIN")
	adminBackend = os.Getenv("ADMIN_BACKEND")
	gateway = os.Getenv("GATEWAY")
}

func GetBackenToken() string {
	LoadEnv()
	return backendToken
}

func GetBackeBaseUrl() string {
	LoadEnv()
	return backendBaseUrl
}

func GetPort() string {
	LoadEnv()
	if port != "" {
		return ":" + port
	}
	return ":6080"
}

func GetRuriShopifyUrl() string {
	LoadEnv()
	return ruriShopify
}

func GetDigitalTwinUrl() string {
	LoadEnv()
	return digitalTwin
}

func GetNftBackendBaseUrl() string {
	LoadEnv()
	return nftBackendUrl
}

func GetAdminBackendUrl() string {
	LoadEnv()
	return adminBackend
}

func GetGatewayUrl() string {
	LoadEnv()
	return gateway
}
