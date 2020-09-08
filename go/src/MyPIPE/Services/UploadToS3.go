package Services

import (
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type UploadToS3 struct {

}

func NewUploadToS3() *UploadToS3{
	return &UploadToS3{}
}

func (u UploadToS3) Upload (file multipart.File){
	creds := credentials.NewStaticCredentials("", "", "")
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: creds,
		Region: aws.String("ap-northeast-1"),
	}))
	bucketName := "mypipe-111"
	objectKey := "TTT.rtf"
	uploader := s3manager.NewUploader(sess)
	var err error
	test := &s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		ContentType: aws.String("text/plain"),
		Body:   file,
	}
	_, err = uploader.Upload(test)

	if err != nil {
		log.Println(err)
	}
	log.Println("done")
}