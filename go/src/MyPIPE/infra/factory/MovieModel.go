package factory

import (
	"MyPIPE/domain/model"
)

type MovieModelFactory struct{}

func NewMovieModelFactory() *MovieModelFactory {
	return &MovieModelFactory{}
}

func (m MovieModelFactory) CreateMovieModel(uploaderID model.UserID, movieFile *model.MovieFile, thumbnail *model.MovieThumbnail) (*model.Movie, error) {

	return model.NewMovie(uploaderID, movieFile, thumbnail), nil
}
