package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	domain_service "MyPIPE/domain/service/PlayList"
	"errors"
)

type AddPlayListItem struct{
	PlayListRepository	repository.PlayListRepository
	PlayListService domain_service.IPlayListService
}

func NewAddPlayListItem(p repository.PlayListRepository,ps domain_service.IPlayListService)*AddPlayListItem{
	return &AddPlayListItem{
		PlayListRepository: p,
		PlayListService: ps,
	}
}

func (a AddPlayListItem)AddPlayListItem(playListItemAddJson AddPlayListItemAddJson)error{

	playList,playListFindErr := a.PlayListRepository.FindByID(playListItemAddJson.PlayListID)
	if playList == nil || playList.UserID != playListItemAddJson.UserID{
		return errors.New("No Such PlayList.")
	}
	if playListFindErr != nil{
		return playListFindErr
	}

	var addPlayListItemErr error

	if a.PlayListService.CanAddItem(playListItemAddJson.MovieID) {
		addPlayListItemErr = playList.AddItem(playListItemAddJson.MovieID)
	}else{
		return errors.New("Can't Add This Movie.")
	}
	if addPlayListItemErr != nil{
		return addPlayListItemErr
	}

	savePlayListErr := a.PlayListRepository.Save(playList)
	if savePlayListErr != nil{
		return savePlayListErr
	}

	return nil
}

type AddPlayListItemAddJson struct{
	PlayListID model.PlayListID
	UserID model.UserID
	MovieID model.MovieID
}