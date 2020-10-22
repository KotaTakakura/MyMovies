package handler

import (
	mock_repository "MyPIPE/Test/mock/repository"
	mock_usecase "MyPIPE/Test/mock/usecase"
	"MyPIPE/domain/model"
	"MyPIPE/handler"
	"MyPIPE/usecase"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"strings"
	"testing"
	"fmt"
	"reflect"
)

func TestCheckUserAlreadyLikedMovie(t *testing.T){
	trueCases := []struct {
		userId uint64
		movieId uint64
	}{
		{userId: 10, movieId: 100},
		{userId: 99999, movieId: 999},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieEvaluationRepository := mock_repository.NewMockMovieEvaluationRepository(ctrl)
	checkAlreadyLikeDMovieUsecase := mock_usecase.NewMockICheckUserAlreadyLikedMovie(ctrl)
	checkUserAlreadyLikedMovieHandler := handler.NewCheckUserAlreadyLikedMovie(movieEvaluationRepository,checkAlreadyLikeDMovieUsecase)

	for _,trueCase := range trueCases{
		// ポストデータ
		bodyReader := strings.NewReader("")

		// リクエスト生成
		req := httptest.NewRequest("GET", "/?user_id=" + fmt.Sprint(trueCase.userId) + "&movie_id=" + fmt.Sprint(trueCase.movieId), bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")


		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req

		checkAlreadyLikeDMovieUsecase.EXPECT().Find(gomock.Any()).DoAndReturn(func(data interface{})error{
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.CheckUserAlreadyLikedMovieFindDTO{})){
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.CheckUserAlreadyLikedMovieFindDTO).UserID != model.UserID(trueCase.userId){
				t.Fatal("UserID Not Match,")
			}
			if data.(*usecase.CheckUserAlreadyLikedMovieFindDTO).MovieID != model.MovieID(trueCase.movieId){
				t.Fatal("MovieID Not Match,")
			}
			return nil
		})

		checkUserAlreadyLikedMovieHandler.CheckUserAlreadyLikedMovie(ginContext)
	}
}
