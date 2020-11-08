package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type IDeleteUser interface {
	DeleteUser(deleteUserDTO *DeleteUserDTO) error
}

type DeleteUser struct {
	UserRepository repository.UserRepository
}

func NewDeleteUser(userRepository repository.UserRepository) *DeleteUser {
	return &DeleteUser{
		UserRepository: userRepository,
	}
}

func (d DeleteUser) DeleteUser(deleteUserDTO *DeleteUserDTO) error {
	deleteUserErr := d.UserRepository.Remove(deleteUserDTO.UserID)
	if deleteUserErr != nil {
		return deleteUserErr
	}
	return nil
}

type DeleteUserDTO struct {
	UserID model.UserID
}

func NewDeleteUserDTO(userId model.UserID) *DeleteUserDTO {
	return &DeleteUserDTO{
		UserID: userId,
	}
}
