package queryService_infra

import (
	"MyPIPE/domain/model"
	queryService "MyPIPE/domain/queryService/UploadedMovies"
	"MyPIPE/infra"
	"github.com/jinzhu/gorm"
)

type UploadedMovies struct {
	DatabaseAccessor *gorm.DB
}

func NewUploadedMovies()*UploadedMovies{
	return &UploadedMovies{
		DatabaseAccessor: infra.ConnectGorm(),
	}
}

func (u UploadedMovies)Get(userId model.UserID)[]queryService.UploadedMoviesDTO{
	var result []queryService.UploadedMoviesDTO
	u.DatabaseAccessor.Table("movies").Where("user_id = ?",userId).Find(&result)
	return result
}