package queryService

type IndexMovieDTO struct {
	Movies []MoviesForIndexMovieDTO `json:"movies"`
	Count  uint64                   `json:"movie_count"`
}

type MoviesForIndexMovieDTO struct {
	MovieID          uint64 `json:"movie_id"`
	MovieDisplayName string `json:"movie_title"`
	UserID           uint64 `json:"user_id"`
	UserName         string `json:"user_name"`
	ThumbnailName    string `json:"movie_thumbnail_name"`
}

type IndexMovieQueryService interface {
	Search(page IndexMovieQueryServicePage, keyWord string, order IndexMovieQueryServiceOrder) IndexMovieDTO
	All(page IndexMovieQueryServicePage, order IndexMovieQueryServiceOrder) IndexMovieDTO
}

type IndexMovieQueryServiceOrder string

func NewIndexMovieQueryServiceOrder(order string) (IndexMovieQueryServiceOrder, error) {
	if order == "desc" {
		return IndexMovieQueryServiceOrder(order), nil
	}
	return IndexMovieQueryServiceOrder("asc"), nil
}

type IndexMovieQueryServicePage uint

func NewIndexMovieQueryServicePage(page uint)(IndexMovieQueryServicePage,error){
	return IndexMovieQueryServicePage(page),nil
}
