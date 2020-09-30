package factory

import (
	"MyPIPE/domain/model"
	"mime/multipart"
)

type IMovieModelFactory interface {
	CreateMovieModel(uploaderID model.UserID,displayName model.MovieDisplayName,fileHeader multipart.FileHeader)(*model.Movie,error)
}
//