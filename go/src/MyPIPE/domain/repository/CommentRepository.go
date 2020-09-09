package repository

import "MyPIPE/domain/model"

type CommentRepository interface {
	GetAll() []model.Comment
	FindById() *model.Comment
	FindByUserId(userId model.UserID) []model.Comment
	FindByMovieId(movieId model.MovieID) []model.Comment
}
