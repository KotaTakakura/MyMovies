package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
)

type IIndexPlayListsInMyPage interface {
	All(indexPlayListsInMyPageDTO *IndexPlayListsInMyPageDTO) *queryService.IndexPlayListsInMyPageDTO
}

type IndexPlayListsInMyPage struct {
	IndexPlayListsInMyPageQueryService queryService.IndexPlayListsInMyPageQueryService
}

func NewIndexPlayListsInMyPage(ipq queryService.IndexPlayListsInMyPageQueryService) *IndexPlayListsInMyPage {
	return &IndexPlayListsInMyPage{
		IndexPlayListsInMyPageQueryService: ipq,
	}
}

func (i IndexPlayListsInMyPage) All(indexPlayListsInMyPageDTO *IndexPlayListsInMyPageDTO) *queryService.IndexPlayListsInMyPageDTO {
	return i.IndexPlayListsInMyPageQueryService.All(indexPlayListsInMyPageDTO.UserID)
}

type IndexPlayListsInMyPageDTO struct {
	UserID model.UserID
}

func NewIndexPlayListsInMyPageDTO(userId model.UserID) *IndexPlayListsInMyPageDTO {
	return &IndexPlayListsInMyPageDTO{
		UserID: userId,
	}
}
