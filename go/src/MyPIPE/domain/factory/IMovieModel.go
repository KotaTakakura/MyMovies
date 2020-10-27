package factory

import (
	"MyPIPE/domain/model"
)

type IMovieModelFactory interface {
	CreateMovieModel(uploaderID model.UserID, movieFile *model.MovieFile, thumbnail *model.MovieThumbnail) (*model.Movie, error)
}
