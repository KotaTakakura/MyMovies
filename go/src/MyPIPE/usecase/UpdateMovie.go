package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type IUpdateMovie interface {
	Update(updateDTO UpdateDTO)(*model.Movie,error)
	UpdateStatus(updateStatusDTO UpdateStatusDTO)(*model.Movie,error)
}

type UpdateMovie struct{
	MovieRepository repository.MovieRepository
}

func NewUpdateMovie(m repository.MovieRepository)*UpdateMovie{
	return &UpdateMovie{
		MovieRepository: m,
	}
}

func (u UpdateMovie)Update(updateDTO UpdateDTO)(*model.Movie,error){
	movie,findMovieErr := u.MovieRepository.FindByUserIdAndMovieId(updateDTO.UserID,updateDTO.MovieID)
	if findMovieErr != nil {
		return nil,findMovieErr
	}
	changeDisplayNameErr := movie.ChangeDisplayName(updateDTO.DisplayName)
	if changeDisplayNameErr != nil{
		return nil, changeDisplayNameErr
	}

	changeDescriptionErr := movie.ChangeDescription(updateDTO.Description)
	if changeDescriptionErr != nil{
		return nil,changeDescriptionErr
	}

	changePublicErr := movie.ChangePublic(updateDTO.Public)
	if changePublicErr != nil{
		return nil,changePublicErr
	}

	changeStatusErr := movie.ChangeStatus(updateDTO.Status)
	if changeStatusErr != nil{
		return nil,changeStatusErr
	}

	updatedMovie,updateMovieErr := u.MovieRepository.Update(*movie)
	if updateMovieErr != nil{
		return nil, updateMovieErr
	}

	return updatedMovie,nil
}

type UpdateDTO struct{
	UserID model.UserID
	MovieID model.MovieID
	DisplayName model.MovieDisplayName
	Description model.MovieDescription
	Public	model.MoviePublic
	Status	model.MovieStatus
}

func (u UpdateMovie)UpdateStatus(updateStatusDTO UpdateStatusDTO)(*model.Movie,error){
	movie,findMovieErr := u.MovieRepository.FindByUserIdAndMovieId(updateStatusDTO.UserID,updateStatusDTO.MovieID)
	if findMovieErr != nil {
		return nil,findMovieErr
	}

	changeStatusErr := movie.ChangeStatus(updateStatusDTO.Status)
	if changeStatusErr != nil{
		return nil,changeStatusErr
	}

	updatedMovie,updateMovieErr := u.MovieRepository.Update(*movie)
	if updateMovieErr != nil{
		return nil, updateMovieErr
	}

	return updatedMovie,nil
}

type UpdateStatusDTO struct{
	UserID model.UserID
	MovieID model.MovieID
	Status	model.MovieStatus
}