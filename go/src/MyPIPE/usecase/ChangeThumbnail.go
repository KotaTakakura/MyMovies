package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type IChangeThumbnail interface {
	ChangeThumbnail(changeThumbnailDTO *ChangeThumbnailDTO) error
}

type ChangeThumbnail struct {
	MovieRepository           repository.MovieRepository
	ThumbnailUploadRepository repository.ThumbnailUploadRepository
}

func NewChangeThumbnail(m repository.MovieRepository, t repository.ThumbnailUploadRepository) *ChangeThumbnail {
	return &ChangeThumbnail{
		MovieRepository:           m,
		ThumbnailUploadRepository: t,
	}
}

func (c ChangeThumbnail) ChangeThumbnail(changeThumbnailDTO *ChangeThumbnailDTO) error {
	movie, findMovieErr := c.MovieRepository.FindByUserIdAndMovieId(changeThumbnailDTO.UserID, changeThumbnailDTO.MovieID)
	if findMovieErr != nil {
		return findMovieErr
	}

	changeThumbnailNameErr := movie.ChangeThumbnailName(changeThumbnailDTO.Thumbnail)
	if changeThumbnailNameErr != nil {
		return changeThumbnailNameErr
	}

	_, updateErr := c.MovieRepository.Update(*movie)
	if updateErr != nil {
		return updateErr
	}

	thumbnailUploadErr := c.ThumbnailUploadRepository.Upload(changeThumbnailDTO.Thumbnail.File, *movie)
	if thumbnailUploadErr != nil {
		return thumbnailUploadErr
	}

	return nil
}

type ChangeThumbnailDTO struct {
	UserID          model.UserID
	MovieID         model.MovieID
	Thumbnail       model.MovieThumbnail
}

func NewChangeThumbnailDTO(userId model.UserID, movieId model.MovieID, thumbnail model.MovieThumbnail) *ChangeThumbnailDTO {
	return &ChangeThumbnailDTO{
		UserID:          userId,
		MovieID:         movieId,
		Thumbnail:       thumbnail,
	}
}
