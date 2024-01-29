package ipfsbusinessfacade

import (
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/dileepaj/tracified-nft-backend/businessFacade/marketplaceBusinessFacade"
	"github.com/dileepaj/tracified-nft-backend/constants"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/services/ipfsservice"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

func UploadCollectionsToIpfs(fileObj models.IpfsObjectForCollections) (string, error) {
	var requestType int
	var folderPath string

	switch fileObj.FileType {
	case constants.ImageFile:
		folderPath = "nftcollection/" + fileObj.CollectionDetails.UserId + "/" + fileObj.CollectionDetails.CollectionName
	// case constants.ImageFile:
	default:
		return "", errors.New("Invalid file type")
	}
	errWhenCreatingFolder := ipfsservice.CreateFolder(fileBaseBucket, folderPath)
	if errWhenCreatingFolder != nil {
		return "", errWhenCreatingFolder
	}
	//Check the request type
	if fileObj.FileType == constants.ImageFile {
		requestType = 2
	} else {
		return "", errors.New("Invalid request type")
	}
	cid, errWhenUploadingFileToIpfs := InitiateCollectionUpload(requestType, fileObj.FileDetails.FileContent, fileObj.FileDetails.FileName, folderPath)
	if errWhenUploadingFileToIpfs != nil {
		return "", errWhenUploadingFileToIpfs
	}
	insertObj := models.NFTCollection{
		UserId:           fileObj.CollectionDetails.UserId,
		Timestamp:        fileObj.CollectionDetails.Timestamp,
		CollectionName:   fileObj.CollectionDetails.CollectionName,
		OrganizationName: fileObj.CollectionDetails.OrganizationName,
		IsPublic:         fileObj.CollectionDetails.IsPublic,
		Publickey:        fileObj.CollectionDetails.Publickey,
	}

	if fileObj.FileType == 2 {
		img := models.ImageObject{
			ImageName: fileObj.FileDetails.FileName,
			ImageCid:  cid,
		}
		insertObj.CID = cid
		insertObj.Images = append(insertObj.Images, img)
	}

	_, errWhenSavingDetails := marketplaceBusinessFacade.CreateCollection(insertObj)
	if errWhenSavingDetails != nil {
		logs.ErrorLogger.Println("Error when saving file details on collection : ", errWhenSavingDetails)
		return "", errWhenSavingDetails
	}
	logs.InfoLogger.Println("CID Hash : " + cid)

	return cid, nil

}

// 1- TDP, 2 - Image
func InitiateCollectionUpload(fileType int, fileContent string, fileName string, folderName string) (string, error) {
	var fileNameInLocation string
	var dec []byte

	if fileType == constants.ImageFile {
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
