package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestEvaluateMovie(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := usecase.EvaluateMovieDTO{
		UserID:     model.UserID(10),
		MovieID:    model.MovieID(100),
		Evaluation: model.Evaluation(1),
	}

	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	movieEvaluationRepository := mock_repository.NewMockMovieEvaluationRepository(ctrl)

	movieRepository.EXPECT().FindById(cases.MovieID).Return(&model.Movie{
		ID:          cases.MovieID,
		StoreName:   "TestStoreName",
		DisplayName: "TestDisplayName",
		UserID:      cases.UserID,
	}, nil)

	movieEvaluationRepository.EXPECT().FindByUserIdAndMovieId(cases.UserID, cases.MovieID).Return(&model.MovieEvaluation{
		UserID:     cases.UserID,
		MovieID:    cases.MovieID,
		Evaluation: cases.Evaluation,
	})

	movieEvaluationRepository.EXPECT().Save(gomock.Any()).DoAndReturn(
		func(movieEvaluation *model.MovieEvaluation) error {
			if movieEvaluation.UserID != cases.UserID {
				t.Error("Invalid UserID.")
			}
			if movieEvaluation.MovieID != cases.MovieID {
				t.Error("Invalid MovieID.")
			}
			if movieEvaluation.Evaluation != cases.Evaluation {
				t.Error("Invalid Evaluation.")
			}
			return nil
		})

	movieEvaluationUsecase := usecase.NewEvaluateUsecase(movieRepository, movieEvaluationRepository)
	err := movieEvaluationUsecase.EvaluateMovie(cases)
	if err != nil {
		t.Error("Usecase Error.")
	}
}
