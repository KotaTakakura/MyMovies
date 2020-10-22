package usecase

import (
	"MyPIPE/domain/factory"
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
)

type IAddPlayListItem interface {
	AddPlayListItem(playListItemAddJson AddPlayListItemAddJson) error
}

type AddPlayListItem struct {
	PlayListRepository      repository.PlayListRepository
	PlayListMovieRepository repository.PlayListMovieRepository
	PlaylistMovieFactory    factory.IPlayListMovie
}

func NewAddPlayListItem(p repository.PlayListRepository, plmr repository.PlayListMovieRepository, plmf factory.IPlayListMovie) *AddPlayListItem {
	return &AddPlayListItem{
		PlayListRepository:      p,
		PlayListMovieRepository: plmr,
		PlaylistMovieFactory:    plmf,
	}
}

func (a AddPlayListItem) AddPlayListItem(playListItemAddJson AddPlayListItemAddJson) error {

	playList, playListFindErr := a.PlayListRepository.FindByID(playListItemAddJson.PlayListID)
	if playList == nil || playList.UserID != playListItemAddJson.UserID {
		return errors.New("No Such PlayList.")
	}
	if playListFindErr != nil {
		return playListFindErr
	}

	playListMovie, err := a.PlaylistMovieFactory.CreatePlayListMovie(playListItemAddJson.PlayListID, playListItemAddJson.MovieID)
	if err != nil {
		return err
	}
	savePlayListMovieErr := a.PlayListMovieRepository.Save(playListMovie)
	if savePlayListMovieErr != nil {
		return savePlayListMovieErr
	}

	return nil

}

type AddPlayListItemAddJson struct {
	PlayListID model.PlayListID
	UserID     model.UserID
	MovieID    model.MovieID
}
