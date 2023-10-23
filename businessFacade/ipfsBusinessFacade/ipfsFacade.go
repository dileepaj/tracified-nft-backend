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
	tenetTdpDetails, errWhenGettingTdpDetails := IpfsRepository.GetTdpDetails("tenetid", fileObj.TDPDetails.TenetID)
	if errWhenGettingTdpDetails != nil {
		return "", errWhenGettingTdpDetails
	} else if tenetTdpDetails.TenetId == "" {
		//New record should be added with the relevant tenet details
		//Create new folder for tenet
		//Create new folder for item
		//Create new folder for batch
		//Create new folder for TDP id
	} else {
		//check if the item is already recoded
		itemTdpDetails, errWhenGettingItemDetails := IpfsRepository.GetTdpDetails("itemid", fileObj.TDPDetails.ItemID)
		if errWhenGettingItemDetails != nil {
			return "", errWhenGettingItemDetails
		} else if itemTdpDetails.ItemId == "" {
			//New record should be added with the relevant item details
			//Create new folder for item
			//Create new folder for batch
			//Create new folder for TDP id
		} else {
			//check if the batch is already added
			batchDetails, errWhenGettingBatchDetails := IpfsRepository.GetTdpDetails("batchid", fileObj.TDPDetails.BatchID)
			if errWhenGettingBatchDetails != nil {
				return "", errWhenGettingBatchDetails
			} else if batchDetails.BatchId == "" {
				//New record should be added with the relevant batch details
				//Create new folder for TDP id
			} else {
				//check if the TDP id exists
				tdpDetails, errWhenGettingTdpDetails := IpfsRepository.GetTdpDetails("tdpid", fileObj.TDPDetails.TdpID)
				if errWhenGettingTdpDetails != nil {
					return "", errWhenGettingTdpDetails
				} else if tdpDetails.TdpId == "" {
					//New record should be added with the relevant tdp details

					//check if the TDP image is available if so upload image
					//upload the tdo
				} else {
					//check if there is an image if so upload the image
					//if not upload the TDP itself
				}
			}
		}
	}

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
