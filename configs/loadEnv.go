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
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		logs.ErrorLogger.Println("Error loading .env1 file")
	}
	backendToken = os.Getenv("BACKEND_TOKEN")
	EnvName = os.Getenv("BRANCH_NAME")
	port = os.Getenv("BE_PORT")
	backendBaseUrl = os.Getenv("BACKEND_BASEURL")
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
	if port != "" {
		return ":" + port
	}
	return ":6080"
}
