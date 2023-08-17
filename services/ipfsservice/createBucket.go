package ipfsservice

import (
	"errors"
	"os"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
)

func InitBucket(bucketName string) error {
	errWhenValidatingBucketName := ValidateBucketName(bucketName)
	if errWhenValidatingBucketName != nil {
		logs.ErrorLogger.Println("Validation failed for the bucket name : ", errWhenValidatingBucketName.Error())
		return errWhenValidatingBucketName
	}

	accessKey := os.Getenv("FILEBASE_ACCESS_KEY")
	secretKey := os.Getenv("FILEBASE_SECRET_KEY")
	endpoint := os.Getenv("FILEBASE_S3_API_ENDPOINT")
	region := os.Getenv("FILEBASE_REGION")
	profile := os.Getenv("FILEBASE_PROFILE")

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
		logs.ErrorLogger.Println("Error when creating filebase AWS session : ", errWhenCreatingSession.Error())
		return errWhenCreatingSession
	}

	s3Client := s3.New(goSession)
	bucket := aws.String(bucketName)

	_, errWhenCreatingBucket := s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: bucket,
	})
	if errWhenCreatingBucket != nil {
		logs.ErrorLogger.Println("Error when creating bucket : ", errWhenCreatingBucket.Error())
		return errWhenCreatingBucket
	}
	logs.InfoLogger.Println(bucketName + " bucket created")
	return nil
}

func ValidateBucketName(bucketName string) error {
	//bucket name should be in between 3-63 characters
	if len(bucketName) < 3 || len(bucketName) > 63 {
		return errors.New("Bucket name must be in between 3 and 63 characters")
	}

	//checking for lowercase characters, numbers, and dashes
	validBucketNameRegex := regexp.MustCompile(`^[a-z0-9\-]+$`)
	if !validBucketNameRegex.MatchString(bucketName) {
		return errors.New("Bucket name can only contain lowercase characters, numbers, and dashes")
	}

	return nil
}
