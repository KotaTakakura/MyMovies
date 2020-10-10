package queryService

import "MyPIPE/domain/model"


type CommentQueryService interface {
	FindByMovieId(movieId model.MovieID)[]FindByUserIAndMovieIdDTO
}

type FindByUserIAndMovieIdDTO struct{
	CommentID model.CommentID	`json:"comment_id"`
	CommentBody model.CommentBody	`json:"comment_body"`
	UserName model.UserName	`json:"user_name"`
	UserID model.UserID	`json:"user_id"`
	MovieID model.MovieID	`json:"movie_id"`
}
