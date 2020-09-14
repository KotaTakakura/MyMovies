package model

import "time"

type PlayListItem struct {
	PlayListID PlayListID
	MovieID    MovieID
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
