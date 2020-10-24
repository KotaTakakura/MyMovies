package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestEvaluateMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	movieEvaluateRepository := mock_repository.NewMockMovieEvaluationRepository(ctrl)
	evaluateMovieUsecase := usecase.NewEvaluateUsecase(movieRepository, movieEvaluateRepository)

	trueCases := []usecase.EvaluateMovieDTO{
		usecase.EvaluateMovieDTO{
			UserID:     model.UserID(10),
			MovieID:    model.MovieID(20),
			Evaluation: model.Evaluation(0),
		},
	}

	for _, trueCase := range trueCases {
		movieRepository.EXPECT().FindById(trueCase.MovieID).Return(&model.Movie{
			ID: trueCase.MovieID,
		}, nil)

		movieEvaluateRepository.EXPECT().FindByUserIdAndMovieId(trueCase.UserID, trueCase.MovieID).Return(&model.MovieEvaluation{
			UserID:  trueCase.UserID,
			MovieID: trueCase.MovieID,
		})

		movieEvaluateRepository.EXPECT().Save(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&model.MovieEvaluation{}) {
				t.Fatal("Type Not Match.")
			}
			if data.(*model.MovieEvaluation).UserID != trueCase.UserID {
				t.Fatal("UserID Not Match.")
			}
			if data.(*model.MovieEvaluation).Evaluation != trueCase.Evaluation {
				t.Fatal("Evaluation Not Match.")
			}
			return nil
		})

		result := evaluateMovieUsecase.EvaluateMovie(&trueCase)
		if result != nil {
			t.Fatal("Usecase Error.")
		}
	}
}
