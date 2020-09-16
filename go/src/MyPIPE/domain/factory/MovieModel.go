package factory

import (
	"MyPIPE/domain/model"
	"github.com/google/uuid"
)

type MovieModelFactory struct {}

func NewMovieModelFactory()*MovieModelFactory{
	return &MovieModelFactory{}
}

func (m MovieModelFactory)CreateMovieModel(uploaderID model.UserID,displayName model.MovieDisplayName,fileExtension string)(*model.Movie,error){
	newMovieUUID,uuidErr := uuid.NewUUID()
	if uuidErr != nil{
		return nil,uuidErr
	}
	newMovieID,movieIdErr := model.NewMovieID(newMovieUUID.String())
	if movieIdErr != nil{
		return nil,movieIdErr
	}

	storeName,storeNameErr := model.NewMovieStoreName(newMovieUUID.String() + fileExtension)
	if storeNameErr != nil{
		return nil,storeNameErr
	}
	return model.NewMovie(newMovieID,uploaderID,storeName,displayName),nil
}