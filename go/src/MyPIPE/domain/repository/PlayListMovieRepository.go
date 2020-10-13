package repository

import "MyPIPE/domain/model"

type PlayListMovieRepository interface {
	Save(playListMovie *model.PlayListMovie)error
	//SaveMultiple(playListMovies []model.PlayListMovie)error
}
