package usecase

import "MyPIPE/domain/queryService"

type IIndexPlaylistItemInMyPage interface {
	Find(indexPlayListMoviesInMyPageDTO IndexPlayListItemInMyPageDTO)*queryService.IndexPlayListMovieInMyPageDTO
}

type IndexPlayListItemInMyPage struct{
	IndexPlayListMovieInMyPage queryService.IndexPlayListMovieQueryService
}

func NewIndexPlayListItemInMyPage(i queryService.IndexPlayListMovieQueryService)*IndexPlayListItemInMyPage{
	return &IndexPlayListItemInMyPage{
		IndexPlayListMovieInMyPage: i,
	}
}

func (i IndexPlayListItemInMyPage)Find(indexPlayListMoviesInMyPageDTO IndexPlayListItemInMyPageDTO)*queryService.IndexPlayListMovieInMyPageDTO{
	return i.IndexPlayListMovieInMyPage.Find(indexPlayListMoviesInMyPageDTO.UserID,indexPlayListMoviesInMyPageDTO.PlayListID)
}

type IndexPlayListItemInMyPageDTO struct{
	UserID uint64
	PlayListID uint64
}

func NewIndexPlayListItemInMyPageDTO(userId uint64,playListId uint64)*IndexPlayListItemInMyPageDTO{
	return &IndexPlayListItemInMyPageDTO{
		UserID: userId,
		PlayListID: playListId,
	}
}