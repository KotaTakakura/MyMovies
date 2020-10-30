package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type IUpdateMovie interface {
	Update(updateDTO *UpdateDTO) (*model.Movie, error)
	UpdateStatus(updateStatusDTO *UpdateStatusDTO) error
	UpdateThumbnailStatus(updateThumbnailStatusDTO *UpdateThumbnailStatusDTO) error
}

type UpdateMovie struct {
	MovieRepository repository.MovieRepository
}

func NewUpdateMovie(m repository.MovieRepository) *UpdateMovie {
	return &UpdateMovie{
		MovieRepository: m,
	}
}

func (u UpdateMovie) Update(updateDTO *UpdateDTO) (*model.Movie, error) {
	movie, findMovieErr := u.MovieRepository.FindByUserIdAndMovieId(updateDTO.UserID, updateDTO.MovieID)
	if findMovieErr != nil {
		return nil, findMovieErr
	}
	changeDisplayNameErr := movie.ChangeDisplayName(updateDTO.DisplayName)
	if changeDisplayNameErr != nil {
		return nil, changeDisplayNameErr
	}

	changeDescriptionErr := movie.ChangeDescription(updateDTO.Description)
	if changeDescriptionErr != nil {
		return nil, changeDescriptionErr
	}

	changeStatusErr := movie.ChangeStatus(updateDTO.Status)
	if changeStatusErr != nil {
		return nil, changeStatusErr
	}

	changePublicErr := movie.ChangePublic(updateDTO.Public)
	if changePublicErr != nil {
		return nil, changePublicErr
	}

	updatedMovie, updateMovieErr := u.MovieRepository.Update(*movie)
	if updateMovieErr != nil {
		return nil, updateMovieErr
	}

	return updatedMovie, nil
}

type UpdateDTO struct {
	UserID      model.UserID
	MovieID     model.MovieID
	DisplayName model.MovieDisplayName
	Description model.MovieDescription
	Public      model.MoviePublic
	Status      model.MovieStatus
}

func (u UpdateMovie) UpdateStatus(updateStatusDTO *UpdateStatusDTO) error {
	movie, findMovieErr := u.MovieRepository.FindById(updateStatusDTO.MovieID)
	if findMovieErr != nil {
		return findMovieErr
	}

	changeStatusErr := movie.Complete()
	if changeStatusErr != nil {
		return changeStatusErr
	}

	_, updateMovieErr := u.MovieRepository.Update(*movie)
	if updateMovieErr != nil {
		return updateMovieErr
	}

	return nil
}

type UpdateStatusDTO struct {
	MovieID model.MovieID
}

func NewUpdateStatusDTO(movieId model.MovieID) *UpdateStatusDTO {
	return &UpdateStatusDTO{
		MovieID: movieId,
	}
}

func (u UpdateMovie) UpdateThumbnailStatus(updateThumbnailStatusDTO *UpdateThumbnailStatusDTO) error {
	movie, findMovieErr := u.MovieRepository.FindById(updateThumbnailStatusDTO.MovieID)
	if findMovieErr != nil {
		return findMovieErr
	}

	changeStatusErr := movie.ChangeThumbnailStatusComplete()
	if changeStatusErr != nil {
		return changeStatusErr
	}

	_, updateMovieErr := u.MovieRepository.Update(*movie)
	if updateMovieErr != nil {
		return updateMovieErr
	}

	return nil
}

type UpdateThumbnailStatusDTO struct {
	MovieID model.MovieID
}

func NewUpdateThumbnailStatusDTO(movieId model.MovieID) *UpdateThumbnailStatusDTO {
	return &UpdateThumbnailStatusDTO{
		MovieID: movieId,
	}
}
