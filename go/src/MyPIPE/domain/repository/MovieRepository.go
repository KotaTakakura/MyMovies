package repository

import "MyPIPE/domain/model"

type MovieRepository interface {
	GetAll() []model.Movie
	FindById(id uint64) *model.Movie
	FindByUserId(userId uint64)	*model.Movie
}