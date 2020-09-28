package domain_service

import "MyPIPE/domain/model"

type IPlayListService interface {
	CanAddItem(movieId model.MovieID)bool
}
