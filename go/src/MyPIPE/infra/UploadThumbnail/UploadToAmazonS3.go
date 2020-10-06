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

type UploadThumbnailToAmazonS3 struct{}

func NewUploadThumbnailToAmazonS3()*UploadThumbnailToAmazonS3{
	return &UploadThumbnailToAmazonS3{}
}

func (u UploadThumbnailToAmazonS3)Upload(file multipart.File,movieFileHeader multipart.FileHeader,movieID model.MovieID) error{
	sess := session.Must(session.NewSession())
	extension := filepath.Ext(movieFileHeader.Filename)
	bucketName := "mypipe-111"
	objectKey := "thumbnails/" + strconv.FormatUint(uint64(movieID), 10) + extension

	// Uploaderを作成し、ローカルファイルをアップロード
	uploader := s3manager.NewUploader(sess)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("done")
	return nil
}