package model

import "time"

type FollowUser struct {
	UserID   UserID	`gorm:"column:user_id"`
	FollowID UserID	`gorm:"column:follow_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewFollowUser(userId UserID,followId UserID) *FollowUser {
	return &FollowUser{
		UserID:	userId,
		FollowID:	followId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}