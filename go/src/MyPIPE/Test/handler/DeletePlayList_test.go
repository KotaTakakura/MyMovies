package handler

import (
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
	"testing"
)

func TestDeletePlayList(t *testing.T) {

	trueCases := []struct {
		userId     uint64
		playListId uint64
	}{
		{userId: 10, playListId: 100},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListRepository := mock_repository.NewMockPlayListRepository(ctrl)
	deletePlayListUsecase := mock_usecase.NewMockIDeletePlayList(ctrl)
	deletePlayListHandler := handler.NewDeletePlayList(playListRepository, deletePlayListUsecase)

	for _, trueCase := range trueCases {

		// リクエスト生成
		req := httptest.NewRequest("GET", "/?play_list_id="+fmt.Sprint(trueCase.playListId), nil)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(trueCase.userId),
		})

		deletePlayListUsecase.EXPECT().Delete(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.DeletePlayListDTO{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.DeletePlayListDTO).UserID != model.UserID(trueCase.userId) {
				t.Fatal("UserID Not Match,")
			}
			if data.(*usecase.DeletePlayListDTO).PlaylistID != model.PlayListID(trueCase.playListId) {
				t.Fatal("PlayListID Not Match,")
			}
			return nil
		})

		deletePlayListHandler.DeletePlayList(ginContext)
	}
}

func TestDeletePlayList_PlayListIDError(t *testing.T) {
	cases := []struct {
		userId     uint64
		playListId uint64
	}{
		{userId: 10, playListId: 100},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListRepository := mock_repository.NewMockPlayListRepository(ctrl)
	deletePlayListUsecase := mock_usecase.NewMockIDeletePlayList(ctrl)
	deletePlayListHandler := handler.NewDeletePlayList(playListRepository, deletePlayListUsecase)

	for _, Case := range cases {

		// リクエスト生成
		//play_list_idが無効
		req := httptest.NewRequest("GET", "/?play_list_id=THISISNOTPLAYLISTID", nil)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(Case.userId),
		})

		deletePlayListHandler.DeletePlayList(ginContext)

		if w.Code != http.StatusBadRequest {
			t.Fatal("Error.")
		}
	}

	for _, Case := range cases {

		// リクエスト生成
		req := httptest.NewRequest("GET", "/", nil)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		//play_list_idが存在しない
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(Case.userId),
		})

		deletePlayListHandler.DeletePlayList(ginContext)

		if w.Code != http.StatusBadRequest {
			t.Fatal("Error.")
		}
	}
}

func TestDeletePlayList_UsecaseError(t *testing.T) {
	cases := []struct {
		userId     uint64
		playListId uint64
	}{
		{userId: 10, playListId: 100},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListRepository := mock_repository.NewMockPlayListRepository(ctrl)
	deletePlayListUsecase := mock_usecase.NewMockIDeletePlayList(ctrl)
	deletePlayListHandler := handler.NewDeletePlayList(playListRepository, deletePlayListUsecase)

	for _, Case := range cases {

		// リクエスト生成
		//play_list_idが無効
		req := httptest.NewRequest("GET", "/?play_list_id="+fmt.Sprint(Case.playListId), nil)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(Case.userId),
		})

		deletePlayListUsecase.EXPECT().Delete(gomock.Any()).Return(errors.New("DeletePlayListUsecase Error."))

		deletePlayListHandler.DeletePlayList(ginContext)

		if w.Code != http.StatusInternalServerError {
			t.Fatal("Error.")
		}
	}
}
