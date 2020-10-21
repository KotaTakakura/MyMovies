package usecase

import "MyPIPE/domain/queryService"

type IIndexMovie interface {
	Search(indexMovieSearchDTO *IndexMovieSearchDTO)queryService.IndexMovieDTO
}

type IndexMovie struct{
	IndexMovieQueryService queryService.IndexMovieQueryService
}

func NewIndexMovie(imq queryService.IndexMovieQueryService)*IndexMovie{
	return &IndexMovie{
		IndexMovieQueryService: imq,
	}
}

func (i IndexMovie)Search(indexMovieSearchDTO *IndexMovieSearchDTO)queryService.IndexMovieDTO{
	if indexMovieSearchDTO.KeyWord == ""{
		return i.IndexMovieQueryService.All(indexMovieSearchDTO.Page,indexMovieSearchDTO.Order)
	}
	return i.IndexMovieQueryService.Search(indexMovieSearchDTO.Page,indexMovieSearchDTO.KeyWord,indexMovieSearchDTO.Order)
}

type IndexMovieSearchDTO struct{
	Page int
	KeyWord string
	Order queryService.IndexMovieQueryServiceOrder
}

func NewIndexMovieSearchDTO(page int,keyWord string,order queryService.IndexMovieQueryServiceOrder)*IndexMovieSearchDTO{
	return &IndexMovieSearchDTO{
		Page:    page,
		KeyWord: keyWord,
		Order:   order,
	}
}