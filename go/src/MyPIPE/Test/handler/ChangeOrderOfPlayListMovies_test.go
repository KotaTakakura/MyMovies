package handler

import (
	mock_repository "MyPIPE/Test/mock/repository"
	mock_usecase "MyPIPE/Test/mock/usecase"
	"MyPIPE/domain/model"
	"MyPIPE/handler"
	"MyPIPE/usecase"
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestChangeOrderOfPlayListMovies(t *testing.T) {

	trueCases := []struct {
		playListId          uint64
		userId              uint64
		playListMovieId1    uint64
		playListMovieOrder1 uint64
		playListMovieId2    uint64
		playListMovieOrder2 uint64
		playListMovieId3    uint64
		playListMovieOrder3 uint64
	}{
		{
			playListId:          10,
			userId:              200,
			playListMovieId1:    1,
			playListMovieOrder1: 1,
			playListMovieId2:    2,
			playListMovieOrder2: 2,
			playListMovieId3:    3,
			playListMovieOrder3: 3,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
	changeOrderOfPlayListMoviesUsecase := mock_usecase.NewMockIChangeOrderOfPlayListMovies(ctrl)
	changeOrderOfPlayListMoviesHandler := handler.NewChangeOrderOfPlayListMovies(playListMovieRepository, changeOrderOfPlayListMoviesUsecase)

	for _, trueCase := range trueCases {

		// ポストデータ
		bodyReader := strings.NewReader(
			`{
				  "play_list_id":` + strconv.FormatUint(trueCase.playListId, 10) + `,
				  "play_list_movie_id_and_order":[
					{
					  "play_list_movie_id":` + strconv.FormatUint(trueCase.playListMovieId1, 10) + `,
					  "play_lise_movie_order":` + strconv.FormatUint(trueCase.playListMovieOrder1, 10) + `
					},
					{
					  "play_list_movie_id":` + strconv.FormatUint(trueCase.playListMovieId2, 10) + `,
					  "play_lise_movie_order":` + strconv.FormatUint(trueCase.playListMovieOrder2, 10) + `
					},
					{
					  "play_list_movie_id":` + strconv.FormatUint(trueCase.playListMovieId3, 10) + `,
					  "play_lise_movie_order":` + strconv.FormatUint(trueCase.playListMovieOrder3, 10) + `
					}
				  ]
				}`)

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

		changeOrderOfPlayListMoviesUsecase.EXPECT().ChangeOrderOfPlayListMovies(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.ChangeOrderOfPlayListMoviesDTO{})) {
				t.Fatal("Type Not Match.")
			}

			if data.(*usecase.ChangeOrderOfPlayListMoviesDTO).UserID != model.UserID(trueCase.userId) {
				t.Fatal("UserID Not Match.")
			}

			if data.(*usecase.ChangeOrderOfPlayListMoviesDTO).PlayListID != model.PlayListID(trueCase.playListId) {
				t.Fatal("PlayListID Not Match.")
			}

			for _, v := range data.(*usecase.ChangeOrderOfPlayListMoviesDTO).MovieIDAndOrder {
				if v.MovieID == model.MovieID(trueCase.playListMovieId1) && v.Order == model.PlayListMovieOrder(trueCase.playListMovieOrder1) {
					continue
				}
				if v.MovieID == model.MovieID(trueCase.playListMovieId2) && v.Order == model.PlayListMovieOrder(trueCase.playListMovieOrder2) {
					continue
				}
				if v.MovieID == model.MovieID(trueCase.playListMovieId3) && v.Order == model.PlayListMovieOrder(trueCase.playListMovieOrder3) {
					continue
				}
				t.Fatal("PlayListOrder Not Match.")
			}
			return nil
		})
		changeOrderOfPlayListMoviesHandler.ChangeOrderOfPlayListMovies(ginContext)
	}
}

func TestChangeOrderOfPlayListMovies_MovieIdError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
	changeOrderOfPlayListMoviesUsecase := mock_usecase.NewMockIChangeOrderOfPlayListMovies(ctrl)
	changeOrderOfPlayListMoviesHandler := handler.NewChangeOrderOfPlayListMovies(playListMovieRepository, changeOrderOfPlayListMoviesUsecase)

	trueCases := []struct {
		playListId          uint64
		userId              uint64
		playListMovieId1    uint64
		playListMovieOrder1 uint64
		playListMovieId2    uint64
		playListMovieOrder2 uint64
		playListMovieId3    uint64
		playListMovieOrder3 uint64
	}{
		{
			playListId:          10,
			userId:              200,
			playListMovieId1:    1,
			playListMovieOrder1: 1,
			playListMovieId2:    2,
			playListMovieOrder2: 2,
			playListMovieId3:    3,
			playListMovieOrder3: 3,
		},
	}

	//MovieIDのエラー
	falseMovieIDCases := trueCases
	for _, falseCase := range falseMovieIDCases {

		// ポストデータ
		//MovieIDが文字列になっている
		bodyReader := strings.NewReader(
			`{
				  "play_list_id":` + strconv.FormatUint(falseCase.playListId, 10) + `,
				  "play_list_movie_id_and_order":[
					{
					  "play_list_movie_id":"Fail",
					  "play_lise_movie_order":` + strconv.FormatUint(falseCase.playListMovieOrder1, 10) + `
					},
					{
					  "play_list_movie_id":` + strconv.FormatUint(falseCase.playListMovieId2, 10) + `,
					  "play_lise_movie_order":` + strconv.FormatUint(falseCase.playListMovieOrder2, 10) + `
					},
					{
					  "play_list_movie_id":` + strconv.FormatUint(falseCase.playListMovieId3, 10) + `,
					  "play_lise_movie_order":` + strconv.FormatUint(falseCase.playListMovieOrder3, 10) + `
					}
				  ]
				}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(falseCase.userId),
		})

		changeOrderOfPlayListMoviesHandler.ChangeOrderOfPlayListMovies(ginContext)

		if w.Code != http.StatusBadRequest {
			t.Fatal("Error.")
		}
	}
}

func TestChangeOrderOfPlayListMovies_OrderError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
	changeOrderOfPlayListMoviesUsecase := mock_usecase.NewMockIChangeOrderOfPlayListMovies(ctrl)
	changeOrderOfPlayListMoviesHandler := handler.NewChangeOrderOfPlayListMovies(playListMovieRepository, changeOrderOfPlayListMoviesUsecase)

	trueCases := []struct {
		playListId          uint64
		userId              uint64
		playListMovieId1    uint64
		playListMovieOrder1 uint64
		playListMovieId2    uint64
		playListMovieOrder2 uint64
		playListMovieId3    uint64
		playListMovieOrder3 uint64
	}{
		{
			playListId:          10,
			userId:              200,
			playListMovieId1:    1,
			playListMovieOrder1: 1,
			playListMovieId2:    2,
			playListMovieOrder2: 2,
			playListMovieId3:    3,
			playListMovieOrder3: 3,
		},
	}
	//Orderのエラー
	falseOrderCases := trueCases
	for _, falseCase := range falseOrderCases {

		// ポストデータ
		//Orderが文字列になっている
		bodyReader := strings.NewReader(
			`{
				  "play_list_id":` + strconv.FormatUint(falseCase.playListId, 10) + `,
				  "play_list_movie_id_and_order":[
					{
					  "play_list_movie_id":` + strconv.FormatUint(falseCase.playListMovieId1, 10) + `,
					  "play_lise_movie_order":"Fail"
					},
					{
					  "play_list_movie_id":` + strconv.FormatUint(falseCase.playListMovieId2, 10) + `,
					  "play_lise_movie_order":` + strconv.FormatUint(falseCase.playListMovieOrder2, 10) + `
					},
					{
					  "play_list_movie_id":` + strconv.FormatUint(falseCase.playListMovieId3, 10) + `,
					  "play_lise_movie_order":` + strconv.FormatUint(falseCase.playListMovieOrder3, 10) + `
					}
				  ]
				}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(falseCase.userId),
		})

		changeOrderOfPlayListMoviesHandler.ChangeOrderOfPlayListMovies(ginContext)

		if w.Code != http.StatusBadRequest {
			t.Fatal("Error.")
		}
	}
}

func TestChangeOrderOfPlayListMovies_PlayListIDError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
	changeOrderOfPlayListMoviesUsecase := mock_usecase.NewMockIChangeOrderOfPlayListMovies(ctrl)
	changeOrderOfPlayListMoviesHandler := handler.NewChangeOrderOfPlayListMovies(playListMovieRepository, changeOrderOfPlayListMoviesUsecase)

	trueCases := []struct {
		playListId          uint64
		userId              uint64
		playListMovieId1    uint64
		playListMovieOrder1 uint64
		playListMovieId2    uint64
		playListMovieOrder2 uint64
		playListMovieId3    uint64
		playListMovieOrder3 uint64
	}{
		{
			playListId:          10,
			userId:              200,
			playListMovieId1:    1,
			playListMovieOrder1: 1,
			playListMovieId2:    2,
			playListMovieOrder2: 2,
			playListMovieId3:    3,
			playListMovieOrder3: 3,
		},
	}

	//PlayListIDのエラー
	falsePlayListIDCases := trueCases
	for _, falseCase := range falsePlayListIDCases {

		// ポストデータ
		//PlayListIDが文字列になっている
		bodyReader := strings.NewReader(
			`{
				  "play_list_id":"10",
				  "play_list_movie_id_and_order":[
					{
					  "play_list_movie_id":` + strconv.FormatUint(falseCase.playListMovieId1, 10) + `,
					 "play_lise_movie_order":` + strconv.FormatUint(falseCase.playListMovieOrder1, 10) + `
					},
					{
					  "play_list_movie_id":` + strconv.FormatUint(falseCase.playListMovieId2, 10) + `,
					  "play_lise_movie_order":` + strconv.FormatUint(falseCase.playListMovieOrder2, 10) + `
					},
					{
					  "play_list_movie_id":` + strconv.FormatUint(falseCase.playListMovieId3, 10) + `,
					  "play_lise_movie_order":` + strconv.FormatUint(falseCase.playListMovieOrder3, 10) + `
					}
				  ]
				}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(falseCase.userId),
		})

		changeOrderOfPlayListMoviesHandler.ChangeOrderOfPlayListMovies(ginContext)

		if w.Code != http.StatusBadRequest {
			t.Fatal("Error.")
		}
	}
}

func TestChangeOrderOfPlayListMovies_UsecaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
	changeOrderOfPlayListMoviesUsecase := mock_usecase.NewMockIChangeOrderOfPlayListMovies(ctrl)
	changeOrderOfPlayListMoviesHandler := handler.NewChangeOrderOfPlayListMovies(playListMovieRepository, changeOrderOfPlayListMoviesUsecase)

	trueCases := []struct {
		playListId          uint64
		userId              uint64
		playListMovieId1    uint64
		playListMovieOrder1 uint64
		playListMovieId2    uint64
		playListMovieOrder2 uint64
		playListMovieId3    uint64
		playListMovieOrder3 uint64
	}{
		{
			playListId:          10,
			userId:              200,
			playListMovieId1:    1,
			playListMovieOrder1: 1,
			playListMovieId2:    2,
			playListMovieOrder2: 2,
			playListMovieId3:    3,
			playListMovieOrder3: 3,
		},
	}
	//Usecaseのエラー
	falseChangeOrderOfPlayListMoviesUsecaseCases := trueCases

	for _, falseCase := range falseChangeOrderOfPlayListMoviesUsecaseCases {

		// ポストデータ
		bodyReader := strings.NewReader(
			`{
				  "play_list_id":` + strconv.FormatUint(falseCase.playListId, 10) + `,
				  "play_list_movie_id_and_order":[
					{
					  "play_list_movie_id":` + strconv.FormatUint(falseCase.playListMovieId1, 10) + `,
					  "play_lise_movie_order":` + strconv.FormatUint(falseCase.playListMovieOrder1, 10) + `
					},
					{
					  "play_list_movie_id":` + strconv.FormatUint(falseCase.playListMovieId2, 10) + `,
					  "play_lise_movie_order":` + strconv.FormatUint(falseCase.playListMovieOrder2, 10) + `
					},
					{
					  "play_list_movie_id":` + strconv.FormatUint(falseCase.playListMovieId3, 10) + `,
					  "play_lise_movie_order":` + strconv.FormatUint(falseCase.playListMovieOrder3, 10) + `
					}
				  ]
				}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(falseCase.userId),
		})

		changeOrderOfPlayListMoviesUsecase.EXPECT().ChangeOrderOfPlayListMovies(gomock.Any()).Return(errors.New("ChangeOrderOfPlayListMovies Error"))
		changeOrderOfPlayListMoviesHandler.ChangeOrderOfPlayListMovies(ginContext)

		if w.Code != http.StatusInternalServerError {
			t.Fatal("Error.")
		}
	}
}
