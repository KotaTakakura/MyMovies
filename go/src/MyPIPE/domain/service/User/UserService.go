package domain_service

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type UserService struct{
	UserRepository	repository.UserRepository
}

func NewUserService(u repository.UserRepository)*UserService{
	return &UserService{
		UserRepository: u,
	}
}

func (c UserService)CheckNameExists(userName model.UserName)bool{
	user,err := c.UserRepository.FindByName(userName)
	if user == nil && err != nil{
		return false
	}
	return true
}