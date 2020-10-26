package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"errors"
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

func TestDeletePlayListMovie_PlayListRepository_FindByID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []usecase.DeletePlayListMovieDTO{
		usecase.DeletePlayListMovieDTO{
			UserID:     model.UserID(10),
			PlayListID: model.PlayListID(100),
			MovieID:    model.MovieID(200),
		},
	}

	for _, Case := range cases {
		playListRepository := mock_repository.NewMockPlayListRepository(ctrl)
		playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
		deletePlayListMovieUsecase := usecase.NewDeletePlayListMovie(playListRepository, playListMovieRepository)

		playListRepository.EXPECT().FindByID(Case.PlayListID).Return(nil, nil)

		result := deletePlayListMovieUsecase.DeletePlayListItem(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}

	for _, Case := range cases {
		playListRepository := mock_repository.NewMockPlayListRepository(ctrl)
		playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
		deletePlayListMovieUsecase := usecase.NewDeletePlayListMovie(playListRepository, playListMovieRepository)

		playListRepository.EXPECT().FindByID(Case.PlayListID).Return(&model.PlayList{
			ID:     Case.PlayListID,
			UserID: Case.UserID + 1,
		}, nil)

		result := deletePlayListMovieUsecase.DeletePlayListItem(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}

	for _, Case := range cases {
		playListRepository := mock_repository.NewMockPlayListRepository(ctrl)
		playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
		deletePlayListMovieUsecase := usecase.NewDeletePlayListMovie(playListRepository, playListMovieRepository)

		playListRepository.EXPECT().FindByID(Case.PlayListID).Return(&model.PlayList{
			ID:     Case.PlayListID,
			UserID: Case.UserID,
		}, errors.New("ERROR"))

		result := deletePlayListMovieUsecase.DeletePlayListItem(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestDeletePlayListMovie_PlayListMovieRepository_Find_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []usecase.DeletePlayListMovieDTO{
		usecase.DeletePlayListMovieDTO{
			UserID:     model.UserID(10),
			PlayListID: model.PlayListID(100),
			MovieID:    model.MovieID(200),
		},
	}

	for _, Case := range cases {
		playListRepository := mock_repository.NewMockPlayListRepository(ctrl)
		playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
		deletePlayListMovieUsecase := usecase.NewDeletePlayListMovie(playListRepository, playListMovieRepository)

		playListRepository.EXPECT().FindByID(Case.PlayListID).Return(&model.PlayList{
			ID:     Case.PlayListID,
			UserID: Case.UserID,
		}, nil)

		playListMovieRepository.EXPECT().Find(Case.UserID, Case.PlayListID, Case.MovieID).Return(nil)

		result := deletePlayListMovieUsecase.DeletePlayListItem(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}

	for _, Case := range cases {
		playListRepository := mock_repository.NewMockPlayListRepository(ctrl)
		playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
		deletePlayListMovieUsecase := usecase.NewDeletePlayListMovie(playListRepository, playListMovieRepository)

		playListRepository.EXPECT().FindByID(Case.PlayListID).Return(&model.PlayList{
			ID:     Case.PlayListID,
			UserID: Case.UserID,
		}, nil)

		playListMovieRepository.EXPECT().Find(Case.UserID, Case.PlayListID, Case.MovieID).Return(&model.PlayListMovie{
			PlayListID: Case.PlayListID,
			MovieID:    Case.MovieID,
		})

		playListMovieRepository.EXPECT().Remove(gomock.Any()).Return(errors.New("ERROR"))

		result := deletePlayListMovieUsecase.DeletePlayListItem(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}
