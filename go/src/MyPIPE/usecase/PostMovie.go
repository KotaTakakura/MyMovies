package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/factory"
	"MyPIPE/domain/repository"
	"mime/multipart"
)

type PostMovie struct{
	FileUploadRepository	repository.FileUpload
	MovieRepository	repository.MovieRepository
	MovieModelFactory	factory.IMovieModelFactory
}

func NewPostMovie(fr repository.FileUpload,mr repository.MovieRepository,mf factory.IMovieModelFactory)*PostMovie{
	return &PostMovie{
		FileUploadRepository: fr,
		MovieRepository: mr,
		MovieModelFactory: mf,
	}
}

func (p *PostMovie)PostMovie(postMovieDTO PostMovieDTO)error{
	uploaderID,userIdErr := model.NewUserID(postMovieDTO.UserID)
	if userIdErr != nil{
		return userIdErr
	}
	displayName,displayNameErr := model.NewMovieDisplayName(postMovieDTO.DisplayName)
	if displayNameErr != nil{
		return displayNameErr
	}
	newMovie,createError := p.MovieModelFactory.CreateMovieModel(uploaderID,displayName,postMovieDTO.FileHeader)
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

	return nil
}

type PostMovieDTO struct{
	File	multipart.File
	FileHeader	multipart.FileHeader
	UserID uint64
	DisplayName	string
}