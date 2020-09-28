package domain_service

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type PlayListService struct{
	MovieRepository repository.MovieRepository
}

func NewPlayListService(m repository.MovieRepository)*PlayListService{
	return &PlayListService{
		MovieRepository: m,
	}
}

func (p PlayListService)CanAddItem(movieId model.MovieID)bool{
	result,_ := p.MovieRepository.FindById(movieId)
	if result == nil{
		return false
	}
	return true
}