package repository

import (
	"MyPIPE/domain/model"
	"mime/multipart"
)

type UserProfileImageRepository interface{
	Upload(file multipart.File,user *model.User)error
}