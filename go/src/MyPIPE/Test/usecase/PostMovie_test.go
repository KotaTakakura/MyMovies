package test

import (
	mock_factory "MyPIPE/Test/mock/factory"
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestPostMovie(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postMovieDTO := usecase.PostMovieDTO{
		UserID:	model.UserID(10),
		DisplayName: model.MovieDisplayName("display_name"),
	}

	MovieRepository := mock_repository.NewMockMovieRepository(ctrl)
	FileRepository := mock_repository.NewMockFileUpload(ctrl)
	MovieFactory := mock_factory.NewMockIMovieModelFactory(ctrl)

	MovieFactory.EXPECT().
		CreateMovieModel(postMovieDTO.UserID,postMovieDTO.DisplayName,postMovieDTO.FileHeader).
		Return(&model.Movie{
			DisplayName: postMovieDTO.DisplayName,
			UserID: postMovieDTO.UserID,
	},nil)

	MovieRepository.EXPECT().Save(gomock.Any()).DoAndReturn(func(movie model.Movie)(*model.Movie,error){
		if movie.DisplayName != postMovieDTO.DisplayName{
			t.Error("Invalid Display Name.")
			return nil,nil
		}
		if movie.UserID != postMovieDTO.UserID{
			t.Error("Invalid UserID Name.")
			return nil,nil
		}

			return &model.Movie{},nil
	})

	FileRepository.EXPECT().Upload(postMovieDTO.File,postMovieDTO.FileHeader,model.Movie{}.ID).Return(nil)

	postMovieUsecase := usecase.NewPostMovie(FileRepository,MovieRepository,MovieFactory)
	result := postMovieUsecase.PostMovie(postMovieDTO)
	if result != nil{
		t.Error("PostMovie Usecase Error.")
	}
}
