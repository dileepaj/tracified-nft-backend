package commons

import (
	"os"

	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func GoDotEnvVariable(key string) string {
	godotenv.Load(".env")
	return os.Getenv(key)
}
