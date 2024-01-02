package configs

import (
	"os"
	"strings"

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
	tillit         = ""
	nftMarketplace = ""
	streamingAPI   = ""
	timelineANDJourneyMapGeneratorAPI = ""
	tenantList     = ""
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
	tillit = os.Getenv("TILLIT")
	nftMarketplace = os.Getenv("NFT_MARKETPLACE")
	streamingAPI = os.Getenv("STREAMING_API")
	timelineANDJourneyMapGeneratorAPI = os.Getenv("TIMELINE_GENERATION_API")
	tenantList = os.Getenv("TENANT_LIST")
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

func GetTillitUrl() string {
	LoadEnv()
	return tillit
}

func GetMarketplaceUrl() string {
	LoadEnv()
	return nftMarketplace
}

func GetStreamingAPIUrl() string {
	LoadEnv()
	return streamingAPI
}

func GetTimelineANDJourneyMapGeneratorAPI() string {
	LoadEnv()
	return timelineANDJourneyMapGeneratorAPI
}

func GetTenantList() []string {
	LoadEnv()
	return strings.Split(tenantList, ",")
}