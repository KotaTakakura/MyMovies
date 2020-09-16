package repository

import "MyPIPE/domain/model"

type MovieRepository interface {
	GetAll() ([]model.Movie, error)
	FindById(id model.MovieID) (*model.Movie, error)
	FindByUserId(userId model.MovieID) (*model.Movie, error)
	Save(model.Movie) (*model.Movie,error)
}
