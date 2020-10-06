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
	MovieDescription	model.MovieDescription	`gorm:"column:description" json:"movie_description"`
	MovieStatus	model.MovieStatus	`gorm:"column:status" json:"movie_status"`
	MoviePublic 	model.MoviePublic 	`gorm:"column:public" json:"movie_public"`
	MovieCreated	time.Time `gorm:"column:created_at" json:"movie_created_at"`
}