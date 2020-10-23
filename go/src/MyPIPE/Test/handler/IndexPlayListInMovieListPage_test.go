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

func TestIndexPlayListInMovieListPage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	indexPlayListInMovieListPageQueryService := mock_queryService.NewMockIndexPlayListInMovieListPageQueryService(ctrl)
	indexPlayListInMovieListPageUsecase := mock_usecase.NewMockIIndexPlayListInMovieListPage(ctrl)
	indexPlayListInMovieListPageHandler := handler.NewIndexPlayListInMovieListPage(indexPlayListInMovieListPageQueryService, indexPlayListInMovieListPageUsecase)

	trueCases := []struct {
		userId  uint64
		movieId uint64
	}{
		{userId: 10, movieId: 100},
	}

	for _, trueCase := range trueCases {

		// ポストデータ
		bodyReader := strings.NewReader("")

		// リクエスト生成
		req := httptest.NewRequest("GET", "/play-lists/10", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req
		ginContext.Params = append(ginContext.Params, gin.Param{Key: "movie_id", Value: fmt.Sprint(trueCase.movieId)})
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(trueCase.userId),
		})

		indexPlayListInMovieListPageUsecase.EXPECT().Find(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.FindDTO{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.FindDTO).UserID != model.UserID(trueCase.userId) {
				t.Fatal("UserID Not Match,")
			}
			if data.(*usecase.FindDTO).MovieID != model.MovieID(trueCase.movieId) {
				t.Fatal("MovieID Not Match,")
			}
			return nil
		})

		indexPlayListInMovieListPageHandler.IndexPlayListInMovieListPage(ginContext)
	}
}
