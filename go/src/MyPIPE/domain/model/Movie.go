package model

import (
	"strconv"
	"time"
)

type Movie struct {
	ID        uint64 `json:"id" gorm:"primaryKey"`
	Name      string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Movie) ReturnURL() string {
	return "http://example.com/v1/movies/" + strconv.FormatUint(m.ID, 10)
}
