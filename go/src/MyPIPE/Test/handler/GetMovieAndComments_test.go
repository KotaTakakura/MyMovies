package handler

import (
	mock_queryService "MyPIPE/Test/mock/queryService"
	mock_usecase "MyPIPE/Test/mock/usecase"
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
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

func TestGetMovieAndComments(t *testing.T) {
	trueCases := []struct {
		movieId uint64
	}{
		{movieId: 10},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	getCommentQueryService := mock_queryService.NewMockCommentQueryService(ctrl)
	getCommentUsecase := mock_usecase.NewMockIGetMovieAndComments(ctrl)
	getMovieAndCommentsHandler := handler.NewGetMovieAndComments(getCommentQueryService, getCommentUsecase)

	for _, trueCase := range trueCases {
		// ポストデータ
		bodyReader := strings.NewReader("")

		// リクエスト生成
		req := httptest.NewRequest("GET", "/?movie_id="+fmt.Sprint(trueCase.movieId), bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req

		result := &(queryService.FindByMovieIdDTO{})

		getCommentUsecase.EXPECT().Get(gomock.Any()).DoAndReturn(func(data interface{}) *queryService.FindByMovieIdDTO {
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.MovieAndGetCommentsDTO{})) {
				fmt.Println(reflect.TypeOf(data))
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.MovieAndGetCommentsDTO).MovieID != model.MovieID(trueCase.movieId) {
				t.Fatal("MovieID Not Match,")
			}
			return result
		})
		getMovieAndCommentsHandler.GetMovieAndComments(ginContext)
	}
}
