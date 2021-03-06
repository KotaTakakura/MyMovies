package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type PlayListMovieOrder int

func NewPlayListMovieOrder(playListOrder int) (PlayListMovieOrder, error) {
	err := validation.Validate(playListOrder,
		validation.Required,
	)
	if err != nil {
		return PlayListMovieOrder(0), err
	}
	return PlayListMovieOrder(playListOrder), nil
}

type PlayListMovie struct {
	PlayListID PlayListID
	MovieID    MovieID
	Order      PlayListMovieOrder
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewPlayListMovie(playListId PlayListID, movieId MovieID, order PlayListMovieOrder) *PlayListMovie {
	return &PlayListMovie{
		PlayListID: playListId,
		MovieID:    movieId,
		Order:      order,
	}
}

func (p *PlayListMovie) ChangeOrder(order PlayListMovieOrder) {
	p.Order = order
}
