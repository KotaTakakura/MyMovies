package model

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"mime/multipart"
	"path/filepath"
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

func NewMovieStoreName(storeName string) (MovieStoreName, error) {
	return MovieStoreName(storeName), nil
}

type MovieDisplayName string

func NewMovieDisplayName(displayName string) (MovieDisplayName, error) {
	return MovieDisplayName(displayName), nil
}

type MovieThumbnailName string

func NewMovieThumbnailName(thumbnailHeader multipart.FileHeader) (MovieThumbnailName, error) {
	extension := filepath.Ext(thumbnailHeader.Filename)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	return MovieThumbnailName("_" + timestamp + extension), nil
}

type MoviePublic uint

func NewMoviePublic(public uint) (MoviePublic, error) {
	return MoviePublic(public), nil
}

type MovieStatus uint

func NewMovieStatus(status uint) (MovieStatus, error) {
	return MovieStatus(status), nil
}

type MovieDescription string

func NewMovieDescription(description string) (MovieDescription, error) {
	return MovieDescription(description), nil
}

type Movie struct {
	ID            MovieID            `json:"id" gorm:"primaryKey"`
	StoreName     MovieStoreName     `gorm:"column:store_name"`
	DisplayName   MovieDisplayName   `gorm:"column:display_name"`
	Description   MovieDescription   `gorm:"column:description"`
	ThumbnailName MovieThumbnailName `gorm:"column:thumbnail_name"`
	UserID        UserID             `gorm:"column:user_id"`
	Public        MoviePublic        `gorm:"column:public"`
	Status        MovieStatus        `gorm:"column:status"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewMovie(uploaderID UserID, storeName MovieStoreName, displayName MovieDisplayName, thumbnailName MovieThumbnailName) *Movie {
	return &Movie{
		StoreName:     storeName,
		DisplayName:   displayName,
		UserID:        uploaderID,
		Description:   MovieDescription(""),
		ThumbnailName: thumbnailName,
		Public:        MoviePublic(0),
		Status:        MovieStatus(0),
	}
}

func (m *Movie) GetURL() string {
	return "http://example.com/v1/movies/" + strconv.FormatUint(uint64(m.ID), 10) + "/" + strconv.FormatUint(uint64(m.ID), 10) + string(m.StoreName)
}

func (m *Movie) ChangeDisplayName(displayName MovieDisplayName) error {
	m.DisplayName = displayName
	return nil
}

func (m *Movie) ChangeDescription(description MovieDescription) error {
	m.Description = description
	return nil
}

func (m *Movie) ChangePublic(publicStatus MoviePublic) error {
	if publicStatus == 1 && m.DisplayName == "" {
		return errors.New("Title Not Set.")
	}

	if publicStatus == 1 && m.Status == 0 {
		return errors.New("Status Not Complete.")
	}
	m.Public = publicStatus
	return nil
}

func (m *Movie) ChangeStatus(status MovieStatus) error {
	m.Status = status
	return nil
}

func (m *Movie) ChangeThumbnailName(thumbnailName MovieThumbnailName) error {
	m.ThumbnailName = thumbnailName
	return nil
}
