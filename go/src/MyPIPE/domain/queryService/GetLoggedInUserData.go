package queryService

import (
	"MyPIPE/domain/model"
	"time"
)

type GetLoggedInUserDataDTO struct{
	ID	uint64	`json:"user_id"`
	Name	string	`json:"user_name"`
	Email string	`json:"user_email"`
	Birthday	time.Time	`json:"user_birthday"`
	ProfileImageName string	`json:"user_profile_image_name"`
	CreatedAt	time.Time	`json:"user_created_datetime"`
	AvatarName	string	`json:"user_avatar_name"`
}

type GetLoggedInUserDataQueryService interface {
	FindByUserId(userId model.UserID)*GetLoggedInUserDataDTO
}