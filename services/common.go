package services

import (
	"fmt"
	"os"

	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

func WriteToFile(location string, fileName string, content string) {
	os.MkdirAll(location, 0755)
	file, err := os.OpenFile(
		location+"/"+fileName,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	file.WriteString(content)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
	defer file.Close()

}

func ReadFromFile(location string) string {
	fmt.Println(location)
	content, err := os.ReadFile(location)
	if err != nil {
		fmt.Println(err)
		logs.ErrorLogger.Println(err.Error())
	}

	return string(content)
}

func MakeDirectory(path string) {

	err := os.MkdirAll(path, 0755)
	if err != nil {
		logs.ErrorLogger.Println(err)
	}
}

func AppendToFile(location string, content string) {

	defFile, err := os.OpenFile(location,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		logs.ErrorLogger.Println(err)
	}

	defer defFile.Close()

	if _, err := defFile.WriteString(content); err != nil {
		logs.ErrorLogger.Println(err)
	}
}
