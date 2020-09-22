package model

import "time"

type PlayListID uint64

type PlayListName string

func NewPlayListName(playListName string)(PlayListName,error){
	return PlayListName(playListName),nil
}

type PlayList struct {
	ID            PlayListID `json:"id" gorm:"primaryKey"`
	UserID        UserID `gorm:"column:user_id"`
	Name          PlayListName `gorm:"column:name"`
	PlayListItems []MovieID `gorm:"-"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewPlayList()*PlayList{
	return &PlayList{}
}

func (PlayList) TableName() string {
	return "play_lists"
}