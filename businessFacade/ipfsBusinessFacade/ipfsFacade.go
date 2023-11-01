package ipfsbusinessfacade

import (
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/constants"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services/ipfsservice"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

var (
	fileBaseBucket = os.Getenv("FILEBASE_BUCKET")
)

// Request type 1 -> TDP, 2 -> Image
func UploadFilesToIpfs(fileObj models.IpfsObjectForTDP) (string, error) {
	var requestType int
	var folderPath string

	//set up the folder path
	switch fileObj.FileType {
	case constants.TdpFile:
		folderPath = "tracabilitydatapackets/" + fileObj.TDPDetails.TenetID + "/" + fileObj.TDPDetails.ItemID + "/" + fileObj.TDPDetails.BatchID + "/" + fileObj.TDPDetails.TdpID
	case constants.ImageFile:
		folderPath = "tracabilitydatapackets/" + fileObj.TDPDetails.TenetID + "/" + fileObj.TDPDetails.ItemID + "/" + fileObj.TDPDetails.BatchID + "/" + fileObj.TDPDetails.TdpID + "/Images"
	default:
		return "", errors.New("Invalid file type")
	}
	errWhenCreatingFolder := ipfsservice.CreateFolder(fileBaseBucket, folderPath)
	if errWhenCreatingFolder != nil {
		return "", errWhenCreatingFolder
	}
	//Check the request type
	if fileObj.FileType == constants.TdpFile {
		requestType = 1
	} else if fileObj.FileType == constants.ImageFile {
		requestType = 2
	} else {
		return "", errors.New("Invalid request type")
	}
	cid, errWhenUploadingFileToIpfs := InitiateUpload(requestType, fileObj.FileDetails.FileContent, fileObj.FileDetails.FileName, folderPath)
	if errWhenUploadingFileToIpfs != nil {
		return "", errWhenUploadingFileToIpfs
	}

	//check if the TDP is already added
	tdpDetails, errWhenGettingTdpDetails := IpfsRepository.GetTdpDetails("tdpid", fileObj.TDPDetails.TdpID)
	if errWhenGettingTdpDetails != nil {
		return "", errWhenGettingTdpDetails
	} else if tdpDetails.TdpId == "" {
		//enter new record
		insertObj := models.TracifiedDataPacket{
			TenetId: fileObj.TDPDetails.TenetID,
			ItemId:  fileObj.TDPDetails.ItemID,
			BatchId: fileObj.TDPDetails.BatchID,
			TdpId:   fileObj.TDPDetails.TdpID,
		}

		if fileObj.FileType == 1 {
			insertObj.TdpCid = cid
			insertObj.Images = nil
		} else if fileObj.FileType == 2 {
			img := models.ImageObject{
				ImageName: fileObj.FileDetails.FileName,
				ImageCid:  cid,
			}
			insertObj.TdpCid = ""
			insertObj.Images = append(insertObj.Images, img)
		}

		_, errWhenSavingDetails := IpfsRepository.SaveFileDetails(insertObj)
		if errWhenSavingDetails != nil {
			logs.ErrorLogger.Println("Error when saving file details on collection : ", errWhenSavingDetails)
			return "", errWhenSavingDetails
		}
	} else {
		//update the current record
		updateObj := models.TracifiedDataPacket{
			TenetId: fileObj.TDPDetails.TenetID,
			ItemId:  fileObj.TDPDetails.ItemID,
			BatchId: fileObj.TDPDetails.BatchID,
			TdpId:   fileObj.TDPDetails.TdpID,
		}
		if fileObj.FileType == 1 {
			updateObj.TdpCid = cid
			updateObj.Images = tdpDetails.Images
		} else if fileObj.FileType == 2 {
			img := models.ImageObject{
				ImageName: fileObj.FileDetails.FileName,
				ImageCid:  cid,
			}
			updateObj.TdpCid = tdpDetails.TdpCid
			updateObj.Images = append(updateObj.Images, img)
		}
		_, errWhenUpdatingDetails := IpfsRepository.UpdateFileDetails(fileObj.TDPDetails.TdpID, updateObj)
		if errWhenUpdatingDetails != nil {
			logs.ErrorLogger.Println("Error when updating the collection : ", errWhenUpdatingDetails)
			return "", errWhenUpdatingDetails
		}

	}

	logs.InfoLogger.Println("CID Hash : " + cid)

	return cid, nil
}

// 1- TDP, 2 - Image
func InitiateUpload(fileType int, fileContent string, fileName string, folderName string) (string, error) {
	var fileNameInLocation string
	var dec []byte

	if fileType == constants.TdpFile {
		//upload TDP
		jsonContent := []byte(fileContent)
		fileNameInLocation = fileName + ".txt"
		fileNameInLocation = strings.ToLower(fileNameInLocation)
		dec = jsonContent
	} else if fileType == constants.ImageFile {
		//upload image
		imageStrArray := strings.Split(fileContent, ";base64,")
		decoded, errWhenDecodingString := base64.StdEncoding.DecodeString(imageStrArray[1])
		extentionType := strings.Split(imageStrArray[0], "data:image/")[1]
		if errWhenDecodingString != nil {
			logs.ErrorLogger.Println("Error when decoding image data : ", errWhenDecodingString)
			return "", errWhenDecodingString
		}
		dec = decoded
		fileNameInLocation = fileName + "." + extentionType
		fileNameInLocation = strings.ToLower(fileNameInLocation)

	} else {
		return "", errors.New("Invalid file type")
	}
	//get current working directory
	workingDirectory, errWhenGettingTheDirectory := os.Getwd()
	if errWhenGettingTheDirectory != nil {
		logs.ErrorLogger.Println("Error when getting the working directory : ", errWhenGettingTheDirectory.Error())
		return "", errWhenGettingTheDirectory
	}

	//create file
	filePath := filepath.Join(workingDirectory, fileNameInLocation)
	file, errWhenCreatingFile := os.Create(filePath)
	if errWhenCreatingFile != nil {
		logs.ErrorLogger.Println("Error when creating")
		return "", errWhenCreatingFile
	}
	defer file.Close()

	//write into the created file
	if _, errWhenWritingToFile := file.Write(dec); errWhenWritingToFile != nil {
		logs.ErrorLogger.Println("Error when writing data into the file : ", errWhenWritingToFile.Error())
		return "", errWhenWritingToFile
	}
	if errWhenSyncing := file.Sync(); errWhenSyncing != nil {
		logs.ErrorLogger.Println("Error when and clearing memory : ", errWhenSyncing.Error())
		return "", errWhenSyncing
	}

	cid, _, errWhenUploadingToIpfs := ipfsservice.UploadFile(filePath, fileNameInLocation, fileBaseBucket, folderName)
	if errWhenUploadingToIpfs != nil {
		logs.ErrorLogger.Println("Error when uploading to IPFS : ", errWhenUploadingToIpfs)
		//delete the image
		errWhenRemovingFile := os.Remove(filePath)
		if errWhenRemovingFile != nil {
			logs.ErrorLogger.Println("Error when removing the file : ", errWhenRemovingFile)
			return "", errWhenRemovingFile
		}
		return "", errWhenUploadingToIpfs
	}

	errWhenClosingTheFile := file.Close()
	if errWhenClosingTheFile != nil {
		logs.ErrorLogger.Println("Error when closing the file : ", errWhenClosingTheFile)
		return "", errWhenClosingTheFile
	}

	errWhenRemovingTheFile := os.Remove(filePath)
	if errWhenRemovingTheFile != nil {
		logs.ErrorLogger.Println("Error when removing the file : ", errWhenRemovingTheFile)
		return "", errWhenRemovingTheFile
	}
	return cid, nil
}
