package usecase

import (
	"MyPIPE/domain/factory"
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type IPostMovie interface {
	PostMovie(postMovieDTO *PostMovieDTO) (*model.Movie, error)
}

type PostMovie struct {
	FileUploadRepository      repository.FileUpload
	ThumbnailUploadRepository repository.ThumbnailUploadRepository
	MovieRepository           repository.MovieRepository
	MovieModelFactory         factory.IMovieModelFactory
}

func NewPostMovie(fr repository.FileUpload, tu repository.ThumbnailUploadRepository, mr repository.MovieRepository, mf factory.IMovieModelFactory) *PostMovie {
	return &PostMovie{
		FileUploadRepository:      fr,
		ThumbnailUploadRepository: tu,
		MovieRepository:           mr,
		MovieModelFactory:         mf,
	}
}

func (p *PostMovie) PostMovie(postMovieDTO *PostMovieDTO) (*model.Movie, error) {
	newMovie, createError := p.MovieModelFactory.CreateMovieModel(postMovieDTO.UserID, postMovieDTO.MovieFile, postMovieDTO.Thumbnail)
	if createError != nil {
		return nil, createError
	}

	savedNewMovie, saveError := p.MovieRepository.Save(*newMovie)
	if saveError != nil {
		return nil, saveError
	}

	err := p.FileUploadRepository.Upload(postMovieDTO.MovieFile.File, postMovieDTO.MovieFile.FileHeader, savedNewMovie.ID)
	if err != nil {
		return nil, err
	}

	thumbnailUploadErr := p.ThumbnailUploadRepository.Upload(postMovieDTO.Thumbnail.File, *savedNewMovie)
	if thumbnailUploadErr != nil {
		return nil, thumbnailUploadErr
	}

	return savedNewMovie, nil
}

type PostMovieDTO struct {
	MovieFile *model.MovieFile
	Thumbnail *model.MovieThumbnail
	UserID    model.UserID
}

func NewPostMovieDTO(movieFile *model.MovieFile, thumbnail *model.MovieThumbnail, userId model.UserID) *PostMovieDTO {
	return &PostMovieDTO{
		MovieFile: movieFile,
		Thumbnail: thumbnail,
		UserID:    userId,
	}
}
