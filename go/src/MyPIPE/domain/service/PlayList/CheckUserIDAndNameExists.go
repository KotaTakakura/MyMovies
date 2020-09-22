package domain_service

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type CheckSameUserIDAndNameExists struct{
	PlayListRepository	repository.PlayListRepository
}

func NewCheckSameUserIDAndNameExists(u repository.PlayListRepository)*CheckSameUserIDAndNameExists{
	return &CheckSameUserIDAndNameExists{
		PlayListRepository: u,
	}
}

func(c CheckSameUserIDAndNameExists)CheckSameUserIDAndNameExists(userId model.UserID,playListName model.PlayListName)(bool,error){
	result,err := c.PlayListRepository.FindByUserIDAndName(userId,playListName)
	if err != nil{
		return false,err
	}
	if len(result) != 0{
		return true,nil
	}
	return false,nil
}