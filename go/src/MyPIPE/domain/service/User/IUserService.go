package domain_service

import "MyPIPE/domain/model"

type IUserService interface {
	CheckNameExists(userName model.UserName) bool
}
