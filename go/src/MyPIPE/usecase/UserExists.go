package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type UserExists struct{
	UserRepository	repository.UserRepository
}

func NewUserExists(u repository.UserRepository) *UserExists{
	return &UserExists{
		UserRepository: u,
	}
}

func (u UserExists)CheckUserExistsForAuth(email model.UserEmail, password string) (*model.User,error){
	user, err := u.UserRepository.FindByEmail(email)

	if user != nil && user.CheckPassword(password) && err == nil {
		return user, nil
	}

	return nil, err

}