package repository

import "MyPIPE/domain/model"

type CommentRepository interface {
	GetAll() ([]model.Comment, error)
	FindById() (*model.Comment, error)
	FindByUserId(userId model.UserID) ([]model.Comment, error)
	FindByMovieId(movieId model.MovieID) ([]model.Comment, error)
}
