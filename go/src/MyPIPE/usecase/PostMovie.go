package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/factory"
	"MyPIPE/domain/repository"
	"mime/multipart"
)

type PostMovie struct{
	FileUploadRepository	repository.FileUpload
	ThumbnailUploadRepository	repository.ThumbnailUploadRepository
	MovieRepository	repository.MovieRepository
	MovieModelFactory	factory.IMovieModelFactory
}

func NewPostMovie(fr repository.FileUpload,tu repository.ThumbnailUploadRepository,mr repository.MovieRepository,mf factory.IMovieModelFactory)*PostMovie{
	return &PostMovie{
		FileUploadRepository: fr,
		ThumbnailUploadRepository: tu,
		MovieRepository: mr,
		MovieModelFactory: mf,
	}
}

func (p *PostMovie)PostMovie(postMovieDTO *PostMovieDTO)error{
	newMovie,createError := p.MovieModelFactory.CreateMovieModel(postMovieDTO.UserID,postMovieDTO.FileHeader,postMovieDTO.ThumbnailHeader)
	if createError != nil{
		return createError
	}

	savedNewMovie,saveError := p.MovieRepository.Save(*newMovie)
	if saveError != nil{
		return saveError
	}

	err := p.FileUploadRepository.Upload(postMovieDTO.File,postMovieDTO.FileHeader,savedNewMovie.ID)
	if err != nil{
		return err
	}

	thumbnailUploadErr := p.ThumbnailUploadRepository.Upload(postMovieDTO.Thumbnail,postMovieDTO.ThumbnailHeader,savedNewMovie.ID)
	if thumbnailUploadErr != nil{
		return thumbnailUploadErr
	}

	return nil
}

type PostMovieDTO struct{
	File	multipart.File
	FileHeader	multipart.FileHeader
	Thumbnail	multipart.File
	ThumbnailHeader	multipart.FileHeader
	UserID model.UserID
}

func NewPostMovieDTO(file multipart.File,file_header multipart.FileHeader,thumbnail multipart.File,thumbnailHeader multipart.FileHeader,userId model.UserID)*PostMovieDTO{
	return &PostMovieDTO{
		File:	file,
		FileHeader: file_header,
		Thumbnail: thumbnail,
		ThumbnailHeader: thumbnailHeader,
		UserID:	userId,
	}
}
