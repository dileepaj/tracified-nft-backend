package ipfsservice

import (
	"errors"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dileepaj/tracified-nft-backend/commons"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

func UploadFile(pathToFile string, keyName string, bucketName string, folderName string) (string, string, error) {
	accessKey := os.Getenv("FILEBASE_ACCESS_KEY")
	secretKey := os.Getenv("FILEBASE_SECRET_KEY")
	endpoint := os.Getenv("FILEBASE_S3_API_ENDPOINT")
	region := os.Getenv("FILEBASE_REGION")
	profile := os.Getenv("FILEBASE_PROFILE")

	//create the configuration
	s3Config := aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(region),
		S3ForcePathStyle: aws.Bool(true),
	}

	goSession, errWhenCreatingSession := session.NewSessionWithOptions(session.Options{
		Config:  s3Config,
		Profile: profile,
	})
	if errWhenCreatingSession != nil {
		logs.ErrorLogger.Println("Error when creating session : ", errWhenCreatingSession.Error())
		return "", "", errWhenCreatingSession
	}

	//create the S3 client session
	s3Client := s3.New(goSession)

	//set the file path to upload
	file, errWhenReadingTheFile := os.Open(pathToFile)
	if errWhenReadingTheFile != nil {
		logs.ErrorLogger.Println("Error when reading the file : ", errWhenReadingTheFile.Error())
		return "", "", errWhenReadingTheFile
	}

	defer file.Close()

	var putObjectInput *s3.PutObjectInput
	if folderName == "" {
		//define key without folder
		putObjectInput = &s3.PutObjectInput{
			Body:   file,
			Bucket: aws.String(bucketName),
			Key:    aws.String(keyName),
		}
	} else {
		//define key with folder
		keyNameWithFolder := folderName + `/` + keyName
		putObjectInput = &s3.PutObjectInput{
			Body:   file,
			Bucket: aws.String(bucketName),
			Key:    aws.String(keyNameWithFolder),
		}
	}

	_, errWhenUploadingTheFile := s3Client.PutObject(putObjectInput)
	if errWhenUploadingTheFile != nil {
		logs.ErrorLogger.Println("Error when uploading the file to the ", bucketName, " bucket in filebase : ", errWhenUploadingTheFile.Error())
		return "", "", errWhenUploadingTheFile
	}

	var resp *s3.HeadObjectOutput
	var errWhenGettingHeadObject error
	if folderName == "" {
		resp, errWhenGettingHeadObject = s3Client.HeadObject(&s3.HeadObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(keyName),
		})
		if errWhenGettingHeadObject != nil {
			logs.ErrorLogger.Println("Error when getting the header object : ", errWhenGettingHeadObject)
			return "", "", errWhenGettingHeadObject
		}
	} else {
		keyNameWithFolder := folderName + `/` + keyName
		resp, errWhenGettingHeadObject = s3Client.HeadObject(&s3.HeadObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(keyNameWithFolder),
		})
		if errWhenGettingHeadObject != nil {
			logs.ErrorLogger.Println("Error when getting the header object : ", errWhenGettingHeadObject)
			return "", "", errWhenGettingHeadObject
		}
	}
	cid := ""
	if resp.Metadata != nil {
		cidValue, ok := resp.Metadata["Cid"]
		if !ok {
			logs.ErrorLogger.Println("Error when getting CID ")
			return "", "", errors.New("Error when getting the D")
		}
		cid = *cidValue
	} else {
		logs.ErrorLogger.Println("CID is not created")
		return "", "", errors.New("No CID is created")
	}

	link := commons.GoDotEnvVariable("IPFSURL") + cid

	logs.InfoLogger.Println("Content uploaded to IPFS at : " + link)
	return cid, link, nil
}
