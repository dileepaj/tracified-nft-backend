package configs

import (
	"os"

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
	godotenv.Load(".env")
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
