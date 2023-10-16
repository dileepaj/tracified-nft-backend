package ipfsbusinessfacade

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services/ipfsservice"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

var (
	fileBaseBucket = os.Getenv("FILEBASE_BUCKET")
)

func UploadFilesToIpfs(fileObj models.IpfsObject) (string, error) {
	//check the file type
	//1 - Image, 2 - TDP_Image, 3 - TDP
	cidHash := ""
	if fileObj.FileType == 1 || fileObj.FileType == 2 {
		//write image to a file
		var imageStrAarray = strings.Split(fileObj.FileContent, ";base64,")
		dec, errWhenDecodingString := base64.StdEncoding.DecodeString(imageStrAarray[1])
		extentionType := strings.Split(imageStrAarray[0], "data:image/")[1]
		if errWhenDecodingString != nil {
			logs.ErrorLogger.Println("Error when decoding image data : ", errWhenDecodingString.Error())
			return "", errWhenDecodingString
		}

		imgName := fileObj.FileName + "." + extentionType
		imgName = strings.ToLower(imgName)

		//get current working directory
		workingDirectory, errWhenGettingTheDirectory := os.Getwd()
		if errWhenGettingTheDirectory != nil {
			logs.ErrorLogger.Println("Error when getting the working directory : ", errWhenGettingTheDirectory.Error())
			return "", errWhenGettingTheDirectory
		}

		//create file
		filePath := filepath.Join(workingDirectory, imgName)
		file, errWhenCreatingFile := os.Create(filePath)
		if errWhenCreatingFile != nil {
			logs.ErrorLogger.Println("Error when creating file : ", errWhenCreatingFile.Error())
			return "", errWhenCreatingFile
		}
		defer file.Close()

		//write into created file
		if _, errWhenWritingToFile := file.Write(dec); errWhenWritingToFile != nil {
			logs.ErrorLogger.Println("Error when writing data into file : ", errWhenWritingToFile.Error())
			return "", errWhenWritingToFile
		}
		if errWhenSyncing := file.Sync(); errWhenSyncing != nil {
			logs.ErrorLogger.Println("Error when syncing and clearing memory : ", errWhenSyncing.Error())
			return "", errWhenSyncing
		}

		folderName := ""
		//folder name based on type
		if fileObj.FileType == 1 {
			//upload image to image bucket
			folderName = "Image"
		} else {
			//upload image to TDP image bucket
			folderName = "TDP/Image"
		}

		cid, _, errWhenUploadingToIpfs := ipfsservice.UploadFile(filePath, fileObj.FileName, fileBaseBucket, folderName)
		if errWhenUploadingToIpfs != nil {
			logs.ErrorLogger.Println("Error when uploading to IPFS")
			//delete image
			errWhenRemovingImage := os.Remove(filePath)
			if errWhenRemovingImage != nil {
				logs.ErrorLogger.Println("Error when removing the file : ", errWhenRemovingImage)
				return "", errWhenRemovingImage
			}
			return "", errWhenUploadingToIpfs
		}

		errWhenClosingFile := file.Close()
		if errWhenClosingFile != nil {
			logs.ErrorLogger.Println("Error when closing the file : ", errWhenClosingFile)
			return "", errWhenClosingFile
		}
		cidHash = cid

		//delete the image
		errWhenRemovingImage := os.Remove(filePath)
		if errWhenRemovingImage != nil {
			logs.ErrorLogger.Println("Error when removing the image file : ", errWhenRemovingImage)
			return "", errWhenRemovingImage
		}

	} else if fileObj.FileType == 3 {
		//write into json file
		jsonContent := []byte(fileObj.FileContent)
		jsonFileName := fileObj.FileName + ".json"
		jsonFileName = strings.ToLower(jsonFileName)

		// Get the current working directory
		workingDirectory, errWhenGettingTheDirectory := os.Getwd()
		if errWhenGettingTheDirectory != nil {
			logs.ErrorLogger.Println("Error when getting the working directory: ", errWhenGettingTheDirectory.Error())
			return "", errWhenGettingTheDirectory
		}

		// Create the JSON file
		jsonFilePath := filepath.Join(workingDirectory, jsonFileName)
		jsonFile, errWhenCreatingJSONFile := os.Create(jsonFilePath)
		if errWhenCreatingJSONFile != nil {
			logs.ErrorLogger.Println("Error when creating JSON file: ", errWhenCreatingJSONFile.Error())
			return "", errWhenCreatingJSONFile
		}
		defer jsonFile.Close()

		// Write JSON content to the file
		_, errWhenWritingJSONToFile := jsonFile.Write(jsonContent)
		if errWhenWritingJSONToFile != nil {
			logs.ErrorLogger.Println("Error when writing JSON data to file: ", errWhenWritingJSONToFile.Error())
			return "", errWhenWritingJSONToFile
		}
		if errWhenSyncing := jsonFile.Sync(); errWhenSyncing != nil {
			logs.ErrorLogger.Println("Error when syncing and clearing memory: ", errWhenSyncing.Error())
			return "", errWhenSyncing
		}

		folderName := "TDP"

		cid, _, errWhenUploadingToIpfs := ipfsservice.UploadFile(jsonFilePath, fileObj.FileName, fileBaseBucket, folderName)
		if errWhenUploadingToIpfs != nil {
			logs.ErrorLogger.Println("Error when uploading to IPFS")
			//delete image
			errWhenRemovingJsonFile := os.Remove(jsonFilePath)
			if errWhenRemovingJsonFile != nil {
				logs.ErrorLogger.Println("Error when removing the file : ", errWhenRemovingJsonFile)
				return "", errWhenRemovingJsonFile
			}
			return "", errWhenUploadingToIpfs
		}

		errWhenClosingFile := jsonFile.Close()
		if errWhenClosingFile != nil {
			logs.ErrorLogger.Println("Error when closing the file : ", errWhenClosingFile)
			return "", errWhenClosingFile
		}
		cidHash = cid

		errWhenRemovingJson := os.Remove(jsonFilePath)
		if errWhenRemovingJson != nil {
			logs.ErrorLogger.Println("Error when removing the Json file : ", errWhenRemovingJson)
			return "", errWhenRemovingJson
		}
	}

	//Add the content details to DB
	insertObj := models.IpfsInsertObject{
		FileType: fileObj.FileType,
		FileName: fileObj.FileName,
		TdpId:    fileObj.TdpId,
		Cid:      cidHash,
	}

	_, errWhenSavingDetails := IpfsRepository.SaveFileDetails(insertObj)
	if errWhenSavingDetails != nil {
		logs.ErrorLogger.Println("Error when saving file details on collection ; ", errWhenSavingDetails)
		return "", errWhenSavingDetails
	}

	logs.InfoLogger.Println("File uploaded to IPFS at : https://ipfs.filebase.io/ipfs/" + cidHash)

	return cidHash, nil
}
