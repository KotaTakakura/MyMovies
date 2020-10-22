package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	domain_service "MyPIPE/domain/service/PlayList"
	"MyPIPE/infra"
	"errors"
)

type ICreatePlayList interface {
	CreatePlayList(createPlayList CreatePlayListDTO) error
}

type CreatePlayList struct {
	UserRepository     repository.UserRepository
	PlayListRepository repository.PlayListRepository
}

func NewCreatePlayList(u repository.UserRepository, p repository.PlayListRepository) *CreatePlayList {
	return &CreatePlayList{
		UserRepository:     u,
		PlayListRepository: p,
	}
}

func (c CreatePlayList) CreatePlayList(createPlayList CreatePlayListDTO) error {
	playList := model.NewPlayList(createPlayList.UserID, createPlayList.PlayListName, createPlayList.PlayListDescription)

	playListRepository := infra.NewPlayListPersistence()
	checkSameUserIDAndNameExistsService := domain_service.NewCheckSameUserIDAndNameExists(playListRepository)
	checkResult, checkResultErr := checkSameUserIDAndNameExistsService.CheckSameUserIDAndNameExists(playList.UserID, playList.Name)
	if checkResult {
		return errors.New("Duplicate PlayList Name.")
	}
	if checkResultErr != nil {
		return checkResultErr
	}

	saveErr := c.PlayListRepository.Save(playList)
	if saveErr != nil {
		return saveErr
	}
	return nil
}

type CreatePlayListDTO struct {
	UserID              model.UserID
	PlayListName        model.PlayListName
	PlayListDescription model.PlayListDescription
}
