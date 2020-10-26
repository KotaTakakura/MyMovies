package factory

import (
	"MyPIPE/domain/model"
	"mime/multipart"
)

type IMovieModelFactory interface {
	CreateMovieModel(uploaderID model.UserID, fileHeader multipart.FileHeader, thumbnail model.MovieThumbnail) (*model.Movie, error)
}
