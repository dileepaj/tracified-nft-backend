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

// 1- TDP_Image, 2 - TDP
func UploadFilesToIpfs(fileObj models.IpfsObjectForTDP) (string, error) {
	//TODO - Check if the TDP details are already entered

	//get the DB object for tenetID
	//If if is there check for ItemID
	//if it is there check for Batch ID
	//if it is there check the TDP id
	//for TDP upload this will not upload the TDP
	//for image reupload the file

	// var itemIndex int
	// resultTdpDetails, errWhenGettingTdpDetails := IpfsRepository.GetTdpDetails(fileObj.TDPDetails.TenetID)
	// if errWhenGettingTdpDetails != nil {
	// 	return "", errWhenGettingTdpDetails
	// } else if resultTdpDetails.TenetId == "" {
	// 	//Insert new object
	// } else {
	// 	//Check the Item ID in the item loop
	// 	for i := 0; i < len(resultTdpDetails.Items); i++ {
	// 		if resultTdpDetails.Items[i].ItemId == fileObj.TDPDetails.ItemID {
	// 			//get the index of the item ID
	// 			itemIndex = i
	// 		}
	// 	}
	// 	//TODO add the item details to the DB

	// }

	//check the file type
	cidHash := ""
	if fileObj.FileDetails.FileType == 2 {
		//write image to a file
		var imageStrAarray = strings.Split(fileObj.FileDetails.FileContent, ";base64,")
		dec, errWhenDecodingString := base64.StdEncoding.DecodeString(imageStrAarray[1])
		extentionType := strings.Split(imageStrAarray[0], "data:image/")[1]
		if errWhenDecodingString != nil {
			logs.ErrorLogger.Println("Error when decoding image data : ", errWhenDecodingString.Error())
			return "", errWhenDecodingString
		}

		imgName := fileObj.FileDetails.FileName + "." + extentionType
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
		folderName := "TDP/Image"

		cid, _, errWhenUploadingToIpfs := ipfsservice.UploadFile(filePath, fileObj.FileDetails.FileName, fileBaseBucket, folderName)
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

	} else if fileObj.FileDetails.FileType == 2 {
		//write into json file
		jsonContent := []byte(fileObj.FileDetails.FileContent)
		jsonFileName := fileObj.FileDetails.FileName + ".json"
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

		cid, _, errWhenUploadingToIpfs := ipfsservice.UploadFile(jsonFilePath, fileObj.FileDetails.FileName, fileBaseBucket, folderName)
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
	insertObj := models.TracifiedDataPacket{}

	_, errWhenSavingDetails := IpfsRepository.SaveFileDetails(insertObj)
	if errWhenSavingDetails != nil {
		logs.ErrorLogger.Println("Error when saving file details on collection ; ", errWhenSavingDetails)
		return "", errWhenSavingDetails
	}

	logs.InfoLogger.Println("CID Hash : " + cidHash)

	return cidHash, nil
}
