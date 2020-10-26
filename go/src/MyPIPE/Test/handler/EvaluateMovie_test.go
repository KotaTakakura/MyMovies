package handler

import (
	mock_repository "MyPIPE/Test/mock/repository"
	mock_usecase "MyPIPE/Test/mock/usecase"
	"MyPIPE/domain/model"
	"MyPIPE/handler"
	"MyPIPE/usecase"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"errors"
)

func TestEvaluateMovie(t *testing.T) {

	trueCases := []struct {
		userId   uint64
		movieId  uint64
		evaluate string
	}{
		{userId: 10, movieId: 100, evaluate: "good"},
		{userId: 10, movieId: 100, evaluate: "bad"},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	movieEvaluationRepository := mock_repository.NewMockMovieEvaluationRepository(ctrl)
	evaluateMovieUsecase := mock_usecase.NewMockIEvaluateMovie(ctrl)
	evaluateMovieHandler := handler.NewEvaluateMovie(movieRepository, movieEvaluationRepository, evaluateMovieUsecase)

	for _, trueCase := range trueCases {
		// ポストデータ
		bodyReader := strings.NewReader(`{"movie_id": ` + fmt.Sprint(trueCase.movieId) + `,"evaluate":"` + trueCase.evaluate + `"}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(trueCase.userId),
		})

		evaluateMovieUsecase.EXPECT().EvaluateMovie(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.EvaluateMovieDTO{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.EvaluateMovieDTO).UserID != model.UserID(trueCase.userId) {
				t.Fatal("UserID Not Match,")
			}
			if data.(*usecase.EvaluateMovieDTO).MovieID != model.MovieID(trueCase.movieId) {
				t.Fatal("MovieID Not Match,")
			}
			evaluation, _ := model.NewEvaluation(trueCase.evaluate)
			if data.(*usecase.EvaluateMovieDTO).Evaluation != evaluation {
				t.Fatal("Evaluation Not Match,")
			}
			return nil
		})

		evaluateMovieHandler.EvaluateMovie(ginContext)
	}

	falseCases := []struct {
		userId   uint64
		movieId  uint64
		evaluate string
	}{
		{userId: 10, movieId: 100, evaluate: "aaaaaa"},
	}

	for _, falseCase := range falseCases {
		// ポストデータ
		bodyReader := strings.NewReader(`{"movie_id": ` + fmt.Sprint(falseCase.movieId) + `,"evaluate":"` + falseCase.evaluate + `"}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(falseCase.userId),
		})

		evaluateMovieHandler.EvaluateMovie(ginContext)
	}
}

func TestEvaluateMovie_MovieIDError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	movieEvaluationRepository := mock_repository.NewMockMovieEvaluationRepository(ctrl)
	evaluateMovieUsecase := mock_usecase.NewMockIEvaluateMovie(ctrl)
	evaluateMovieHandler := handler.NewEvaluateMovie(movieRepository, movieEvaluationRepository, evaluateMovieUsecase)

	cases := []struct {
		userId   uint64
		movieId  uint64
		evaluate string
	}{
		{userId: 10, movieId: 100, evaluate: "good"},
		{userId: 10, movieId: 100, evaluate: "bad"},
	}

	for _,Case := range cases {
		// ポストデータ
		//moviIDが無効
		bodyReader := strings.NewReader(`{"movie_id": "INVALIDMOVIEID","evaluate":"` + Case.evaluate + `"}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(Case.userId),
		})

		evaluateMovieHandler.EvaluateMovie(ginContext)

		if w.Code != http.StatusBadRequest{
			t.Fatal("Error.")
		}
	}
}

func TestEvaluateMovie_UsecaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	movieEvaluationRepository := mock_repository.NewMockMovieEvaluationRepository(ctrl)
	evaluateMovieUsecase := mock_usecase.NewMockIEvaluateMovie(ctrl)
	evaluateMovieHandler := handler.NewEvaluateMovie(movieRepository, movieEvaluationRepository, evaluateMovieUsecase)

	cases := []struct {
		userId   uint64
		movieId  uint64
		evaluate string
	}{
		{userId: 10, movieId: 100, evaluate: "good"},
		{userId: 10, movieId: 100, evaluate: "bad"},
	}

	for _,Case := range cases {
		// ポストデータ
		bodyReader := strings.NewReader(`{"movie_id": ` + fmt.Sprint(Case.movieId) + `,"evaluate":"` + Case.evaluate + `"}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(Case.userId),
		})

		evaluateMovieUsecase.EXPECT().EvaluateMovie(gomock.Any()).Return(errors.New("EvaluateMovieUsecase Error."))

		evaluateMovieHandler.EvaluateMovie(ginContext)

		if w.Code != http.StatusInternalServerError{
			t.Fatal("Error.")
		}
	}
}
