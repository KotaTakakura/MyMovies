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

func TestChangeOrderOfPlayListMovies(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
	changeOrderOfPlayListMoviesUsecase := usecase.NewChangeOrderOfPlayListMovies(playListMovieRepository)

	trueCases := []usecase.ChangeOrderOfPlayListMoviesDTO{
		usecase.ChangeOrderOfPlayListMoviesDTO{
			UserID:     model.UserID(20),
			PlayListID: model.PlayListID(10),
			MovieIDAndOrder: []usecase.MovieIdAndOrderForChangeOrderOfPlayListMoviesDTO{
				usecase.MovieIdAndOrderForChangeOrderOfPlayListMoviesDTO{
					MovieID: model.MovieID(100),
					Order:   model.PlayListMovieOrder(5),
				},
			},
		},
	}

	for _, trueCase := range trueCases {
		playListMovieRepository.EXPECT().FindAll(trueCase.UserID, trueCase.PlayListID).Return([]model.PlayListMovie{
			model.PlayListMovie{
				PlayListID: trueCase.PlayListID,
				MovieID:    trueCase.MovieIDAndOrder[0].MovieID,
			},
		})
		playListMovieRepository.EXPECT().SaveAll(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf([]model.PlayListMovie{}) {
				t.Fatal("Type Not Match.")
			}
			return nil
		})

		result := changeOrderOfPlayListMoviesUsecase.ChangeOrderOfPlayListMovies(&trueCase)
		if result != nil {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestChangeOrderOfPlayListMovies_PlayListMovieRepository_SaveAll_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []usecase.ChangeOrderOfPlayListMoviesDTO{
		usecase.ChangeOrderOfPlayListMoviesDTO{
			UserID:     model.UserID(20),
			PlayListID: model.PlayListID(10),
			MovieIDAndOrder: []usecase.MovieIdAndOrderForChangeOrderOfPlayListMoviesDTO{
				usecase.MovieIdAndOrderForChangeOrderOfPlayListMoviesDTO{
					MovieID: model.MovieID(100),
					Order:   model.PlayListMovieOrder(5),
				},
			},
		},
	}

	for _, Case := range cases {
		playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
		changeOrderOfPlayListMoviesUsecase := usecase.NewChangeOrderOfPlayListMovies(playListMovieRepository)

		playListMovieRepository.EXPECT().FindAll(Case.UserID, Case.PlayListID).Return([]model.PlayListMovie{
			model.PlayListMovie{
				PlayListID: Case.PlayListID,
				MovieID:    Case.MovieIDAndOrder[0].MovieID,
			},
		})
		playListMovieRepository.EXPECT().SaveAll(gomock.Any()).Return(errors.New("ERROR"))

		result := changeOrderOfPlayListMoviesUsecase.ChangeOrderOfPlayListMovies(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}
