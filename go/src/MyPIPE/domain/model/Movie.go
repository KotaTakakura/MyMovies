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
	err := validation.Validate(displayName,
		validation.RuneLength(0, 200),
	)
	if err != nil {
		return MovieDisplayName(""), err
	}
	return MovieDisplayName(displayName), nil
}

type MovieThumbnailName string

func NewMovieThumbnailName(thumbnailHeader multipart.FileHeader) (MovieThumbnailName, error) {
	extension := filepath.Ext(thumbnailHeader.Filename)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	return MovieThumbnailName("_" + timestamp + extension), nil
}

type MovieThumbnail struct {
	Name       string
	File       multipart.File
	FileHeader multipart.FileHeader
}

func NewMovieThumbnail(file multipart.File, fileHeader multipart.FileHeader) (*MovieThumbnail, error) {
	extension := filepath.Ext(fileHeader.Filename)
	if !(extension == ".jpg" || extension == ".JPG" || extension == ".jpeg" || extension == ".JPEG" || extension == ".png" || extension == ".PNG" || extension == ".bmp" || extension == ".BMP" || extension == ".gif" || extension == ".GIF") {
		return nil, errors.New("Image File Only.")
	}
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	if fileHeader.Size > 2000000 {
		return nil, errors.New("Too Large File.")
	}

	return &MovieThumbnail{
		Name:       "_" + timestamp + extension,
		File:       file,
		FileHeader: fileHeader,
	}, nil
}

type MoviePublic uint

func NewMoviePublic(public uint) (MoviePublic, error) {
	if public != 0 && public != 1 && public != 2 && public != 3 {
		return MoviePublic(100), errors.New("Invalid Public Status.")
	}
	return MoviePublic(public), nil
}

type MovieStatus uint

func NewMovieStatus(status uint) (MovieStatus, error) {
	if status != 0 && status != 1 && status != 2 && status != 3 {
		return MovieStatus(100), errors.New("Invalid Status.")
	}
	return MovieStatus(status), nil
}

type MovieDescription string

func NewMovieDescription(description string) (MovieDescription, error) {
	return MovieDescription(description), nil
}

type MovieFile struct {
	StoreName  string
	File       multipart.File
	FileHeader multipart.FileHeader
}

func NewMovieFile(file multipart.File, fileHeader multipart.FileHeader) (*MovieFile, error) {
	extension := filepath.Ext(fileHeader.Filename)
	if !(extension == ".mp4" || extension == ".mov") {
		return nil, errors.New("Movie File Only.")
	}

	if fileHeader.Size > 1000000000 {
		return nil, errors.New("Too Large File.")
	}

	return &MovieFile{
		StoreName:  extension,
		File:       file,
		FileHeader: fileHeader,
	}, nil
}

type MovieThumbnailStatus uint

type Movie struct {
	ID              MovieID          `json:"id" gorm:"primaryKey"`
	StoreName       string           `gorm:"column:store_name"`
	DisplayName     MovieDisplayName `gorm:"column:display_name"`
	Description     MovieDescription `gorm:"column:description"`
	ThumbnailName   string           `gorm:"column:thumbnail_name"`
	UserID          UserID           `gorm:"column:user_id"`
	Public          MoviePublic      `gorm:"column:public"`
	Status          MovieStatus      `gorm:"column:status"`
	ThumbnailStatus MovieThumbnailStatus
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewMovie(uploaderID UserID, movieFile *MovieFile, thumbnail *MovieThumbnail) *Movie {
	return &Movie{
		StoreName:       movieFile.StoreName,
		DisplayName:     MovieDisplayName(""),
		UserID:          uploaderID,
		Description:     MovieDescription(""),
		ThumbnailName:   thumbnail.Name,
		Public:          MoviePublic(0),
		Status:          MovieStatus(0),
		ThumbnailStatus: MovieThumbnailStatus(0),
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

func (m *Movie) Complete() error {
	m.Status = MovieStatus(1)
	return nil
}

func (m *Movie) ChangeThumbnailStatusComplete() error {
	m.ThumbnailStatus = MovieThumbnailStatus(1)
	return nil
}

func (m *Movie) ChangeThumbnailStatusInComplete() error {
	m.ThumbnailStatus = MovieThumbnailStatus(0)
	return nil
}

func (m *Movie) ChangeThumbnailName(thumbnail *MovieThumbnail) error {
	m.ThumbnailName = thumbnail.Name
	m.ThumbnailStatus = MovieThumbnailStatus(0)
	return nil
}
