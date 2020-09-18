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
	err := validation.Validate(displayName,
		validation.Required,
	)
	if err != nil {
		return MovieDisplayName(""), err
	}
	return MovieDisplayName(displayName),nil
}

type Movie struct {
	ID          MovieID `json:"id" gorm:"primaryKey"`
	StoreName   MovieStoreName
	DisplayName MovieDisplayName
	UserID      UserID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewMovie(uploaderID UserID,storeName MovieStoreName,displayName MovieDisplayName) *Movie {
	return &Movie{
		StoreName:	storeName,
		DisplayName: displayName,
		UserID: uploaderID,
	}
}

func (m *Movie) GetURL() string {
	return "http://example.com/v1/movies/" + strconv.FormatUint(uint64(m.ID), 10) + "/" + strconv.FormatUint(uint64(m.ID), 10) + string(m.StoreName)
}
