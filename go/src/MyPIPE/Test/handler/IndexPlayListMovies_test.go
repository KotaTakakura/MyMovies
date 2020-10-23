package handler

import (
	mock_queryService "MyPIPE/Test/mock/queryService"
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

func TestIndexPlayListMovies(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	indexPlayListMoviesQueryService := mock_queryService.NewMockIndexPlayListMovieQueryService(ctrl)
	indexPlayListMoviesUsecase := mock_usecase.NewMockIIndexPlaylistItemInMyPage(ctrl)
	indexPlayListMoviesHandler := handler.NewIndexPlaylistMovies(indexPlayListMoviesQueryService, indexPlayListMoviesUsecase)

	trueCases := []struct {
		userId     uint64
		playListId uint64
	}{
		{userId: 10, playListId: 100},
	}

	for _, trueCase := range trueCases {
		// ポストデータ
		bodyReader := strings.NewReader("")

		// リクエスト生成
		req := httptest.NewRequest("GET", "/play-list-items", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req
		ginContext.Params = append(ginContext.Params, gin.Param{Key: "play_list_id", Value: fmt.Sprint(trueCase.playListId)})
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(trueCase.userId),
		})

		indexPlayListMoviesUsecase.EXPECT().Find(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.IndexPlayListItemInMyPageDTO{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.IndexPlayListItemInMyPageDTO).UserID != model.UserID(trueCase.userId) {
				t.Fatal("UserID Not Match,")
			}
			if data.(*usecase.IndexPlayListItemInMyPageDTO).PlayListID != model.PlayListID(trueCase.playListId) {
				t.Fatal("PlayListID Not Match,")
			}
			return nil
		})
		indexPlayListMoviesHandler.IndexPlaylistMovies(ginContext)

	}
}
