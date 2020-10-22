package queryService_infra

import (
	"MyPIPE/domain/queryService"
	"MyPIPE/infra"
)

type IndexPlayListMovieInMyPage struct{}

func NewIndexPlayListMovieInMyPage() *IndexPlayListMovieInMyPage {
	return &IndexPlayListMovieInMyPage{}
}

func (i IndexPlayListMovieInMyPage) Find(userId uint64, playListId uint64) *queryService.IndexPlayListMovieInMyPageDTO {
	db := infra.ConnectGorm()
	defer db.Close()
	var result queryService.IndexPlayListMovieInMyPageDTO

	var playListMovies []queryService.PlayListMovieForIndexPlayListMovieInMyPageDTO
	db.Raw("select movies.id as movie_id,movies.display_name as movie_title,movies.description as movie_description, movies.thumbnail_name as movie_thumbnail_name,play_list_movies.order as `order` "+
		"from (select * from play_lists where user_id = ? and id = ?) as s "+
		"left join play_list_movies on play_list_movies.play_list_id = s.id "+
		"left join movies on play_list_movies.movie_id = movies.id "+
		"order by `order`", userId, playListId).
		Scan(&playListMovies)
	if len(playListMovies) == 1 && playListMovies[0].MovieID == 0 {
		playListMovies = nil
	}

	var playList queryService.PlayListForIndexPlayListMovieInMyPageDTO
	db.Raw("select id as play_list_id,name as play_list_name,description as play_list_description from play_lists where id = ? and user_id = ?", playListId, userId).Scan(&playList)

	result.PlayList = playList
	result.PlayListMovies = playListMovies

	return &result
}
