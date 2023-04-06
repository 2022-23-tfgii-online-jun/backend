package aws

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	configs "github.com/emur-uy/backend/config"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Digital Ocean Spaces Configs
var endpoint = configs.Get().AwsEndpoint
var region = configs.Get().AwsRegionName
var accessKey = configs.Get().AwsAccessKey
var secretKey = configs.Get().AwsSecretKey
var bucketName = configs.Get().AwsBucketName

var sess *session.Session

// Init initializes the AWS package
func Init() (err error) {
	sess, err = session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:    &endpoint,
	},
	)
	return
}

// GetHeaderType determines the file header type based on the file name
func GetHeaderType(fileName string) string {
	if strings.Contains(fileName, "jpeg") || strings.Contains(fileName, "jpg") {
		return "image/jpeg"
	}

	if strings.Contains(fileName, "png") {
		return "image/png"
	}

	return "application/octet-stream"
}

// UploadFileToS3 uploads a file (fileName) to a specified bucket path (uploadPath)
func UploadFileToS3(fileName, uploadPath string) (string, error) {

	contentType := GetHeaderType(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("file open failed for %v, err %v", fileName, err.Error())
	}
	defer file.Close()
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(
		&s3manager.UploadInput{
			ACL:         aws.String("public-read"),
			Bucket:      aws.String(bucketName),
			Key:         aws.String(uploadPath),
			Body:        file,
			ContentType: &contentType,
			// Digital Ocean Spaces encrypts data at rest automatically, no additional settings needed
		})
	if err != nil {
		return "", fmt.Errorf("unable to upload %q to %q, err %v", fileName, bucketName, err)
	}
	return fmt.Sprintf("https://%s.%s.digitaloceanspaces.com/%s", bucketName, endpoint, uploadPath), nil
}

// UploadFileToS3Stream uploads a file to S3 using a stream.
// UploadFileToS3Stream uploads a file to S3 using a stream.
func UploadFileToS3Stream(file io.Reader, uploadPath string) (string, error) {
	// Create a new S3 service client
	s3Client := s3.New(sess)

	// Read the entire file into a buffer (this assumes the file is not too large to fit in memory)
	buffer := new(bytes.Buffer)
	if _, err := io.Copy(buffer, file); err != nil {
		return "", fmt.Errorf("failed to read file into buffer: %v", err)
	}

	// Upload the file to S3
	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(uploadPath),
		Body:   bytes.NewReader(buffer.Bytes()),
	}

	_, err := s3Client.PutObject(input)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %v", err)
	}

	// Generate the URL of the uploaded file
	fileURL := fmt.Sprintf("https://%s.%s.digitaloceanspaces.com/%s", bucketName, endpoint, uploadPath)

	return fileURL, nil
}

// GetFileLinkUsingKey generates a pre-signed URL for the specified bucket path (uploadPath) with an expiration time
func GetFileLinkUsingKey(uploadPath string, expirySeconds int) (string, error) {

	svc := s3.New(sess)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(uploadPath),
	})
	urlStr, err := req.Presign(time.Duration(expirySeconds) * time.Second)

	if err != nil {
		return "", fmt.Errorf("unable to get presigned link for %q, err %v", uploadPath, err)
	}
	return urlStr, err
}

// GetFileFromS3 downloads a file (fileName) from the bucket and saves it to the specified local path (outPath)
func GetFileFromS3(fileName, outPath string) (string, error) {
	downloader := s3manager.NewDownloader(sess)
	outFile := filepath.Join(outPath, fileName)
	file, err := os.Create(outFile)
	if err != nil {
		return "", fmt.Errorf("not able to create file, err : %v", err)
	}
	defer file.Close()
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(fileName),
		})
	if err != nil {
		return "", fmt.Errorf("unable to get %q from %q, err %v", fileName, bucketName, err)
	}
	return outFile, nil
}

// GetFileFromS3WithFileName downloads a file (fileName) from the bucket and saves it to the specified local path (outPath) with a custom name
func GetFileFromS3WithFileName(fileName, outPath, customFileName string) (string, error) {
	downloader := s3manager.NewDownloader(sess)
	outFile := filepath.Join(outPath, customFileName)
	file, err := os.Create(outFile)
	if err != nil {
		return "", fmt.Errorf("not able to create file, err : %v", err)
	}
	defer file.Close()
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(fileName),
		})
	if err != nil {
		return "", fmt.Errorf("unable to get %q from %q, err %v", fileName, bucketName, err)
	}
	return outFile, nil
}

// FileExistsS3 checks if the file (fileName) exists in the bucket
func FileExistsS3(fileName string) (bool, error) {
	svc := s3.New(sess)
	_, err := svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "NotFound": // s3.ErrCodeNoSuchKey does not work, AWS is missing this error code so we hardwire a string
				return false, nil
			default:
				return false, err
			}
		}
		return false, err
	}
	return true, nil
}

// GetFileBytesFromS3 retrieves the bytes of a file (fileName) from the bucket
func GetFileBytesFromS3(fileName string) ([]byte, error) {
	downloader := s3manager.NewDownloader(sess)

	buffer := &aws.WriteAtBuffer{}
	_, err := downloader.Download(buffer,
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(fileName),
		})
	if err != nil {
		return nil, fmt.Errorf("Unable to get %q from %q, err %v", fileName, bucketName, err)
	}
	return buffer.Bytes(), nil
}

// GetListFromS3 retrieves the list of files from the specified subfolder in the bucket
func GetListFromS3(subFolder string) ([]string, error) {
	svc := s3.New(sess)
	params := &s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(subFolder),
	}
	fmt.Println(params)
	resp, err := svc.ListObjects(params)
	if err != nil {
		return nil, fmt.Errorf("Unable to get list from %q, err %v", subFolder, err)
	}
	var result []string
	for _, key := range resp.Contents {
		result = append(result, *key.Key)
	}
	return result, nil
}

// DeleteObjectFromS3 deletes a file from the bucket at the specified path
func DeleteObjectFromS3(filePath string) error {
	svc := s3.New(sess)
	params := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filePath),
	}
	_, err := svc.DeleteObject(params)
	return err
}
