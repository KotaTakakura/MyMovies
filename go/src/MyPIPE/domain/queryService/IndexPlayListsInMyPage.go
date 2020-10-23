package queryService

import "MyPIPE/domain/model"

type IndexPlayListsInMyPageDTO struct {
	PlayLists      []PlayListForIndexPlayListsInMyPageDTO `json:"play_lists"`
	PlayListsCount uint64                                 `json:"play_lists_count"`
}

type PlayListForIndexPlayListsInMyPageDTO struct {
	PlayListID                      uint64 `json:"play_list_id"`
	PlayListName                    string `json:"play_list_name"`
	PlayListDescription             string `json:"play_list_description"`
	PlayListFirstMovieID            uint64 `json:"play_list_first_movie_id"`
	PlayListFirstMovieThumbnailName string `json:"play_list_thumbnail_name"`
}

type IndexPlayListsInMyPageQueryService interface {
	All(userId model.UserID) *IndexPlayListsInMyPageDTO
}
