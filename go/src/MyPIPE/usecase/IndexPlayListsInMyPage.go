package usecase

import "MyPIPE/domain/queryService"

type IndexPlayListsInMyPage struct{
	IndexPlayListsInMyPageQueryService queryService.IndexPlayListsInMyPageQueryService
}

func NewIndexPlayListsInMyPage(ipq queryService.IndexPlayListsInMyPageQueryService)*IndexPlayListsInMyPage{
	return &IndexPlayListsInMyPage{
		IndexPlayListsInMyPageQueryService:	ipq,
	}
}

func (i IndexPlayListsInMyPage)All(userId uint64)*queryService.IndexPlayListsInMyPageDTO{
	return i.IndexPlayListsInMyPageQueryService.All(userId)
}