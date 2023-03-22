package commons

import (
	"os"

	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil || os.Getenv("BRANCH_NAME") == "qa" || os.Getenv("BRANCH_NAME") == "staging" {
		logs.InfoLogger.Println("Info Issue with loading .env1 file")
	}
	return os.Getenv(key)
}
