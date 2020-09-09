package model

import (
	"strconv"
	"time"
)

type Movie struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	StoreName   string
	DisplayName string
	UserID      uint64
	User        User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewMovie() *Movie {
	return &Movie{}
}

func (m *Movie) ReturnURL() string {
	return "http://example.com/v1/movies/" + strconv.FormatUint(m.ID, 10)
}
