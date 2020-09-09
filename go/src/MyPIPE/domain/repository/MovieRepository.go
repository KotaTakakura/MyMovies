package repository

import "MyPIPE/domain/model"

type MovieRepository interface {
	GetAll() []model.Movie
	FindById(id model.MovieID) *model.Movie
	FindByUserId(userId model.MovieID) *model.Movie
}
