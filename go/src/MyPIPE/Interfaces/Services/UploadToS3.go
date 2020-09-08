package Iservices

import(
	"mime/multipart"
)

type IUploadToS3 interface{
	Upload(multipart.File)
}