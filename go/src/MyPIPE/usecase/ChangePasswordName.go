package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type ChangePassword struct{
	UserRepository repository.UserRepository
}

func NewChangePassword(u repository.UserRepository)*ChangePassword{
	return &ChangePassword{
		UserRepository: u,
	}
}

func (c ChangePassword)ChangePassword(changePasswordDTO *ChangePasswordDTO)error{
	user,findUserErr := c.UserRepository.FindById(changePasswordDTO.UserID)
	if findUserErr != nil{
		return findUserErr
	}

	changePasswordErr := user.ChangePassword(changePasswordDTO.Password)
	if changePasswordErr != nil{
		return changePasswordErr
	}

	updateUserErr := c.UserRepository.UpdateUser(user)
	if updateUserErr != nil{
		return updateUserErr
	}

	return nil
}

type ChangePasswordDTO struct{
	UserID model.UserID
	Password model.UserPassword
}

func NewChangePasswordDTO(userId model.UserID,password model.UserPassword)*ChangePasswordDTO{
	return &ChangePasswordDTO{
		UserID: userId,
		Password: password,
	}
}