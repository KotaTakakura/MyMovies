package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
)

type IIndexPlaylistItemInMyPage interface {
	Find(indexPlayListMoviesInMyPageDTO *IndexPlayListItemInMyPageDTO) *queryService.IndexPlayListMovieInMyPageDTO
}

type IndexPlayListItemInMyPage struct {
	IndexPlayListMovieInMyPage queryService.IndexPlayListMovieQueryService
}

func NewIndexPlayListItemInMyPage(i queryService.IndexPlayListMovieQueryService) *IndexPlayListItemInMyPage {
	return &IndexPlayListItemInMyPage{
		IndexPlayListMovieInMyPage: i,
	}
}

func (i IndexPlayListItemInMyPage) Find(indexPlayListMoviesInMyPageDTO *IndexPlayListItemInMyPageDTO) *queryService.IndexPlayListMovieInMyPageDTO {
	return i.IndexPlayListMovieInMyPage.Find(indexPlayListMoviesInMyPageDTO.UserID, indexPlayListMoviesInMyPageDTO.PlayListID)
}

type IndexPlayListItemInMyPageDTO struct {
	UserID     model.UserID
	PlayListID model.PlayListID
}

func NewIndexPlayListItemInMyPageDTO(userId model.UserID, playListId model.PlayListID) *IndexPlayListItemInMyPageDTO {
	return &IndexPlayListItemInMyPageDTO{
		UserID:     userId,
		PlayListID: playListId,
	}
}
