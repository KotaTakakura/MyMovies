package usecase

import (
	"MyPIPE/domain/factory"
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"mime/multipart"
	"path/filepath"
)

type PostMovie struct{
	FileUploadRepository	repository.FileUpload
	MovieRepository	repository.MovieRepository
	MovieModelFactory	factory.MovieModelFactory
}

func NewPostMovie(fr repository.FileUpload,mr repository.MovieRepository,mf factory.MovieModelFactory)*PostMovie{
	return &PostMovie{
		FileUploadRepository: fr,
		MovieRepository: mr,
		MovieModelFactory: mf,
	}
}

func (p *PostMovie)PostMovie(movieFile multipart.File,postMovieDTO PostMovieDTO)error{
	uploaderID,_ := model.NewUserID(postMovieDTO.UserID)
	displayName,_ := model.NewMovieDisplayName(postMovieDTO.DisplayName)
	newMovie,createError := p.MovieModelFactory.CreateMovieModel(uploaderID,displayName,filepath.Ext(postMovieDTO.FileHeader.Filename))
	if createError != nil{
		return createError
	}

	saveError := p.MovieRepository.Save(*newMovie)
	if saveError != nil{
		return saveError
	}

	err := p.FileUploadRepository.Upload(postMovieDTO.File,postMovieDTO.FileHeader,newMovie.ID)
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
