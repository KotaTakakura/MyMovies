package repository

import (
	"MyPIPE/domain/model"
	"mime/multipart"
)

type ThumbnailUploadRepository interface {
	Upload(file multipart.File,movie model.Movie) error
}