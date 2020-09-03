package Services

import(
	"fmt"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFileToS3(file multipart.File){

	// セッションの作成（認証はここで行う）
    sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String("ap-northeast-1"),
			Credentials: credentials.NewEnvCredentials(),
		})

	if err != nil {
        panic(err)
	}

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)
	bucketName := "mypipe-112233"

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("sample11111.txt"),
		Body:   file,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(result.Location)
}