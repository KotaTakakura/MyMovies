package repository

import "MyPIPE/domain/model"

type MovieRepository interface {
	GetAll() ([]model.Movie, error)
	FindById(id model.MovieID) (*model.Movie, error)
	FindByUserId(userId model.MovieID) (*model.Movie, error)
	FindByUserIdAndMovieId(userId model.UserID, movieId model.MovieID) (*model.Movie, error)
	Save(model.Movie) (*model.Movie, error)
	Update(movie model.Movie) (*model.Movie, error)
}
