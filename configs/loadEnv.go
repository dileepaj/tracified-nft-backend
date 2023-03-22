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
)

func LoadEnv() {
	godotenv.Load(".env")
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
	LoadEnv()
	if port != "" {
		return ":" + port
	}
	return ":6080"
}
