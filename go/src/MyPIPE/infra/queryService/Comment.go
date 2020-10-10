package queryService_infra

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
	"MyPIPE/infra"
	"fmt"
)

type CommentQueryService struct{}

func NewCommentQueryService()*CommentQueryService{
	return &CommentQueryService{}
}

func (c CommentQueryService)FindByMovieId(movieId model.MovieID)[]queryService.FindByUserIAndMovieIdDTO{
	db := infra.ConnectGorm()
	defer db.Close()
	var result []queryService.FindByUserIAndMovieIdDTO
	db.Table("comments").
		Select("comments.id as comment_id, comments.body as comment_body, comments.movie_id as movie_id,comments.user_id as user_id,users.name as user_name").
		Joins("inner join users on comments.user_id = users.id").
		Where("comments.movie_id = ?",movieId).
		Order("comments.created_at desc").
		Scan(&result)
	fmt.Println(result)
	return result
}