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

func (m MovieModelFactory)CreateMovieModel(uploaderID model.UserID,fileHeader multipart.FileHeader,thumbnailHeader multipart.FileHeader)(*model.Movie,error){
	storeName,storeNameErr := model.NewMovieStoreName(filepath.Ext(fileHeader.Filename))
	if storeNameErr != nil{
		return nil,storeNameErr
	}

	displayName,displayNameErr := model.NewMovieDisplayName("")
	if displayNameErr != nil{
		return nil,displayNameErr
	}

	thumbnailName,thumbnailNameErr := model.NewMovieThumbnailName(thumbnailHeader)
	if thumbnailNameErr != nil{
		return nil,thumbnailNameErr
	}

	return model.NewMovie(uploaderID,storeName,displayName,thumbnailName),nil
}