package ipfsservice

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

func CreateFolder(bucketName string, folderName string) error {
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
		return errWhenCreatingSession
	}

	s3Client := s3.New(goSession)

	folderKey := folderName + `/`

	// Create an empty object with the specified key (folder)
	_, errWhenCreatingFolder := s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(folderKey),
		Body:   nil, // Empty body for a folder
	})
	if errWhenCreatingFolder != nil {
		logs.ErrorLogger.Println("Error when creating folder : ", errWhenCreatingFolder.Error())
		return errWhenCreatingFolder
	}

	logs.InfoLogger.Println("Folder created under bucket " + bucketName)

	return nil

}
