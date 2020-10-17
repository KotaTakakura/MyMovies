package infra

import (
	"MyPIPE/domain/model"
	"github.com/jinzhu/gorm"
)

type MovieEvaluatePersistence struct{
	DatabaseAccessor *gorm.DB
}

func NewMovieEvaluatePersistence()*MovieEvaluatePersistence{
	return &MovieEvaluatePersistence{
		DatabaseAccessor:	ConnectGorm(),
	}
}

func (m MovieEvaluatePersistence) FindByUserIdAndMovieId(userId model.UserID, movieId model.MovieID) *model.MovieEvaluation {
	var evaluation model.MovieEvaluation
	var count int
	m.DatabaseAccessor.Where("movie_id = ? and user_id = ?",movieId,userId).Take(&evaluation).Count(&count)
	if count == 0{
		return nil
	}
	return &evaluation
}

func (m MovieEvaluatePersistence) FindByUserIdAndMovieIdAndEvaluation(userId model.UserID, movieId model.MovieID,evaluation model.Evaluation) *model.MovieEvaluation {
	var movieEvaluation model.MovieEvaluation
	var count int
	m.DatabaseAccessor.Where("movie_id = ? and user_id = ? and evaluation = ?",movieId,userId,evaluation).Take(&movieEvaluation).Count(&count)
	if count == 0{
		return nil
	}
	return &movieEvaluation
}

func (m MovieEvaluatePersistence) Save(evaluation *model.MovieEvaluation)error{
	var alreadyExistsEvaluation model.MovieEvaluation
	result := m.DatabaseAccessor.Where("movie_id = ? and user_id = ?",evaluation.MovieID,evaluation.UserID).Take(&alreadyExistsEvaluation)
	if result.RowsAffected == 0{
		m.DatabaseAccessor.Create(&evaluation)
		return nil
	}
	result = m.DatabaseAccessor.
		Model(&evaluation).
		Table("movie_evaluations").
		Where("movie_id = ? and user_id = ?",evaluation.MovieID,evaluation.UserID).
		Update("evaluation",evaluation.Evaluation)
	if result.Error != nil{
		return result.Error
	}
	return nil
}