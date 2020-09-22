package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
	"errors"
)

type PlayListID uint64

func NewPlayListID(playListID uint64)(PlayListID,error){
	err := validation.Validate(playListID,
		validation.Required,
	)
	if err != nil {
		return PlayListID(0), err
	}
	return PlayListID(playListID), nil
}

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

func (p *PlayList)AddItem(movieId MovieID)error{
	for _,playListItem := range p.PlayListItems{
		if playListItem == movieId{
			return errors.New("Duplicate PlayList Item.")
		}
	}
	p.PlayListItems = append(p.PlayListItems,movieId)
	return nil
}