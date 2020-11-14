package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type IUpdateMovie interface {
	Update(updateDTO *UpdateDTO) (*model.Movie, error)
	UpdateStatus(updateStatusDTO *UpdateStatusDTO) error
	UpdateStatusError(updateStatusErrorDTO *UpdateStatusErrorDTO) error
	UpdateThumbnailStatus(updateThumbnailStatusDTO *UpdateThumbnailStatusDTO) error
}

type UpdateMovie struct {
	MovieRepository repository.MovieRepository
	MovieStatusRepository repository.MovieStatusRepository
}

func NewUpdateMovie(m repository.MovieRepository,msr repository.MovieStatusRepository) *UpdateMovie {
	return &UpdateMovie{
		MovieRepository: m,
		MovieStatusRepository: msr,
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
	movieStatus, findMovieErr := u.MovieStatusRepository.Find(updateStatusDTO.MovieID)
	if findMovieErr != nil {
		return findMovieErr
	}

	changeStatusErr := movieStatus.Complete()
	if changeStatusErr != nil {
		return changeStatusErr
	}

	updateMovieErr := u.MovieStatusRepository.Save(movieStatus)
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

func (u UpdateMovie) UpdateStatusError(updateStatusErrorDTO *UpdateStatusErrorDTO) error {
	movieStatus, findMovieErr := u.MovieStatusRepository.Find(updateStatusErrorDTO.MovieID)
	if findMovieErr != nil {
		return findMovieErr
	}

	changeStatusErr := movieStatus.Error()
	if changeStatusErr != nil {
		return changeStatusErr
	}

	updateMovieErr := u.MovieStatusRepository.Save(movieStatus)
	if updateMovieErr != nil {
		return updateMovieErr
	}

	return nil
}

type UpdateStatusErrorDTO struct {
	MovieID model.MovieID
}

func NewUpdateStatusErrorDTO(movieId model.MovieID) *UpdateStatusErrorDTO {
	return &UpdateStatusErrorDTO{
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
