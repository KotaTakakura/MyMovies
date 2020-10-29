package queryService_infra

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
	"MyPIPE/infra"
)

type IndexPlayListMovieInMyPage struct{}

func NewIndexPlayListMovieInMyPage() *IndexPlayListMovieInMyPage {
	return &IndexPlayListMovieInMyPage{}
}

func (i IndexPlayListMovieInMyPage) Find(userId model.UserID, playListId model.PlayListID) *queryService.IndexPlayListMovieInMyPageDTO {
	db := infra.ConnectGorm()
	defer db.Close()
	var result queryService.IndexPlayListMovieInMyPageDTO

	var playListMovies []queryService.PlayListMovieForIndexPlayListMovieInMyPageDTO
	db.Raw("select movies.id as movie_id,"+
		"movies.display_name as movie_title,"+
		"movies.description as movie_description,"+
		"movies.thumbnail_name as movie_thumbnail_name,"+
		"play_list_movies.order as `order` "+
		"from (select * from play_lists where user_id = ? and id = ?) as s "+
		"left join play_list_movies on play_list_movies.play_list_id = s.id "+
		"left join movies on play_list_movies.movie_id = movies.id "+
		"order by `order`", userId, playListId).
		Scan(&playListMovies)
	if len(playListMovies) == 1 && playListMovies[0].MovieID == 0 {
		playListMovies = nil
	}

	var playList queryService.PlayListForIndexPlayListMovieInMyPageDTO
	db.Raw("select play_lists.id as play_list_id,"+
		"name as play_list_name,play_lists.description as play_list_description,"+
		"first_movie.thumbnail_name as play_list_first_movie_thumbnail_name,"+
		"thumbnail_movie.id as play_list_thumbnail_movie_id,"+
		"first_movie.id as play_list_first_movie_id,"+
		"thumbnail_movie.thumbnail_name as play_list_thumbnail_name "+
		"from play_lists "+
		"left join (select play_list_id,movie_id,MAX(created_at) from play_list_movies group by play_list_id) as s on play_lists.id = s.play_list_id "+
		"left join movies as thumbnail_movie on play_lists.thumbnail_movie_id = thumbnail_movie.id "+
		"left join movies as first_movie on s.movie_id = first_movie.id "+
		"where play_lists.id = ? and play_lists.user_id = ?", playListId, userId).Scan(&playList)

	result.PlayList = playList
	result.PlayListMovies = playListMovies

	return &result
}
