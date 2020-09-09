package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type User struct {
	UserRepository repository.UserRepository
}

func NewUser(userRepository repository.UserRepository) *User {
	return &User{
		UserRepository: userRepository,
	}
}

func (u *User) RegisterUser(newUser *model.User) {
	u.UserRepository.SetUser(newUser)
}
