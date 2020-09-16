package infra

import (
	"MyPIPE/domain/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"mime/multipart"
	"path/filepath"
	"strconv"
)

type UploadToAmazonS3 struct{}

func NewUploadToAmazonS3()*UploadToAmazonS3{
	return &UploadToAmazonS3{}
}

func (u UploadToAmazonS3)Upload(movieFile multipart.File,movieFileHeader multipart.FileHeader,movieID model.MovieID)error{
	sess := session.Must(session.NewSession())
	extension := filepath.Ext(movieFileHeader.Filename)
	bucketName := "mypipe-before-encoded"
	objectKey := strconv.FormatUint(uint64(movieID), 10) + extension

	// Uploaderを作成し、ローカルファイルをアップロード
	uploader := s3manager.NewUploader(sess)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   movieFile,
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("done")
	return nil
}