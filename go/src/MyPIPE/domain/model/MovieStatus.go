package model

import (
	"errors"
	"time"
)

type MovieStatusModel struct{
	MovieID MovieID	`gorm:"column:id"`
	MovieStatus	MovieStatusValue	`gorm:"column:movie_status"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewMovieStatusModel(movieId MovieID,movieStatusValue MovieStatusValue)*MovieStatusModel{
	return &MovieStatusModel{
		MovieID: movieId,
		MovieStatus: movieStatusValue,
	}
}

type MovieStatusValue int

func NewMovieStatusValue(status uint) (MovieStatusValue, error) {
	if status != 0 && status != 1 && status != 2 && status != 3 {
		return MovieStatusValue(100), errors.New("Invalid Status.")
	}
	return MovieStatusValue(status), nil
}

func (m *MovieStatusModel)Complete()error{
	m.MovieStatus = 1
	return nil
}

func (m *MovieStatusModel)Error()error{
	m.MovieStatus = 2
	return nil
}