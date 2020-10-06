package usecase

import (
	"MyPIPE/domain/model"
	queryService "MyPIPE/domain/queryService/UploadedMovies"
)

type UploadedMovies struct{
	UploadedMoviesQueryService	queryService.UploadedMovies
}

func NewUploadedMovies(umq queryService.UploadedMovies)*UploadedMovies{
	return &UploadedMovies{
		UploadedMoviesQueryService: umq,
	}
}

func (u UploadedMovies)Get(userId model.UserID)[]queryService.UploadedMoviesDTO{
	result := u.UploadedMoviesQueryService.Get(userId)
	return result
}