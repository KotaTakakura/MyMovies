package handler

import (
	mock_factory "MyPIPE/Test/mock/factory"
	mock_repository "MyPIPE/Test/mock/repository"
	mock_usecase "MyPIPE/Test/mock/usecase"
	"MyPIPE/domain/model"
	"MyPIPE/handler"
	"MyPIPE/usecase"
	"errors"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestPlayListItemAddPlayListMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListRepsitory := mock_repository.NewMockPlayListRepository(ctrl)
	playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
	playListMovieFactory := mock_factory.NewMockIPlayListMovie(ctrl)
	addPlayListItemAddUsecase := mock_usecase.NewMockIAddPlayListItem(ctrl)
	deletePlayListMovieUsecase := mock_usecase.NewMockIDeletePlayListMovie(ctrl)
	playListItemHandler := handler.NewPlayListItem(playListRepsitory, playListMovieRepository, playListMovieFactory, addPlayListItemAddUsecase, deletePlayListMovieUsecase)

	trueCases := []struct {
		userId     uint64
		movieId    uint64
		playListId uint64
	}{
		{userId: 10, movieId: 100, playListId: 200},
	}

	for _, trueCase := range trueCases {
		// ポストデータ
		bodyReader := strings.NewReader(`{"movie_id":` + fmt.Sprint(trueCase.movieId) + `,"play_list_id":` + fmt.Sprint(trueCase.playListId) + `}`)

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

		addPlayListItemAddUsecase.EXPECT().AddPlayListItem(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.AddPlayListItemAddJson{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.AddPlayListItemAddJson).UserID != model.UserID(trueCase.userId) {
				t.Fatal("UserID Not Match,")
			}
			if data.(*usecase.AddPlayListItemAddJson).MovieID != model.MovieID(trueCase.movieId) {
				t.Fatal("MovieID Not Match,")
			}
			if data.(*usecase.AddPlayListItemAddJson).PlayListID != model.PlayListID(trueCase.playListId) {
				t.Fatal("PlayListID Not Match,")
			}
			return nil
		})

		playListItemHandler.AddPlayListMovie(ginContext)
	}
}

func TestPlayListItemAddPlayListMovie_ParametersError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListRepsitory := mock_repository.NewMockPlayListRepository(ctrl)
	playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
	playListMovieFactory := mock_factory.NewMockIPlayListMovie(ctrl)
	addPlayListItemAddUsecase := mock_usecase.NewMockIAddPlayListItem(ctrl)
	deletePlayListMovieUsecase := mock_usecase.NewMockIDeletePlayListMovie(ctrl)
	playListItemHandler := handler.NewPlayListItem(playListRepsitory, playListMovieRepository, playListMovieFactory, addPlayListItemAddUsecase, deletePlayListMovieUsecase)

	cases := []struct {
		userId     uint64
		movieId    uint64
		playListId uint64
	}{
		{userId: 10, movieId: 0, playListId: 200},
		{userId: 10, movieId: 10, playListId: 0},
	}

	for _, Case := range cases {
		// ポストデータ
		bodyReader := strings.NewReader(`{"movie_id":` + fmt.Sprint(Case.movieId) + `,"play_list_id":` + fmt.Sprint(Case.playListId) + `}`)

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

		playListItemHandler.AddPlayListMovie(ginContext)

		if w.Code != http.StatusBadRequest {
			t.Fatal("Error")
		}
	}
}

func TestPlayListItemAddPlayListMovie_UsecaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListRepsitory := mock_repository.NewMockPlayListRepository(ctrl)
	playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
	playListMovieFactory := mock_factory.NewMockIPlayListMovie(ctrl)
	addPlayListItemAddUsecase := mock_usecase.NewMockIAddPlayListItem(ctrl)
	deletePlayListMovieUsecase := mock_usecase.NewMockIDeletePlayListMovie(ctrl)
	playListItemHandler := handler.NewPlayListItem(playListRepsitory, playListMovieRepository, playListMovieFactory, addPlayListItemAddUsecase, deletePlayListMovieUsecase)

	trueCases := []struct {
		userId     uint64
		movieId    uint64
		playListId uint64
	}{
		{userId: 10, movieId: 100, playListId: 200},
	}

	for _, trueCase := range trueCases {
		// ポストデータ
		bodyReader := strings.NewReader(`{"movie_id":` + fmt.Sprint(trueCase.movieId) + `,"play_list_id":` + fmt.Sprint(trueCase.playListId) + `}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(trueCase.userId),
		})

		addPlayListItemAddUsecase.EXPECT().AddPlayListItem(gomock.Any()).Return(errors.New("Usecase Error"))

		playListItemHandler.AddPlayListMovie(ginContext)

		if w.Code != http.StatusInternalServerError {
			t.Fatal("Error.")
		}
	}
}

func TestPlayListItemDeletePlayListMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListRepsitory := mock_repository.NewMockPlayListRepository(ctrl)
	playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
	playListMovieFactory := mock_factory.NewMockIPlayListMovie(ctrl)
	addPlayListItemAddUsecase := mock_usecase.NewMockIAddPlayListItem(ctrl)
	deletePlayListMovieUsecase := mock_usecase.NewMockIDeletePlayListMovie(ctrl)
	playListItemHandler := handler.NewPlayListItem(playListRepsitory, playListMovieRepository, playListMovieFactory, addPlayListItemAddUsecase, deletePlayListMovieUsecase)

	trueCases := []struct {
		userId     uint64
		movieId    uint64
		playListId uint64
	}{
		{userId: 10, movieId: 100, playListId: 200},
	}

	for _, trueCase := range trueCases {
		// ポストデータ
		bodyReader := strings.NewReader(``)

		// リクエスト生成
		req := httptest.NewRequest("DELETE", `/?movie_id=`+fmt.Sprint(trueCase.movieId)+`&play_list_id=`+fmt.Sprint(trueCase.playListId), bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(trueCase.userId),
		})

		deletePlayListMovieUsecase.EXPECT().DeletePlayListItem(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.DeletePlayListMovieDTO{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.DeletePlayListMovieDTO).UserID != model.UserID(trueCase.userId) {
				t.Fatal("UserID Not Match,")
			}
			if data.(*usecase.DeletePlayListMovieDTO).MovieID != model.MovieID(trueCase.movieId) {
				t.Fatal("MovieID Not Match,")
			}
			if data.(*usecase.DeletePlayListMovieDTO).PlayListID != model.PlayListID(trueCase.playListId) {
				t.Fatal("PlayListID Not Match,")
			}
			return nil
		})

		playListItemHandler.DeletePlayListMovie(ginContext)
	}
}

func TestPlayListItemDeletePlayListMovie_ParametersError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListRepsitory := mock_repository.NewMockPlayListRepository(ctrl)
	playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
	playListMovieFactory := mock_factory.NewMockIPlayListMovie(ctrl)
	addPlayListItemAddUsecase := mock_usecase.NewMockIAddPlayListItem(ctrl)
	deletePlayListMovieUsecase := mock_usecase.NewMockIDeletePlayListMovie(ctrl)
	playListItemHandler := handler.NewPlayListItem(playListRepsitory, playListMovieRepository, playListMovieFactory, addPlayListItemAddUsecase, deletePlayListMovieUsecase)

	cases := []struct {
		userId     uint64
		movieId    uint64
		playListId uint64
	}{
		{userId: 10, movieId: 0, playListId: 200},
		{userId: 10, movieId: 10, playListId: 0},
	}

	for _, Case := range cases {
		// ポストデータ
		bodyReader := strings.NewReader(``)

		// リクエスト生成
		req := httptest.NewRequest("DELETE", `/?movie_id=`+fmt.Sprint(Case.movieId)+`&play_list_id=`+fmt.Sprint(Case.playListId), bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(Case.userId),
		})

		playListItemHandler.DeletePlayListMovie(ginContext)

		if w.Code != http.StatusBadRequest {
			t.Fatal("Error.")
		}
	}
}

func TestPlayListItemDeletePlayListMovie_UsecaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListRepsitory := mock_repository.NewMockPlayListRepository(ctrl)
	playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
	playListMovieFactory := mock_factory.NewMockIPlayListMovie(ctrl)
	addPlayListItemAddUsecase := mock_usecase.NewMockIAddPlayListItem(ctrl)
	deletePlayListMovieUsecase := mock_usecase.NewMockIDeletePlayListMovie(ctrl)
	playListItemHandler := handler.NewPlayListItem(playListRepsitory, playListMovieRepository, playListMovieFactory, addPlayListItemAddUsecase, deletePlayListMovieUsecase)

	trueCases := []struct {
		userId     uint64
		movieId    uint64
		playListId uint64
	}{
		{userId: 10, movieId: 100, playListId: 200},
	}

	for _, trueCase := range trueCases {
		// ポストデータ
		bodyReader := strings.NewReader(``)

		// リクエスト生成
		req := httptest.NewRequest("DELETE", `/?movie_id=`+fmt.Sprint(trueCase.movieId)+`&play_list_id=`+fmt.Sprint(trueCase.playListId), bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(trueCase.userId),
		})

		deletePlayListMovieUsecase.EXPECT().DeletePlayListItem(gomock.Any()).Return(errors.New("Usecase Error"))
		playListItemHandler.DeletePlayListMovie(ginContext)
		if w.Code != http.StatusInternalServerError {
			t.Fatal("Error")
		}
	}

}
