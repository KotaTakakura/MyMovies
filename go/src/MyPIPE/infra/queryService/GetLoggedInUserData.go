package queryService_infra

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
	"MyPIPE/infra"
)

type GetLoggedInUserData struct{}

func NewGetLoggedInUserData() *GetLoggedInUserData {
	return &GetLoggedInUserData{}
}

func (g GetLoggedInUserData) FindByUserId(userId model.UserID) *queryService.GetLoggedInUserDataDTO {
	db := infra.ConnectGorm()
	defer db.Close()
	var getLoggedInUserData queryService.GetLoggedInUserDataDTO
	var count int
	db.Table("users").Where("id = ?", userId).First(&getLoggedInUserData).Count(&count)
	if count == 0 {
		return nil
	}
	return &getLoggedInUserData
}
