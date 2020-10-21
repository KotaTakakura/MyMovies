package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/factory"
	"MyPIPE/domain/repository"
	"mime/multipart"
)

type IPostMovie interface {
	PostMovie(postMovieDTO *PostMovieDTO)(*model.Movie,error)
}

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

func (p *PostMovie)PostMovie(postMovieDTO *PostMovieDTO)(*model.Movie,error){
	newMovie,createError := p.MovieModelFactory.CreateMovieModel(postMovieDTO.UserID,postMovieDTO.FileHeader,postMovieDTO.ThumbnailHeader)
	if createError != nil{
		return nil,createError
	}

	savedNewMovie,saveError := p.MovieRepository.Save(*newMovie)
	if saveError != nil{
		return nil,saveError
	}

	err := p.FileUploadRepository.Upload(postMovieDTO.File,postMovieDTO.FileHeader,savedNewMovie.ID)
	if err != nil{
		return nil,err
	}

	thumbnailUploadErr := p.ThumbnailUploadRepository.Upload(postMovieDTO.Thumbnail,*savedNewMovie)
	if thumbnailUploadErr != nil{
		return nil,thumbnailUploadErr
	}

	return savedNewMovie,nil
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
