package model

import "time"

type PlayListID uint64

type PlayListName string

type PlayList struct {
	ID            PlayListID `json:"id" gorm:"primaryKey"`
	UserID        UserID
	Name          PlayListName
	PlayListItems []MovieID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
