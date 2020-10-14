package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
	queryService_infra "MyPIPE/infra/queryService"
)

type IndexPlayListInMovieListPage struct{}

func NewIndexPlayListInMovieListPage()*IndexPlayListInMovieListPage{
	return &IndexPlayListInMovieListPage{}
}

func (i IndexPlayListInMovieListPage)Find(findDTO FindDTO)*queryService.IndexPlayListInMovieListPageDTO{
	return queryService_infra.NewIndexPlayListInMovieListPage().Find(findDTO.UserID,findDTO.MovieID)
}

type FindDTO struct{
	UserID model.UserID
	MovieID model.MovieID
}

func NewFindDTO(userId model.UserID,movieId model.MovieID)*FindDTO{
	return &FindDTO{
		UserID:  userId,
		MovieID: movieId,
	}
}