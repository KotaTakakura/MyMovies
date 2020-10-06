package queryService_infra

import (
	"MyPIPE/domain/model"
	queryService "MyPIPE/domain/queryService/UploadedMovies"
	"MyPIPE/infra"
)

type UploadedMovies struct {}

func NewUploadedMovies()*UploadedMovies{
	return &UploadedMovies{}
}

func (u UploadedMovies)Get(userId model.UserID)[]queryService.UploadedMoviesDTO{
	db := infra.ConnectGorm()
	defer db.Close()
	var result []queryService.UploadedMoviesDTO
	db.Table("movies").Where("user_id = ?",userId).Find(&result)
	return result
}