package model

import (
	"errors"
)

type Evaluation uint64

func NewEvaluation(evaluation string) (Evaluation, error) {
	switch evaluation {
	case "good":
		return Evaluation(0), nil
	case "bad":
		return Evaluation(1), nil
	case "default":
		return Evaluation(2), nil
	}
	return Evaluation(999), errors.New("Invalid Evaluation.")
}

type MovieEvaluation struct {
	UserID     UserID     `gorm:"column:user_id"`
	MovieID    MovieID    `gorm:"column:movie_id"`
	Evaluation Evaluation `gorm:"column:evaluation"`
}

func NewMovieEvaluation(userId UserID, movieId MovieID, evaluation Evaluation) *MovieEvaluation {
	return &MovieEvaluation{
		UserID:     userId,
		MovieID:    movieId,
		Evaluation: evaluation,
	}
}

func (m *MovieEvaluation) EvaluateMovie(evaluation Evaluation) error {
	m.Evaluation = evaluation
	return nil
}
