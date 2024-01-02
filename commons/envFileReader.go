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

// ContainsString checks if a given string is present in the array.
func ContainsString(array []string, target string) bool {
	for _, value := range array {
		if value == target {
			return true
		}
	}
	return false
}