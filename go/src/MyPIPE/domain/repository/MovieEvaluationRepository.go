package repository

import "MyPIPE/domain/model"

type MovieEvaluationRepository interface {
	FindByUserIdAndMovieId(userId model.UserID, movieId model.MovieID) *model.MovieEvaluation
	Save(evaluation *model.MovieEvaluation) error
	FindByUserIdAndMovieIdAndEvaluation(userId model.UserID, movieId model.MovieID, evaluation model.Evaluation) *model.MovieEvaluation
}
