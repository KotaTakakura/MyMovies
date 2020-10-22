package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type IChangeUserName interface {
	ChangeUserName(changeUserNameDTO *ChangeUserNameDTO) error
}

type ChangeUserName struct {
	UserRepository repository.UserRepository
}

func NewChangeUserName(u repository.UserRepository) *ChangeUserName {
	return &ChangeUserName{
		UserRepository: u,
	}
}

func (c ChangeUserName) ChangeUserName(changeUserNameDTO *ChangeUserNameDTO) error {
	user, findUserErr := c.UserRepository.FindById(changeUserNameDTO.UserID)
	if findUserErr != nil {
		return findUserErr
	}

	changeUserNameErr := user.ChangeName(changeUserNameDTO.UserName)
	if changeUserNameErr != nil {
		return changeUserNameErr
	}

	updateUserErr := c.UserRepository.UpdateUser(user)
	if updateUserErr != nil {
		return updateUserErr
	}

	return nil
}

type ChangeUserNameDTO struct {
	UserID   model.UserID
	UserName model.UserName
}

func NewChangeUserNameDTO(userId model.UserID, userName model.UserName) *ChangeUserNameDTO {
	return &ChangeUserNameDTO{
		UserID:   userId,
		UserName: userName,
	}
}
