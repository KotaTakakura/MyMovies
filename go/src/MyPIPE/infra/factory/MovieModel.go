package factory

import (
	"MyPIPE/domain/model"
)

type MovieModelFactory struct{}

func NewMovieModelFactory() *MovieModelFactory {
	return &MovieModelFactory{}
}

func (m MovieModelFactory) CreateMovieModel(uploaderID model.UserID, movieFile *model.MovieFile, thumbnail *model.MovieThumbnail) (*model.Movie, error) {

	displayName, displayNameErr := model.NewMovieDisplayName("")
	if displayNameErr != nil {
		return nil, displayNameErr
	}

	return model.NewMovie(uploaderID, movieFile, displayName, thumbnail), nil
}
