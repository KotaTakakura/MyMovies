package queryService

import "MyPIPE/domain/model"

type CommentQueryService interface {
	FindByMovieId(movieId model.MovieID) FindByMovieIdDTO
}

type FindByMovieIdDTO struct {
	Movie          MovieForFindByMovieIdDTO `json:"movie"`
	MovieLikeCount int                      `json:"movie_like_count"`
	User           UserForFindByMovieIdDTO  `json:"posted_user"`
	Comments       []CommentsFoundByMovieId `json:"comments"`
}

type MovieForFindByMovieIdDTO struct {
	ID          uint64
	UserID      uint64
	DisplayName string
	Description string
	Public		uint64
}

type UserForFindByMovieIdDTO struct {
	ID               uint64
	Name             string
	ProfileImageName string
}

type CommentsFoundByMovieId struct {
	CommentID            uint64 `json:"comment_id"`
	CommentBody          string `json:"comment_body"`
	UserName             string `json:"user_name"`
	UserProfileImageName string `json:"user_profile_image_name"`
	UserID               uint64 `json:"user_id"`
	MovieID              uint64 `json:"movie_id"`
}
