package queryService

import (
	"MyPIPE/domain/model"
	"time"
)

type GetLoggedInUserDataDTO struct{
	Name	string	`json:"user_name"`
	Email string	`json:"user_email"`
	Birthday	time.Time	`json:"user_birthday"`
	CreatedAt	time.Time	`json:"user_created_datetime"`
	AvatarName	string	`json:"user_avatar_name"`
}

type GetLoggedInUserDataQueryService interface {
	FindByUserId(userId model.UserID)*GetLoggedInUserDataDTO
}