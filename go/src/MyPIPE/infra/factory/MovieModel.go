package factory

import (
	"MyPIPE/domain/model"
	"mime/multipart"
	"path/filepath"
)

type MovieModelFactory struct {}

func NewMovieModelFactory()*MovieModelFactory{
	return &MovieModelFactory{}
}

func (m MovieModelFactory)CreateMovieModel(uploaderID model.UserID,displayName model.MovieDisplayName,fileHeader multipart.FileHeader)(*model.Movie,error){
	storeName,storeNameErr := model.NewMovieStoreName(filepath.Ext(fileHeader.Filename))
	if storeNameErr != nil{
		return nil,storeNameErr
	}
	return model.NewMovie(uploaderID,storeName,displayName),nil
}