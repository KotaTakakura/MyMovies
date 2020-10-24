package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type ICreatePlayList interface {
	CreatePlayList(createPlayList *CreatePlayListDTO) error
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

func (c CreatePlayList) CreatePlayList(createPlayList *CreatePlayListDTO) error {
	playList := model.NewPlayList(createPlayList.UserID, createPlayList.PlayListName, createPlayList.PlayListDescription)

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

func NewCreatePlayListDTO(userId model.UserID, playListName model.PlayListName, playListDescription model.PlayListDescription) *CreatePlayListDTO {
	return &CreatePlayListDTO{
		UserID:              userId,
		PlayListName:        playListName,
		PlayListDescription: playListDescription,
	}
}
