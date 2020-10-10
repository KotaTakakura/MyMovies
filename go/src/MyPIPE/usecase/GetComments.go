package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
)

type GetComments struct{
	CommentQueryService queryService.CommentQueryService
}

func NewGetComments(cqs queryService.CommentQueryService)*GetComments{
	return &GetComments{
		CommentQueryService: cqs,
	}
}

func (g GetComments)Get(getDTO GetCommentsDTO)[]queryService.FindByUserIAndMovieIdDTO{
	return g.CommentQueryService.FindByMovieId(getDTO.MovieID)
}

type GetCommentsDTO struct{
	MovieID model.MovieID
}

func NewGetCommentsDTO(movieId model.MovieID)*GetCommentsDTO{
	return &GetCommentsDTO{
		MovieID: movieId,
	}
}