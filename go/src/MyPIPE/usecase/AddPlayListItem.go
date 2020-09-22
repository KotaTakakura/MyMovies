package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
	"fmt"
)

type AddPlayListItem struct{
	PlayListRepository	repository.PlayListRepository
	MovieRepository repository.MovieRepository
}

func NewAddPlayListItem(p repository.PlayListRepository,m repository.MovieRepository)*AddPlayListItem{
	return &AddPlayListItem{
		PlayListRepository: p,
		MovieRepository: m,
	}
}

func (a AddPlayListItem)AddPlayListItem(playListItemAddJson PlayListItemAddJson)error{
	playListID,playListIDErr := model.NewPlayListID(playListItemAddJson.PlayListID)
	if playListIDErr != nil{
		fmt.Println("111")
		return playListIDErr
	}

	movieId,movieIdErr := model.NewMovieID(playListItemAddJson.MovieID)
	if movieIdErr != nil{
		fmt.Println("222")
		return movieIdErr
	}

	playList,playListFindErr := a.PlayListRepository.FindByID(playListID)
	if playList == nil || uint64(playList.UserID) != playListItemAddJson.UserID{
		return errors.New("No Such PlayList.")
	}
	if playListFindErr != nil{
		return playListFindErr
	}

	movie,movieFindErr := a.MovieRepository.FindById(movieId)
	if movie == nil{
		return errors.New("No Such Movie.")
	}
	if movieFindErr != nil{
		return movieFindErr
	}

	addPlayListItemErr := playList.AddItem(movieId)
	if addPlayListItemErr != nil{
		return addPlayListItemErr
	}

	savePlayListErr := a.PlayListRepository.Save(playList)
	if savePlayListErr != nil{
		return savePlayListErr
	}

	return nil
}

type PlayListItemAddJson struct{
	PlayListID uint64 `json:"play_list_id"`
	UserID uint64 `json:"user_id"`
	MovieID uint64 `json:"movie_id"`
}