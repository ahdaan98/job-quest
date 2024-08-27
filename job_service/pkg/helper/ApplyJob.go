package helper

import (
	"bytes"
	"fmt"
	cfg "job_service/pkg/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Helper struct {
	cfg cfg.Config
}

func NewHelper(cfg cfg.Config) *Helper {
	return &Helper{cfg: cfg}
}

func (h *Helper) AddImageToAwsS3(file []byte, filename string) (string, error) {
	config, err := cfg.LoadConfig()
	if err != nil {
		return "", err
	}

	fmt.Println("pppppppp", config.DBHost)
	fmt.Println("print1", config.AWSRegion)
	fmt.Println("print2", config.Access_key_ID)
	fmt.Println("print3", config.Secret_access_key)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			config.Access_key_ID,
			config.Secret_access_key,
			"",
		),
	})
	if err != nil {
		fmt.Println("erorrrr here", err)
		return "", err
	}

	uploader := s3manager.NewUploader(sess)
	bucketName := "jobquestbucket"

	fmt.Printf("Uploading file of size: %d bytes\n", len(file))

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(file),
		ContentType: aws.String("application/pdf"),
	})

	if err != nil {
		fmt.Println("erroorrrr 2", err)
		return "", err
	}
	fmt.Println("Bucket: ",bucketName)
	fmt.Println("aws region: ",config.AWSRegion)
	fmt.Println("file name: ", filename)
	fmt.Printf("Upload result: %+v\n", result)
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, config.AWSRegion, filename)
	return url, nil
}