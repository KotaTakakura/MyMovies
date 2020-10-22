package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type IDeletePlayList interface {
	Delete(deletePlayListDTO *DeletePlayListDTO) error
}

type DeletePlayList struct {
	PlayListRepository repository.PlayListRepository
}

func NewDeletePlayList(p repository.PlayListRepository) *DeletePlayList {
	return &DeletePlayList{
		PlayListRepository: p,
	}
}

func (d DeletePlayList) Delete(deletePlayListDTO *DeletePlayListDTO) error {
	result := d.PlayListRepository.Remove(deletePlayListDTO.UserID, deletePlayListDTO.PlaylistID)
	if result != nil {
		return result
	}
	return nil
}

type DeletePlayListDTO struct {
	UserID     model.UserID
	PlaylistID model.PlayListID
}

func NewDeletePlayListDTO(userId model.UserID, playListId model.PlayListID) *DeletePlayListDTO {
	return &DeletePlayListDTO{
		UserID:     userId,
		PlaylistID: playListId,
	}
}
