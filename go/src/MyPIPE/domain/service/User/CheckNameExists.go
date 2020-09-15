package domain_service

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type CheckNameExists struct{
	UserRepository	repository.UserRepository
}

func NewCheckNameExists(u repository.UserRepository)*CheckNameExists{
	return &CheckNameExists{
		UserRepository: u,
	}
}

func (c CheckNameExists)CheckNameExists(userName model.UserName)bool{
	user,err := c.UserRepository.FindByName(userName)
	if user == nil && err != nil{
		return false
	}
	return true
}