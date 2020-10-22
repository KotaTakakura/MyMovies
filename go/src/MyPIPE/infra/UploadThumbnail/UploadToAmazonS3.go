package infra

import (
	"MyPIPE/domain/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"mime/multipart"
	"strconv"
)

type UploadThumbnailToAmazonS3 struct{}

func NewUploadThumbnailToAmazonS3() *UploadThumbnailToAmazonS3 {
	return &UploadThumbnailToAmazonS3{}
}

func (u UploadThumbnailToAmazonS3) Upload(file multipart.File, movie model.Movie) error {
	sess := session.Must(session.NewSession())
	bucketName := "mypipe-111"
	movieIdString := strconv.FormatUint(uint64(movie.ID), 10)
	objectKey := "thumbnails/" + movieIdString + string(movie.ThumbnailName)

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
