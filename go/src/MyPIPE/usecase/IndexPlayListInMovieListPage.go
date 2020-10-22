package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
)

type IIndexPlayListInMovieListPage interface {
	Find(findDTO FindDTO) *queryService.IndexPlayListInMovieListPageDTO
}

type IndexPlayListInMovieListPage struct {
	IndexPlayListInMovieListPageQueryService queryService.IndexPlayListInMovieListPageQueryService
}

func NewIndexPlayListInMovieListPage(indexPlayListInMovieListPage queryService.IndexPlayListInMovieListPageQueryService) *IndexPlayListInMovieListPage {
	return &IndexPlayListInMovieListPage{
		IndexPlayListInMovieListPageQueryService: indexPlayListInMovieListPage,
	}
}

func (i IndexPlayListInMovieListPage) Find(findDTO FindDTO) *queryService.IndexPlayListInMovieListPageDTO {
	return i.IndexPlayListInMovieListPageQueryService.Find(findDTO.UserID, findDTO.MovieID)
}

type FindDTO struct {
	UserID  model.UserID
	MovieID model.MovieID
}

func NewFindDTO(userId model.UserID, movieId model.MovieID) *FindDTO {
	return &FindDTO{
		UserID:  userId,
		MovieID: movieId,
	}
}
