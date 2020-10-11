package usecase

import "MyPIPE/domain/queryService"

type IndexMovie struct{
	IndexMovieQueryService queryService.IndexMovieQueryService
}

func NewIndexMovie(imq queryService.IndexMovieQueryService)*IndexMovie{
	return &IndexMovie{
		IndexMovieQueryService: imq,
	}
}

func (i IndexMovie)Search(page int,keyWord string)queryService.IndexMovieDTO{
	if keyWord == ""{
		return i.IndexMovieQueryService.All(page)
	}
	return i.IndexMovieQueryService.Search(page,keyWord)
}