package queryService_infra

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
	"MyPIPE/infra"
)

type CommentQueryService struct{}

func NewCommentQueryService()*CommentQueryService{
	return &CommentQueryService{}
}

func (c CommentQueryService)FindByMovieId(movieId model.MovieID)queryService.FindByMovieIdDTO{
	db := infra.ConnectGorm()
	defer db.Close()
	var comments []queryService.CommentsFoundByMovieId
	db.Table("comments").
		Select("comments.id as comment_id, comments.body as comment_body, comments.movie_id as movie_id,comments.user_id as user_id,users.name as user_name").
		Joins("inner join users on comments.user_id = users.id").
		Where("comments.movie_id = ?",movieId).
		Order("comments.created_at desc").
		Scan(&comments)

	var movie  queryService.MovieForFindByMovieIdDTO
	db.Table("movies").
		Where("id = ?",movieId).
		Take(&movie)

	var likeCount int
	db.Table("movie_evaluations").
		Where("movie_id = ? and evaluation = 0",movieId).
		Count(&likeCount)

	var user queryService.UserForFindByMovieIdDTO
	db.Table("users").
		Where("id = ?",movie.UserID).
		Take(&user)

	result := queryService.FindByMovieIdDTO{
		Movie:    movie,
		MovieLikeCount: likeCount,
		Comments: comments,
		User: user,
	}

	return result
}