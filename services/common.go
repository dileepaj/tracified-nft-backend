package services

import (
	"os"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

//Read the file from a specific location
func ReadFromFile(location string) string {
	content, err := os.ReadFile(location)
	if err != nil {
		logs.ErrorLogger.Println(err.Error())
	}
	readLine := strings.TrimSuffix(string(content), "\r\n")
	return readLine
}

//make the file directory
func MakeDirectory(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
}
