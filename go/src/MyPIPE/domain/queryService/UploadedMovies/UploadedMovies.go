package queryService

import (
	"MyPIPE/domain/model"
	"time"
)

type UploadedMovies interface {
	Get(userId model.UserID)[]UploadedMoviesDTO
}

type UploadedMoviesDTO struct{
	MovieID	model.MovieID	`gorm:"column:id" json:"movie_id"`
	MovieName	model.MovieDisplayName	`gorm:"column:display_name" json:"movie_name"`
	MovieDescription	string	`gorm:"column:description" json:"movie_description"`
	MovieProgress	int	`gorm:"column:progress" json:"movie_progress"`
	MovieCreated	time.Time `gorm:"column:created_at" json:"movie_created_at"`
}