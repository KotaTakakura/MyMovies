package infra

import (
	"MyPIPE/domain/model"
)

type MovieEvaluatePersistence struct{}

func NewMovieEvaluatePersistence() *MovieEvaluatePersistence {
	return &MovieEvaluatePersistence{}
}

func (m MovieEvaluatePersistence) FindByUserIdAndMovieId(userId model.UserID, movieId model.MovieID) *model.MovieEvaluation {
	db := ConnectGorm()
	defer db.Close()
	var evaluation model.MovieEvaluation
	var count int
	db.Where("movie_id = ? and user_id = ?", movieId, userId).Take(&evaluation).Count(&count)
	if count == 0 {
		return nil
	}
	return &evaluation
}

func (m MovieEvaluatePersistence) FindByUserIdAndMovieIdAndEvaluation(userId model.UserID, movieId model.MovieID, evaluation model.Evaluation) *model.MovieEvaluation {
	db := ConnectGorm()
	defer db.Close()
	var movieEvaluation model.MovieEvaluation
	var count int
	db.Where("movie_id = ? and user_id = ? and evaluation = ?", movieId, userId, evaluation).Take(&movieEvaluation).Count(&count)
	if count == 0 {
		return nil
	}
	return &movieEvaluation
}

func (m MovieEvaluatePersistence) Save(evaluation *model.MovieEvaluation) error {
	db := ConnectGorm()
	defer db.Close()
	var alreadyExistsEvaluation model.MovieEvaluation
	result := db.Where("movie_id = ? and user_id = ?", evaluation.MovieID, evaluation.UserID).Take(&alreadyExistsEvaluation)
	if result.RowsAffected == 0 {
		db.Create(&evaluation)
		return nil
	}
	result = db.
		Model(&evaluation).
		Table("movie_evaluations").
		Where("movie_id = ? and user_id = ?", evaluation.MovieID, evaluation.UserID).
		Update("evaluation", evaluation.Evaluation)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
