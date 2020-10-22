package factory

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"errors"
)

type PlayListMovieFactory struct{}

func NewPlayListMovieFactory() *PlayListMovieFactory {
	return &PlayListMovieFactory{}
}

func (p PlayListMovieFactory) CreatePlayListMovie(playListId model.PlayListID, movieId model.MovieID) (*model.PlayListMovie, error) {
	db := infra.ConnectGorm()
	defer db.Close()
	var movieCount int
	db.Table("movies").Where("id = ?", movieId).Count(&movieCount)
	if movieCount == 0 {
		return nil, errors.New("No Such Movie.")
	}

	type PlayListMovieCount struct {
		LastCount int
	}
	var playListMovieCount PlayListMovieCount
	db.Table("play_list_movies").Where("play_list_id = ?", playListId).Select("Max(`order`) as last_count").First(&playListMovieCount)
	playListMovieOrder, err := model.NewPlayListMovieOrder(playListMovieCount.LastCount + 1)
	if err != nil {
		return nil, err
	}

	playListMovie := model.NewPlayListMovie(playListId, movieId, playListMovieOrder)
	return playListMovie, nil
}
