package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
)

type IDeletePlayListMovie interface {
	DeletePlayListItem(playListItemDeleteJson *DeletePlayListMovieDTO) error
}

type DeletePlayListMovie struct {
	PlayListRepository      repository.PlayListRepository
	PlayListMovieRepository repository.PlayListMovieRepository
}

func NewDeletePlayListMovie(p repository.PlayListRepository, plmr repository.PlayListMovieRepository) *DeletePlayListMovie {
	return &DeletePlayListMovie{
		PlayListRepository:      p,
		PlayListMovieRepository: plmr,
	}
}

func (a DeletePlayListMovie) DeletePlayListItem(playListItemDeleteJson *DeletePlayListMovieDTO) error {

	playList, playListFindErr := a.PlayListRepository.FindByID(playListItemDeleteJson.PlayListID)
	if playList == nil || playList.UserID != playListItemDeleteJson.UserID {
		return errors.New("No Such PlayList.")
	}
	if playListFindErr != nil {
		return playListFindErr
	}

	playListMovie := a.PlayListMovieRepository.Find(playListItemDeleteJson.UserID, playListItemDeleteJson.PlayListID, playListItemDeleteJson.MovieID)
	if playListMovie == nil {
		return errors.New("No Such PlayListMovie.")
	}

	removePlayListMovieErr := a.PlayListMovieRepository.Remove(playListMovie)
	if removePlayListMovieErr != nil {
		return removePlayListMovieErr
	}

	return nil

}

type DeletePlayListMovieDTO struct {
	PlayListID model.PlayListID
	UserID     model.UserID
	MovieID    model.MovieID
}

func NewDeletePlayListMovieJson(playListID model.PlayListID, userId model.UserID, movieId model.MovieID) *DeletePlayListMovieDTO {
	return &DeletePlayListMovieDTO{
		PlayListID: playListID,
		UserID:     userId,
		MovieID:    movieId,
	}
}
