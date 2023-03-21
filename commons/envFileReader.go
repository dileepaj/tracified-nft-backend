package commons

import (
	"os"
)

// use godot package to load/read the .env file and
// return the value of the key
func GoDotEnvVariable(key string) string {
	return os.Getenv(key)
}
