package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"strconv"
	"time"
)

type MovieID uint64

func NewMovieID(movieId uint64) (MovieID, error) {
	err := validation.Validate(movieId,
		validation.Required,
	)
	if err != nil {
		return MovieID(0), err
	}
	return MovieID(movieId), nil
}

type MovieStoreName string

func NewMovieStoreName(storeName string) (MovieStoreName,error) {
	return MovieStoreName(storeName),nil
}

type MovieDisplayName string

func NewMovieDisplayName(displayName string) (MovieDisplayName,error) {
	return MovieDisplayName(displayName),nil
}

type MoviePublic uint

func NewMoviePublic(public uint) (MoviePublic,error){
	err := validation.Validate(public,
		validation.Required,
		validation.In(0,1),
	)
	if err != nil{
		return	MoviePublic(0),err
	}
	return MoviePublic(public),nil
}

type MovieStatus uint

func NewMovieStatus(status uint) (MovieStatus,error){
	err := validation.Validate(status,
		validation.Required,
		validation.In(0,1,2,3),
	)
	if err != nil{
		return MovieStatus(0),err
	}
	return MovieStatus(status),nil
}

type Movie struct {
	ID          MovieID `json:"id" gorm:"primaryKey"`
	StoreName   MovieStoreName
	DisplayName MovieDisplayName
	UserID      UserID
	Public  	MoviePublic
	Status		MovieStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewMovie(uploaderID UserID,storeName MovieStoreName,displayName MovieDisplayName) *Movie {
	return &Movie{
		StoreName:	storeName,
		DisplayName: displayName,
		UserID: uploaderID,
		Public:	MoviePublic(0),
		Status: MovieStatus(0),
	}
}

func (m *Movie) GetURL() string {
	return "http://example.com/v1/movies/" + strconv.FormatUint(uint64(m.ID), 10) + "/" + strconv.FormatUint(uint64(m.ID), 10) + string(m.StoreName)
}
