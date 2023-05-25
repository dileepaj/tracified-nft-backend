package commons

import (
	"os"

	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func GoDotEnvVariable(key string) string {
	loadEnvErr := godotenv.Load(".env")
	if loadEnvErr != nil {
		logs.ErrorLogger.Println("Failed to load ENV : ", loadEnvErr.Error())
	}
	return os.Getenv(key)
}
