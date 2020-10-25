package handler

import (
	mock_repository "MyPIPE/Test/mock/repository"
	mock_usecase "MyPIPE/Test/mock/usecase"
	"MyPIPE/domain/model"
	"MyPIPE/handler"
	"MyPIPE/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestCheckUserAlreadyLikedMovie(t *testing.T) {
	trueCases := []struct {
		userId       uint64
		movieId      uint64
		alreadyLiked bool
	}{
		{userId: 10, movieId: 100, alreadyLiked: true},
		{userId: 99999, movieId: 999, alreadyLiked: false},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieEvaluationRepository := mock_repository.NewMockMovieEvaluationRepository(ctrl)
	checkAlreadyLikeDMovieUsecase := mock_usecase.NewMockICheckUserAlreadyLikedMovie(ctrl)
	checkUserAlreadyLikedMovieHandler := handler.NewCheckUserAlreadyLikedMovie(movieEvaluationRepository, checkAlreadyLikeDMovieUsecase)

	for _, trueCase := range trueCases {
		// ポストデータ
		bodyReader := strings.NewReader("")

		// リクエスト生成
		req := httptest.NewRequest("GET", "/?user_id="+fmt.Sprint(trueCase.userId)+"&movie_id="+fmt.Sprint(trueCase.movieId), bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req

		checkAlreadyLikeDMovieUsecase.EXPECT().Find(gomock.Any()).DoAndReturn(func(data interface{}) bool {
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.CheckUserAlreadyLikedMovieFindDTO{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.CheckUserAlreadyLikedMovieFindDTO).UserID != model.UserID(trueCase.userId) {
				t.Fatal("UserID Not Match,")
			}
			if data.(*usecase.CheckUserAlreadyLikedMovieFindDTO).MovieID != model.MovieID(trueCase.movieId) {
				t.Fatal("MovieID Not Match,")
			}
			return trueCase.alreadyLiked
		})

		checkUserAlreadyLikedMovieHandler.CheckUserAlreadyLikedMovie(ginContext)
	}
}

func TestCheckUserAlreadyLikedMovie_UserIDError(t *testing.T) {
	trueCases := []struct {
		userId  uint64
		movieId uint64
	}{
		{userId: 10, movieId: 100},
		{userId: 99999, movieId: 999},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieEvaluationRepository := mock_repository.NewMockMovieEvaluationRepository(ctrl)
	checkAlreadyLikeDMovieUsecase := mock_usecase.NewMockICheckUserAlreadyLikedMovie(ctrl)
	checkUserAlreadyLikedMovieHandler := handler.NewCheckUserAlreadyLikedMovie(movieEvaluationRepository, checkAlreadyLikeDMovieUsecase)

	for _, trueCase := range trueCases {
		// ポストデータ
		bodyReader := strings.NewReader("")

		// リクエスト生成
		//UserIDが設定されていない
		req := httptest.NewRequest("GET", "/?movie_id="+fmt.Sprint(trueCase.movieId), bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req

		checkUserAlreadyLikedMovieHandler.CheckUserAlreadyLikedMovie(ginContext)
	}

	for _, trueCase := range trueCases {
		// ポストデータ
		bodyReader := strings.NewReader("")

		// リクエスト生成
		//UserIDが無効
		req := httptest.NewRequest("GET", "/?user_id=THISISNOTUSERID&movie_id="+fmt.Sprint(trueCase.movieId), bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req

		checkUserAlreadyLikedMovieHandler.CheckUserAlreadyLikedMovie(ginContext)
	}
}

func TestCheckUserAlreadyLikedMovie_MovieIDError(t *testing.T) {
	trueCases := []struct {
		userId  uint64
		movieId uint64
	}{
		{userId: 10, movieId: 100},
		{userId: 99999, movieId: 999},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieEvaluationRepository := mock_repository.NewMockMovieEvaluationRepository(ctrl)
	checkAlreadyLikeDMovieUsecase := mock_usecase.NewMockICheckUserAlreadyLikedMovie(ctrl)
	checkUserAlreadyLikedMovieHandler := handler.NewCheckUserAlreadyLikedMovie(movieEvaluationRepository, checkAlreadyLikeDMovieUsecase)

	for _, trueCase := range trueCases {
		// ポストデータ
		bodyReader := strings.NewReader("")

		// リクエスト生成
		//MovieIDが設定されていない
		req := httptest.NewRequest("GET", "/?user_id="+fmt.Sprint(trueCase.userId), bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req

		checkUserAlreadyLikedMovieHandler.CheckUserAlreadyLikedMovie(ginContext)
	}

	for _, trueCase := range trueCases {
		// ポストデータ
		bodyReader := strings.NewReader("")

		// リクエスト生成
		//MovieIDが無効
		req := httptest.NewRequest("GET", "/?user_id="+fmt.Sprint(trueCase.userId)+"&movie_id=THISISNOTMOVIEID", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req

		checkUserAlreadyLikedMovieHandler.CheckUserAlreadyLikedMovie(ginContext)
	}
}
