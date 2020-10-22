package repository

import "MyPIPE/domain/model"

type PlayListMovieRepository interface {
	Find(userId model.UserID, playListId model.PlayListID, movieId model.MovieID) *model.PlayListMovie
	FindAll(userId model.UserID, playListId model.PlayListID) []model.PlayListMovie
	Save(playListMovie *model.PlayListMovie) error
	Remove(*model.PlayListMovie) error
	SaveAll([]model.PlayListMovie) error
	//RemoveAll([]model.PlayListMovie)error
	//SaveMultiple(playListMovies []model.PlayListMovie)error
}
