package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	domain_service "MyPIPE/domain/service/PlayList"
	"MyPIPE/infra"
	"errors"
)

type CreatePlayList struct{
	UserRepository repository.UserRepository
	PlayListRepository repository.PlayListRepository
}

func NewCreatePlayList(u repository.UserRepository,p repository.PlayListRepository)*CreatePlayList{
	return &CreatePlayList{
		UserRepository: u,
		PlayListRepository: p,
	}
}

func (c CreatePlayList)CreatePlayList(createPlayList CreatePlayListJson)error{
	playList := model.NewPlayList()
	var playListUserIDErr error
	playList.UserID,playListUserIDErr = model.NewUserID(createPlayList.UserID)
	if playListUserIDErr != nil{
		return playListUserIDErr
	}

	var playListNameErr error
	playList.Name,playListNameErr = model.NewPlayListName(createPlayList.PlayListName)
	if playListNameErr != nil{
		return playListNameErr
	}

	playListRepository := infra.NewPlayListPersistence()
	checkSameUserIDAndNameExistsService := domain_service.NewCheckSameUserIDAndNameExists(playListRepository)
	checkResult,checkResultErr := checkSameUserIDAndNameExistsService.CheckSameUserIDAndNameExists(playList.UserID,playList.Name)
	if checkResult {
		return errors.New("Duplicate PlayList Name.")
	}
	if checkResultErr != nil{
		return checkResultErr
	}

	saveErr := c.PlayListRepository.Save(playList)
	if saveErr != nil{
		return saveErr
	}
	return nil
}

type CreatePlayListJson struct{
	UserID uint64
	PlayListName string `json:"playlist_name"`
}