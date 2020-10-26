package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCheckUserAlreadyLikedMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieEvaluationRepository := mock_repository.NewMockMovieEvaluationRepository(ctrl)
	checkUserAlreadyLikedMovieUsecase := usecase.NewCheckUserAlreadyLikedMovie(movieEvaluationRepository)

	trueCases := []usecase.CheckUserAlreadyLikedMovieFindDTO{
		usecase.CheckUserAlreadyLikedMovieFindDTO{
			UserID:  model.UserID(10),
			MovieID: model.MovieID(100),
		},
	}

	for _, trueCase := range trueCases {
		movieEvaluationRepository.EXPECT().FindByUserIdAndMovieIdAndEvaluation(trueCase.UserID, trueCase.MovieID, model.Evaluation(0)).Return(
			&model.MovieEvaluation{
				UserID:     trueCase.UserID,
				MovieID:    trueCase.MovieID,
				Evaluation: model.Evaluation(0),
			},
		)

		result := checkUserAlreadyLikedMovieUsecase.Find(&trueCase)
		if !result {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestCheckUserAlreadyLikedMovie_MovieEvaluationRepository_FindByUserIdAndMovieIdAndEvaluation_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieEvaluationRepository := mock_repository.NewMockMovieEvaluationRepository(ctrl)
	checkUserAlreadyLikedMovieUsecase := usecase.NewCheckUserAlreadyLikedMovie(movieEvaluationRepository)

	trueCases := []usecase.CheckUserAlreadyLikedMovieFindDTO{
		usecase.CheckUserAlreadyLikedMovieFindDTO{
			UserID:  model.UserID(10),
			MovieID: model.MovieID(100),
		},
	}

	for _, trueCase := range trueCases {
		movieEvaluationRepository.EXPECT().FindByUserIdAndMovieIdAndEvaluation(trueCase.UserID, trueCase.MovieID, model.Evaluation(0)).Return(nil)

		result := checkUserAlreadyLikedMovieUsecase.Find(&trueCase)
		if result {
			t.Fatal("Usecase Error.")
		}
	}
}
