package repository

import "MyPIPE/domain/model"

type CommentRepository interface {
	GetAll() ([]model.Comment, error)
	FindById(commentID model.CommentID) (*model.Comment, error)
	FindByIdAndUserID(commentID model.CommentID,userID model.UserID) (*model.Comment, error)
	FindByUserId(userId model.UserID) ([]model.Comment, error)
	FindByMovieId(movieId model.MovieID) ([]model.Comment, error)
	Save(*model.Comment) error
	Remove(*model.Comment) error
}
