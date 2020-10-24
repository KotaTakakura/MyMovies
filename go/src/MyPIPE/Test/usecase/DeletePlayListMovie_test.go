package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestDeletePlayListMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListRepository := mock_repository.NewMockPlayListRepository(ctrl)
	playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
	deletePlayListMovieUsecase := usecase.NewDeletePlayListMovie(playListRepository, playListMovieRepository)

	trueCases := []usecase.DeletePlayListMovieDTO{
		usecase.DeletePlayListMovieDTO{
			UserID:     model.UserID(10),
			PlayListID: model.PlayListID(100),
			MovieID:    model.MovieID(200),
		},
	}

	for _, trueCase := range trueCases {
		playListRepository.EXPECT().FindByID(trueCase.PlayListID).Return(&model.PlayList{
			ID:     trueCase.PlayListID,
			UserID: trueCase.UserID,
		}, nil)

		playListMovieRepository.EXPECT().Find(trueCase.UserID, trueCase.PlayListID, trueCase.MovieID).Return(&model.PlayListMovie{
			PlayListID: trueCase.PlayListID,
			MovieID:    trueCase.MovieID,
		})

		playListMovieRepository.EXPECT().Remove(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&model.PlayListMovie{}) {
				t.Fatal("Type Not Match.")
			}
			if data.(*model.PlayListMovie).PlayListID != trueCase.PlayListID {
				t.Fatal("PlayListID Not Match.")
			}
			if data.(*model.PlayListMovie).MovieID != trueCase.MovieID {
				t.Fatal("MovieID Not Match.")
			}
			return nil
		})

		result := deletePlayListMovieUsecase.DeletePlayListItem(&trueCase)
		if result != nil {
			t.Fatal("Usecase Error.")
		}
	}
}
