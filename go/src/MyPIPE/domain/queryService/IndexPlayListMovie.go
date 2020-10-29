package queryService

import "MyPIPE/domain/model"

type IndexPlayListMovieInMyPageDTO struct {
	PlayList       PlayListForIndexPlayListMovieInMyPageDTO        `json:"play_list"`
	PlayListMovies []PlayListMovieForIndexPlayListMovieInMyPageDTO `json:"play_list_movies"`
}

type PlayListForIndexPlayListMovieInMyPageDTO struct {
	PlayListID                      uint64 `json:"play_list_id"`
	PlayListName                    string `json:"play_list_name"`
	PlaylistDescription             string `json:"play_list_description"`
	PlayListFirstMovieID            uint64 `json:"play_list_first_movie_id"`
	PlayListFirstMovieThumbnailName string `json:"play_list_first_movie_thumbnail_name"`
	PlayListThumbnailMovieID        uint64 `json:"play_list_thumbnail_movie_id"`
	PlayListThumbnailName           string `json:"play_list_thumbnail_name"`
}

type PlayListMovieForIndexPlayListMovieInMyPageDTO struct {
	MovieID            uint64 `json:"movie_id"`
	MovieTitle         string `json:"movie_title"`
	MovieDescription   string `json:"movie_description"`
	MovieThumbnailName string `json:"movie_thumbnail_name"`
	Order              int    `json:"order"`
}

type IndexPlayListMovieQueryService interface {
	Find(userId model.UserID, playListId model.PlayListID) *IndexPlayListMovieInMyPageDTO
}
