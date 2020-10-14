package queryService

import "MyPIPE/domain/model"

type IndexPlayListInMovieListPageDTO struct{
	PlayLists	[]PlayListForIndexPlayListInMovieListPageDTO
}

type PlayListForIndexPlayListInMovieListPageDTO struct{
	PlayListID uint64
	PlayListName string
	PlayListMovies []PlayListMoviesForIndexPlayListInMovieListPageDTO	`gorm:"ForeignKey:PlayListID;AssociationForeignKey:play_list_id"`
}

func (PlayListForIndexPlayListInMovieListPageDTO) TableName() string { return "play_lists" }

type PlayListMoviesForIndexPlayListInMovieListPageDTO struct{
	MovieID uint64
	PlayListID uint64
}

func (PlayListMoviesForIndexPlayListInMovieListPageDTO) TableName() string { return "play_list_movies" }

type IndexPlayListInMovieListPageQueryService interface{
	Find(userId model.UserID,movieId model.MovieID)*IndexPlayListInMovieListPageDTO
}
