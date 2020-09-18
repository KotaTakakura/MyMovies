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
		UserID:	uint64(10),
		DisplayName: string("display_name"),
	}

	MovieRepository := mock_repository.NewMockMovieRepository(ctrl)
	FileRepository := mock_repository.NewMockFileUpload(ctrl)
	MovieFactory := mock_factory.NewMockIMovieModelFactory(ctrl)

	uploaderID,_ := model.NewUserID(postMovieDTO.UserID)
	displayName,_ := model.NewMovieDisplayName(postMovieDTO.DisplayName)

	MovieFactory.EXPECT().
		CreateMovieModel(uploaderID,displayName,postMovieDTO.FileHeader).
		Return(&model.Movie{},nil)

	MovieRepository.EXPECT().Save(model.Movie{}).Return(&model.Movie{},nil)

	FileRepository.EXPECT().Upload(postMovieDTO.File,postMovieDTO.FileHeader,model.Movie{}.ID)

	postMovieUsecase := usecase.NewPostMovie(FileRepository,MovieRepository,MovieFactory)
	result := postMovieUsecase.PostMovie(postMovieDTO)
	if result != nil{
		t.Error("errorrrr")
	}
}
