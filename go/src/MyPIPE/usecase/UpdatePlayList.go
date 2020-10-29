package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
)

type IUpdatePlayList interface {
	Update(updatePlayListDTO *UpdatePlayListDTO) error
}

type UpdatePlayList struct {
	PlayListRepository repository.PlayListRepository
}

func NewUpdatePlayList(playListRepository repository.PlayListRepository) *UpdatePlayList {
	return &UpdatePlayList{PlayListRepository: playListRepository}
}

func (u UpdatePlayList) Update(updatePlayListDTO *UpdatePlayListDTO) error {
	playList, playListFindErr := u.PlayListRepository.FindByIDAndUserID(updatePlayListDTO.PlayListID, updatePlayListDTO.UserID)
	if playListFindErr != nil {
		return playListFindErr
	}
	if playList == nil {
		return errors.New("No Such PlayList.")
	}

	changePlayListNameErr := playList.ChangeName(updatePlayListDTO.PlayListName)
	if changePlayListNameErr != nil {
		return changePlayListNameErr
	}

	changePlayListDescriptionErr := playList.ChangeDescription(updatePlayListDTO.PlayListDescription)
	if changePlayListDescriptionErr != nil {
		return changePlayListDescriptionErr
	}

	savePlayListErr := u.PlayListRepository.Save(playList)
	if savePlayListErr != nil {
		return savePlayListErr
	}

	return nil
}

type UpdatePlayListDTO struct {
	UserID              model.UserID
	PlayListID          model.PlayListID
	PlayListName        model.PlayListName
	PlayListDescription model.PlayListDescription
}

func NewUpdatePlayListDTO(userId model.UserID, playListId model.PlayListID, playListName model.PlayListName, playListDescription model.PlayListDescription) *UpdatePlayListDTO {
	return &UpdatePlayListDTO{
		UserID:              userId,
		PlayListID:          playListId,
		PlayListName:        playListName,
		PlayListDescription: playListDescription,
	}
}
