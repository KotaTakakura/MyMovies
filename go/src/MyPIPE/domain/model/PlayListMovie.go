package model

import "time"

type PlayListMovieOrder int

func NewPlayListMovieOrder(playListOrder int)(PlayListMovieOrder,error){
	return PlayListMovieOrder(playListOrder),nil
}

type PlayListMovie struct {
	PlayListID PlayListID
	MovieID    MovieID
	Order	   PlayListMovieOrder
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewPlayListMovie(playListId PlayListID,movieId MovieID,order PlayListMovieOrder)*PlayListMovie{
	return &PlayListMovie{
		PlayListID: playListId,
		MovieID:    movieId,
		Order:      order,
	}
}
