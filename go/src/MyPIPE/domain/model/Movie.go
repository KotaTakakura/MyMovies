package model

import "time"

type Movie struct{
	Name	string
	URL		string
	UserID	uint
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
