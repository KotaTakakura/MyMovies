package queryService

type IndexPlayListMovieInMyPageDTO struct {
	PlayList       PlayListForIndexPlayListMovieInMyPageDTO        `json:"play_list"`
	PlayListMovies []PlayListMovieForIndexPlayListMovieInMyPageDTO `json:"play_list_movies"`
}

type PlayListForIndexPlayListMovieInMyPageDTO struct {
	PlayListID          uint64 `json:"play_list_id"`
	PlayListName        string `json:"play_list_name"`
	PlaylistDescription string `json:"play_list_description"`
}

type PlayListMovieForIndexPlayListMovieInMyPageDTO struct {
	MovieID            uint64 `json:"movie_id"`
	MovieTitle         string `json:"movie_title"`
	MovieDescription   string `json:"movie_description"`
	MovieThumbnailName string `json:"movie_thumbnail_name"`
	Order              int    `json:"order"`
}

type IndexPlayListMovieQueryService interface {
	Find(userId uint64, playListId uint64) *IndexPlayListMovieInMyPageDTO
}
