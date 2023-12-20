package composerimgservice

import (
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/constants"
	"github.com/dileepaj/tracified-nft-backend/services/ipfsservice"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

var (
	fileBaseBucket = os.Getenv("FILEBASE_BUCKET")
)

func UploadImageToIpfsWithFolder(widgetType int, imageContent string, projectId string, widgetId string, tenetId string, imageTitle string) (string, error) {
	var childFolderName string

	//set up folder path
	switch widgetType {
	case constants.ImageWidget:
		childFolderName = "ImageWidget_" + widgetId
	case constants.TimelineWidget:
		childFolderName = "TimelineWidget_" + widgetId
	default:
		return "", errors.New("Invalid widget type")
	}

	folderPath := "Marketplace/" + tenetId + "/" + "NFT" + "/" + projectId + "/" + childFolderName

	errWhenCreatingFolder := ipfsservice.CreateFolder(fileBaseBucket, folderPath)
	if errWhenCreatingFolder != nil {
		return "", errWhenCreatingFolder
	}

	contentStrArr := strings.Split(imageContent, ";base64,")
	decoded, errWhenDecodingString := base64.StdEncoding.DecodeString(contentStrArr[1])
	fileType := strings.Split(contentStrArr[0], "data:image/")[1]
	if errWhenDecodingString != nil {
		logs.ErrorLogger.Println("Error when decoding data : ", errWhenDecodingString.Error())
		return "", errWhenDecodingString
	}

	fileName := imageTitle + "." + fileType
	fileName = strings.ToLower(fileName)

	directory, errWhenGettingDirectory := os.Getwd()
	if errWhenGettingDirectory != nil {
		logs.ErrorLogger.Println("Error when getting the working directory : ", errWhenGettingDirectory.Error())
		return "", errWhenGettingDirectory
	}

	filePath := filepath.Join(directory, fileName)
	f, errWhenCreatingFile := os.Create(filePath)
	if errWhenCreatingFile != nil {
		logs.ErrorLogger.Println("Error when writing data to file : ", errWhenCreatingFile.Error())
		return "", errWhenCreatingFile
	}

	defer f.Close()

	if _, errWhenWritingToFile := f.Write(decoded); errWhenWritingToFile != nil {
		logs.ErrorLogger.Println("Error when writing into the file : ", errWhenWritingToFile.Error())
		return "", errWhenWritingToFile
	}

	if errWhenSyncing := f.Sync(); errWhenSyncing != nil {
		logs.ErrorLogger.Println("Error when syncing : ", errWhenSyncing.Error())
		return "", errWhenSyncing
	}

	defer f.Close()

	cid, _, errWhenUploadingToIpfs := ipfsservice.UploadFile(filePath, fileName, fileBaseBucket, folderPath)
	if errWhenUploadingToIpfs != nil {
		logs.ErrorLogger.Println("Error when uploading to IPFS : ", errWhenUploadingToIpfs.Error())

		errWhenRemovingFile := os.Remove(filePath)
		if errWhenRemovingFile != nil {
			logs.ErrorLogger.Println("Error when removing the file : ", errWhenRemovingFile.Error())
			return "", errWhenRemovingFile
		}
		return "", errWhenRemovingFile
	}

	errWhenClosingFile := f.Close()
	if errWhenClosingFile != nil {
		logs.ErrorLogger.Println("Error when closing file : ", errWhenClosingFile.Error())
		return "", errWhenClosingFile
	}

	errWhenRemovingFile := os.Remove(filePath)
	if errWhenRemovingFile != nil {
		logs.ErrorLogger.Println("Error when removing file : ", errWhenRemovingFile.Error())
		return "", errWhenRemovingFile
	}

	return cid, nil
}
