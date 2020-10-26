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

func TestEvaluateMovie_MovieRepository_FindById_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []usecase.EvaluateMovieDTO{
		usecase.EvaluateMovieDTO{
			UserID:     model.UserID(10),
			MovieID:    model.MovieID(20),
			Evaluation: model.Evaluation(0),
		},
	}

	for _, Case := range cases {
		movieRepository := mock_repository.NewMockMovieRepository(ctrl)
		movieEvaluateRepository := mock_repository.NewMockMovieEvaluationRepository(ctrl)
		evaluateMovieUsecase := usecase.NewEvaluateUsecase(movieRepository, movieEvaluateRepository)

		movieRepository.EXPECT().FindById(Case.MovieID).Return(&model.Movie{
			ID: Case.MovieID,
		}, errors.New("ERROR"))

		result := evaluateMovieUsecase.EvaluateMovie(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}

	for _, Case := range cases {
		movieRepository := mock_repository.NewMockMovieRepository(ctrl)
		movieEvaluateRepository := mock_repository.NewMockMovieEvaluationRepository(ctrl)
		evaluateMovieUsecase := usecase.NewEvaluateUsecase(movieRepository, movieEvaluateRepository)

		movieRepository.EXPECT().FindById(Case.MovieID).Return(nil, nil)

		result := evaluateMovieUsecase.EvaluateMovie(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestEvaluateMovie_MovieEvaluationRepository_FindByUserIdAndMovieId_Nil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []usecase.EvaluateMovieDTO{
		usecase.EvaluateMovieDTO{
			UserID:     model.UserID(10),
			MovieID:    model.MovieID(20),
			Evaluation: model.Evaluation(0),
		},
	}

	for _, Case := range cases {
		movieRepository := mock_repository.NewMockMovieRepository(ctrl)
		movieEvaluateRepository := mock_repository.NewMockMovieEvaluationRepository(ctrl)
		evaluateMovieUsecase := usecase.NewEvaluateUsecase(movieRepository, movieEvaluateRepository)

		movieRepository.EXPECT().FindById(Case.MovieID).Return(&model.Movie{
			ID: Case.MovieID,
		}, nil)

		movieEvaluateRepository.EXPECT().FindByUserIdAndMovieId(Case.UserID, Case.MovieID).Return(nil)

		movieEvaluateRepository.EXPECT().Save(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(model.MovieEvaluation{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*model.MovieEvaluation).UserID != Case.UserID {
				t.Fatal("UserID Not Match,")
			}
			if data.(*model.MovieEvaluation).MovieID != Case.MovieID {
				t.Fatal("MovieID Not Match,")
			}
			if data.(*model.MovieEvaluation).Evaluation != Case.Evaluation {
				t.Fatal("Evaluation Not Match,")
			}
			return nil
		})

		result := evaluateMovieUsecase.EvaluateMovie(&Case)
		if result != nil {
			t.Fatal("Usecase Error.")
		}
	}

	for _, Case := range cases {
		movieRepository := mock_repository.NewMockMovieRepository(ctrl)
		movieEvaluateRepository := mock_repository.NewMockMovieEvaluationRepository(ctrl)
		evaluateMovieUsecase := usecase.NewEvaluateUsecase(movieRepository, movieEvaluateRepository)

		movieRepository.EXPECT().FindById(Case.MovieID).Return(&model.Movie{
			ID: Case.MovieID,
		}, nil)

		movieEvaluateRepository.EXPECT().FindByUserIdAndMovieId(Case.UserID, Case.MovieID).Return(nil)

		movieEvaluateRepository.EXPECT().Save(gomock.Any()).Return(errors.New("ERROR"))

		result := evaluateMovieUsecase.EvaluateMovie(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestEvaluateMovie_MovieEvaluationRepository_Save_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []usecase.EvaluateMovieDTO{
		usecase.EvaluateMovieDTO{
			UserID:     model.UserID(10),
			MovieID:    model.MovieID(20),
			Evaluation: model.Evaluation(0),
		},
	}

	for _, Case := range cases {
		movieRepository := mock_repository.NewMockMovieRepository(ctrl)
		movieEvaluateRepository := mock_repository.NewMockMovieEvaluationRepository(ctrl)
		evaluateMovieUsecase := usecase.NewEvaluateUsecase(movieRepository, movieEvaluateRepository)

		movieRepository.EXPECT().FindById(Case.MovieID).Return(&model.Movie{
			ID: Case.MovieID,
		}, nil)

		movieEvaluateRepository.EXPECT().FindByUserIdAndMovieId(Case.UserID, Case.MovieID).Return(&model.MovieEvaluation{
			UserID:  Case.UserID,
			MovieID: Case.MovieID,
		})

		movieEvaluateRepository.EXPECT().Save(gomock.Any()).Return(errors.New("ERROR"))

		result := evaluateMovieUsecase.EvaluateMovie(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}
