package queryService_infra

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
	"MyPIPE/infra"
)

type IndexPlayListInMovieListPage struct{}

func NewIndexPlayListInMovieListPage()*IndexPlayListInMovieListPage{
	return &IndexPlayListInMovieListPage{}
}

func (i IndexPlayListInMovieListPage)Find(userId model.UserID,movieId model.MovieID)*queryService.IndexPlayListInMovieListPageDTO{
	db := infra.ConnectGorm()
	defer db.Close()

	var playListForIndexPlayListInMovieListPageDTO []queryService.PlayListForIndexPlayListInMovieListPageDTO
	db.Table("play_lists").Where("play_lists.user_id = ?",userId).Select("id as play_list_id, name as play_list_name").Preload("PlayListMovies","movie_id = ?",movieId).Find(&playListForIndexPlayListInMovieListPageDTO)

	var indexPlayListInMovieListPageDTO queryService.IndexPlayListInMovieListPageDTO
	indexPlayListInMovieListPageDTO.PlayLists = playListForIndexPlayListInMovieListPageDTO

	return &indexPlayListInMovieListPageDTO
}
