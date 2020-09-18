package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestPostMovie(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MovieRepository := mock_repository.NewMockMovieRepository(ctrl)
	FileRepository := mock_repository.NewMockFileUpload(ctrl)
	MovieRepository := mock_repository.NewMockMovieRepository(ctrl)
	postMovieUsecase := usecase.NewPostMovie(FileRepository,MovieRepository)

}
