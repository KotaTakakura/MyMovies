package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type ICheckUserAlreadyLikedMovie interface {
	Find(checkUserAlreadyLikedMovieFindDTO *CheckUserAlreadyLikedMovieFindDTO) bool
}

type CheckUserAlreadyLikedMovie struct {
	MovieEvaluationRepository repository.MovieEvaluationRepository
}

func NewCheckUserAlreadyLikedMovie(mer repository.MovieEvaluationRepository) *CheckUserAlreadyLikedMovie {
	return &CheckUserAlreadyLikedMovie{
		MovieEvaluationRepository: mer,
	}
}

func (c CheckUserAlreadyLikedMovie) Find(checkUserAlreadyLikedMovieFindDTO *CheckUserAlreadyLikedMovieFindDTO) bool {
	evaluation, _ := model.NewEvaluation("good")
	movieEvaluation := c.MovieEvaluationRepository.FindByUserIdAndMovieIdAndEvaluation(checkUserAlreadyLikedMovieFindDTO.UserID, checkUserAlreadyLikedMovieFindDTO.MovieID, evaluation)
	if movieEvaluation == nil {
		return false
	}
	return true
}

type CheckUserAlreadyLikedMovieFindDTO struct {
	UserID  model.UserID
	MovieID model.MovieID
}

func NewCheckUserAlreadyLikedMovieFindDTO(userId model.UserID, movieId model.MovieID) *CheckUserAlreadyLikedMovieFindDTO {
	return &CheckUserAlreadyLikedMovieFindDTO{
		UserID:  userId,
		MovieID: movieId,
	}
}
