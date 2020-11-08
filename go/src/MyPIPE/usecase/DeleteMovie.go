package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
)

type IDeleteMovie interface {
	DeleteMovie(deleteMovieDTO *DeleteMovieDTO) error
}

type DeleteMovie struct {
	MovieRepository repository.MovieRepository
}

func NewDeleteMovie(movieRepository repository.MovieRepository) *DeleteMovie {
	return &DeleteMovie{
		MovieRepository: movieRepository,
	}
}

func (d DeleteMovie) DeleteMovie(deleteMovieDTO *DeleteMovieDTO) error {
	movie,findMovieErr := d.MovieRepository.FindByUserIdAndMovieId(deleteMovieDTO.UserID,deleteMovieDTO.MovieID)
	if findMovieErr != nil{
		return findMovieErr
	}
	if movie == nil{
		return errors.New("No Such Movie.")
	}

	deleteMovieErr := d.MovieRepository.Remove(deleteMovieDTO.UserID, deleteMovieDTO.MovieID)
	if deleteMovieErr != nil {
		return deleteMovieErr
	}
	return nil
}

type DeleteMovieDTO struct {
	UserID  model.UserID
	MovieID model.MovieID
}

func NewDeleteMovieDTO(userId model.UserID, movieId model.MovieID) *DeleteMovieDTO {
	return &DeleteMovieDTO{
		UserID:  userId,
		MovieID: movieId,
	}
}
