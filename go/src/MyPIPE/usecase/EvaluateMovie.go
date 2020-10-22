package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
)

type IEvaluateMovie interface {
	EvaluateMovie(evaluateMovieDTO *EvaluateMovieDTO) error
}

type EvaluateMovie struct {
	MovieRepository           repository.MovieRepository
	MovieEvaluationRepository repository.MovieEvaluationRepository
}

func NewEvaluateUsecase(m repository.MovieRepository, me repository.MovieEvaluationRepository) *EvaluateMovie {
	return &EvaluateMovie{
		MovieRepository:           m,
		MovieEvaluationRepository: me,
	}
}

func (e EvaluateMovie) EvaluateMovie(evaluateMovieDTO *EvaluateMovieDTO) error {
	movie, movieErr := e.MovieRepository.FindById(evaluateMovieDTO.MovieID)
	if movieErr != nil {
		return movieErr
	}

	if movie == nil {
		return errors.New("No Such Movie.")
	}

	movieEvaluation := e.MovieEvaluationRepository.FindByUserIdAndMovieId(evaluateMovieDTO.UserID, evaluateMovieDTO.MovieID)

	if movieEvaluation == nil {
		newMovieEvaluation := model.NewMovieEvaluation(evaluateMovieDTO.UserID, evaluateMovieDTO.MovieID, evaluateMovieDTO.Evaluation)
		evaluationSaveErr := e.MovieEvaluationRepository.Save(newMovieEvaluation)
		if evaluationSaveErr != nil {
			return evaluationSaveErr
		}
		return nil
	}

	evaluationErr := movieEvaluation.EvaluateMovie(evaluateMovieDTO.Evaluation)
	if evaluationErr != nil {
		return evaluationErr
	}

	evaluationSaveErr := e.MovieEvaluationRepository.Save(movieEvaluation)
	if evaluationSaveErr != nil {
		return evaluationSaveErr
	}
	return nil
}

type EvaluateMovieDTO struct {
	UserID     model.UserID
	MovieID    model.MovieID
	Evaluation model.Evaluation
}
