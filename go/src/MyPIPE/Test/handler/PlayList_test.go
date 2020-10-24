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
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestPlayList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_repository.NewMockUserRepository(ctrl)
	playListRepository := mock_repository.NewMockPlayListRepository(ctrl)
	createPlayListUsecase := mock_usecase.NewMockICreatePlayList(ctrl)
	createPlayListHandler := handler.NewCreatePlayList(userRepository, playListRepository, createPlayListUsecase)

	trueCases := []struct {
		userId              uint64
		playListName        string
		playListDescription string
	}{
		{userId: 10, playListName: "TestPlaylistName", playListDescription: "TestPlayListDescription"},
	}

	falseCases := []struct {
		userId              uint64
		playListName        string
		playListDescription string
	}{
		{userId: 10, playListName: "", playListDescription: "TestPlayListDescription"},
	}

	for _, trueCase := range trueCases {
		// ポストデータ
		bodyReader := strings.NewReader(`
			{
				"play_list_name":"` + fmt.Sprint(trueCase.playListName) + `",
				"play_list_description":"` + fmt.Sprint(trueCase.playListDescription) + `"
			}
`)

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

		createPlayListUsecase.EXPECT().CreatePlayList(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.CreatePlayListDTO{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.CreatePlayListDTO).UserID != model.UserID(trueCase.userId) {
				t.Fatal("UserID Not Match,")
			}
			if data.(*usecase.CreatePlayListDTO).PlayListName != model.PlayListName(trueCase.playListName) {
				t.Fatal("PlayListName Not Match,")
			}
			if data.(*usecase.CreatePlayListDTO).PlayListDescription != model.PlayListDescription(trueCase.playListDescription) {
				t.Fatal("PlayListDescription Not Match,")
			}
			return nil
		})

		createPlayListHandler.CreatePlayList(ginContext)
	}

	for _, falseCase := range falseCases {
		// ポストデータ
		bodyReader := strings.NewReader(`
			{
				"play_list_name":"` + fmt.Sprint(falseCase.playListName) + `",
				"play_list_description":"` + fmt.Sprint(falseCase.playListDescription) + `"
			}
`)

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

		createPlayListHandler.CreatePlayList(ginContext)
	}
}
