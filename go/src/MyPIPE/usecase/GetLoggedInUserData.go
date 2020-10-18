package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
)

type GetLoggedInUserData struct{
	GetLoggedInUserDataQueryService queryService.GetLoggedInUserDataQueryService
}

func NewGetLoggedInUserData(g queryService.GetLoggedInUserDataQueryService)*GetLoggedInUserData{
	return &GetLoggedInUserData{
		GetLoggedInUserDataQueryService: g,
	}
}

func (g GetLoggedInUserData)Find(getLoggedInUserDataDTO *GetLoggedInUserDataDTO)*queryService.GetLoggedInUserDataDTO{
	return g.GetLoggedInUserDataQueryService.FindByUserId(getLoggedInUserDataDTO.UserID)
}

type GetLoggedInUserDataDTO struct{
	UserID model.UserID
}

func NewGetLoggedInUserDataDTO(userId model.UserID)*GetLoggedInUserDataDTO{
	return &GetLoggedInUserDataDTO{
		UserID: userId,
	}
}