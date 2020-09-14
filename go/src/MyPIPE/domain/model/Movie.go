package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"strconv"
	"time"
)

type MovieID uint64

func NewMovieID(movieId uint64) (MovieID,error) {
	err := validation.Validate(movieId,
		validation.Required,
	)
	if err != nil {
		return MovieID(0), err
	}
	return MovieID(movieId), nil
}

type MovieStoreName string

func NewMovieStoreName(storeName string) MovieStoreName {
	return MovieStoreName(storeName)
}

type MovieDisplayName string

func NewMovieDisplayName(displayName string) MovieDisplayName {
	return MovieDisplayName(displayName)
}

type Movie struct {
	ID          MovieID `json:"id" gorm:"primaryKey"`
	StoreName   MovieStoreName
	DisplayName MovieDisplayName
	UserID      UserID
	User        User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewMovie() *Movie {
	return &Movie{}
}

func (m *Movie) ReturnURL() string {
	return "http://example.com/v1/movies/" + strconv.FormatUint(uint64(m.ID), 10)
}
