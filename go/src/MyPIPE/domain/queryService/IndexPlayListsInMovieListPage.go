package queryService

import "MyPIPE/domain/model"

type IndexPlayListInMovieListPageDTO struct {
	PlayLists []PlayListForIndexPlayListInMovieListPageDTO `json:"play_lists"`
}

type PlayListForIndexPlayListInMovieListPageDTO struct {
	PlayListID     uint64                                             `json:"play_list_id"`
	PlayListName   string                                             `json:"play_list_name"`
	PlayListMovies []PlayListMoviesForIndexPlayListInMovieListPageDTO `json:"play_list_movies" gorm:"ForeignKey:PlayListID;AssociationForeignKey:play_list_id"`
}

func (PlayListForIndexPlayListInMovieListPageDTO) TableName() string { return "play_lists" }

type PlayListMoviesForIndexPlayListInMovieListPageDTO struct {
	MovieID    uint64 `json:"movie_id"`
	PlayListID uint64 `json:"play_list_id"`
}

func (PlayListMoviesForIndexPlayListInMovieListPageDTO) TableName() string { return "play_list_movies" }

type IndexPlayListInMovieListPageQueryService interface {
	Find(userId model.UserID, movieId model.MovieID) *IndexPlayListInMovieListPageDTO
}
