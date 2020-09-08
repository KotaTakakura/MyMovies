package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"github.com/google/uuid"
	"time"
)

type UserTemporaryRegistration struct {
	UserRepository	repository.UserRepository
}

func NewUserTemporaryRegistration(userRepository repository.UserRepository) *UserTemporaryRegistration{
	return &UserTemporaryRegistration{
		UserRepository: userRepository,
	}
}

func (u *UserTemporaryRegistration)TemporaryRegister(user *model.User){
	registeredUser := u.UserRepository.FindByEmail(user.Email)
	if registeredUser != nil{
		return
	}
	user.Token = uuid.New().String()
	user.Birthday = time.Date(1000, 1, 1, 0, 0, 0, 0, time.Local)
	u.UserRepository.SetUser(user)
}