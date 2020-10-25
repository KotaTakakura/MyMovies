package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestUpdateMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	updateMovieUsecase := usecase.NewUpdateMovie(movieRepository)

	trueCases := []usecase.UpdateDTO{
		usecase.UpdateDTO{
			UserID:      model.UserID(10),
			MovieID:     model.MovieID(20),
			DisplayName: model.MovieDisplayName("TestNewDisplayName"),
			Description: model.MovieDescription("TestNewDescription"),
			Public:      model.MoviePublic(0),
			Status:      model.MovieStatus(0),
		},
	}

	for _, trueCase := range trueCases {
		movieRepository.EXPECT().FindByUserIdAndMovieId(trueCase.UserID, trueCase.MovieID).Return(&model.Movie{
			ID:            trueCase.MovieID,
			StoreName:     "TestStoreName",
			DisplayName:   "OldDisplayName",
			Description:   "OldDescription",
			ThumbnailName: "TestThumbnailname",
			UserID:        trueCase.UserID,
			Public:        model.MoviePublic(1),
			Status:        model.MovieStatus(1),
		}, nil)

		movieRepository.EXPECT().Update(gomock.Any()).DoAndReturn(func(data interface{}) (*model.Movie, error) {
			if reflect.TypeOf(data) != reflect.TypeOf(model.Movie{}) {
				t.Fatal("Type Not Match.")
			}
			if data.(model.Movie).ID != trueCase.MovieID {
				t.Fatal("MovieID Not Match.")
			}
			if data.(model.Movie).DisplayName != trueCase.DisplayName {
				t.Fatal("DisplayName Not Match.")
			}
			if data.(model.Movie).Description != trueCase.Description {
				t.Fatal("Description Not Match.")
			}
			if data.(model.Movie).Public != trueCase.Public {
				t.Fatal("Public Not Match.")
			}
			if data.(model.Movie).Status != trueCase.Status {
				t.Fatal("Status Not Match.")
			}
			updatedMovie := data.(model.Movie)
			return &updatedMovie, nil
		})

		result, err := updateMovieUsecase.Update(&trueCase)
		if err != nil {
			t.Fatal("Status Not Match.")
		}
		if result.ID != trueCase.MovieID {
			t.Fatal("MovieID Not Match.")
		}
		if result.DisplayName != trueCase.DisplayName {
			t.Fatal("DisplayName Not Match.")
		}
		if result.Description != trueCase.Description {
			t.Fatal("Description Not Match.")
		}
		if result.Public != trueCase.Public {
			t.Fatal("Public Not Match.")
		}
		if result.Status != trueCase.Status {
			t.Fatal("Status Not Match.")
		}
	}
}
