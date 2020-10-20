package infra

import (
	"MyPIPE/domain/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"mime/multipart"
	"strconv"
)

type UserProfileImagePersistence struct{}

func NewUserProfileImagePersistence()*UserProfileImagePersistence{
	return &UserProfileImagePersistence{}
}

func (u UserProfileImagePersistence)Upload(file multipart.File,user *model.User)error{
	sess := session.Must(session.NewSession())
	bucketName := "mypipe-111"
	userIdString := strconv.FormatUint(uint64(user.ID),10)
	objectKey := "profile_images/" + userIdString + "/" + string(user.ProfileImageName)

	uploader := s3manager.NewUploader(sess)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})

	if err != nil {
		return err
	}

	return nil
}