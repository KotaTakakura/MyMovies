package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"fmt"
)

type IChangeOrderOfPlayListMovies interface {
	ChangeOrderOfPlayListMovies(changeOrderOfPlayListMoviesDTO *ChangeOrderOfPlayListMoviesDTO)error
}

type ChangeOrderOfPlayListMovies struct{
	PlayListMovieRepository	repository.PlayListMovieRepository
}

func NewChangeOrderOfPlayListMovies(p repository.PlayListMovieRepository)*ChangeOrderOfPlayListMovies{
	return &ChangeOrderOfPlayListMovies{
		PlayListMovieRepository: p,
	}
}

func (c ChangeOrderOfPlayListMovies)ChangeOrderOfPlayListMovies(changeOrderOfPlayListMoviesDTO *ChangeOrderOfPlayListMoviesDTO)error{
	playListMovies := c.PlayListMovieRepository.FindAll(changeOrderOfPlayListMoviesDTO.UserID,changeOrderOfPlayListMoviesDTO.PlayListID)
	movieIdAndOrderMap := make(map[model.MovieID]model.PlayListMovieOrder)

	for _,movieAndOrder := range changeOrderOfPlayListMoviesDTO.MovieIDAndOrder{
		movieIdAndOrderMap[movieAndOrder.MovieID] = movieAndOrder.Order
	}

	for i := 0; i < len(playListMovies); i++ {
		playListMovies[i].ChangeOrder(movieIdAndOrderMap[playListMovies[i].MovieID])
	}

	fmt.Println(playListMovies)

	saveErr := c.PlayListMovieRepository.SaveAll(playListMovies)
	if saveErr != nil {
		return saveErr
	}
	return nil
}

type ChangeOrderOfPlayListMoviesDTO struct{
	UserID model.UserID
	PlayListID model.PlayListID
	MovieIDAndOrder []MovieIdAndOrderForChangeOrderOfPlayListMoviesDTO
}

type MovieIdAndOrderForChangeOrderOfPlayListMoviesDTO struct{
	MovieID model.MovieID
	Order	model.PlayListMovieOrder
}