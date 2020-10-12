package queryService_infra

import (
	"MyPIPE/domain/queryService"
	"MyPIPE/infra"
)

type IndexPlayListsInMyPage struct{}

func NewIndexPlayListsInMyPage()*IndexPlayListsInMyPage{
	return &IndexPlayListsInMyPage{}
}

func (i IndexPlayListsInMyPage)All(userId uint64)*queryService.IndexPlayListsInMyPageDTO{
	db := infra.ConnectGorm()
	defer db.Close()
	var playLists []queryService.PlayListForIndexPlayListsInMyPageDTO
	searchResult := db.Raw(
		"select play_lists.id as play_list_id, play_lists.name as play_list_name,play_lists.description as play_list_description,s.movie_id as play_list_first_movie_id,movies.thumbnail_name as play_list_thumbnail_name " +
			"from play_lists " +
			"left join (select play_list_id,movie_id,MAX(created_at) from play_list_movies group by play_list_id) as s on play_lists.id = s.play_list_id " +
			"left join movies on s.movie_id = movies.id " +
			"where play_lists.user_id = ?",userId).
		Scan(&playLists)
	var count uint64
	count = uint64(searchResult.RowsAffected)
	
	result := &queryService.IndexPlayListsInMyPageDTO{
		PlayLists:      playLists,
		PlayListsCount: count,
	}

	return result
}