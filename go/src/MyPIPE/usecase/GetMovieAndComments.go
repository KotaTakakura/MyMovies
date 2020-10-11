package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
)

type GetMovieAndComments struct{
	CommentQueryService queryService.CommentQueryService
}

func NewGetMovieAndComments(cqs queryService.CommentQueryService)*GetMovieAndComments{
	return &GetMovieAndComments{
		CommentQueryService: cqs,
	}
}

func (g GetMovieAndComments)Get(getDTO MovieAndGetCommentsDTO)queryService.FindByMovieIdDTO{
	return g.CommentQueryService.FindByMovieId(getDTO.MovieID)
}

type MovieAndGetCommentsDTO struct{
	MovieID model.MovieID
}

func NewGetMovieAndCommentsDTO(movieId model.MovieID)*MovieAndGetCommentsDTO{
	return &MovieAndGetCommentsDTO{
		MovieID: movieId,
	}
}