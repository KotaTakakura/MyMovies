package model

type FollowUser struct {
	UserID   UserID
	FollowID UserID
}

func NewFollowUser() *FollowUser {
	return &FollowUser{}
}
