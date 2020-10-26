package handler

import (
	mock_queryService "MyPIPE/Test/mock/queryService"
	mock_usecase "MyPIPE/Test/mock/usecase"
	//"MyPIPE/domain/model"
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

func TestIndexMovie(t *testing.T) {

	trueCases := []struct {
		keyWord string
		order   string
		page    int
		url     string
	}{
		{keyWord: "game", order: "asc", page: 10, url: "/?keyWord=game&order=asc&page=10"},
		{keyWord: "movie", order: "desc", page: -1, url: "/?keyWord=movie&order=desc&page=-1"},    //page = 1に変換される
		{keyWord: "movie", order: "invalid", page: 10, url: "/?keyWord=movie&order=desc&page=-1"}, //order = "asc"に変換される
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	indexMovieUsecase := mock_usecase.NewMockIIndexMovie(ctrl)
	indexMovieQueryService := mock_queryService.NewMockIndexMovieQueryService(ctrl)
	indexMovieHandler := handler.NewIndexMovie(indexMovieQueryService, indexMovieUsecase)

	for _, trueCase := range trueCases {
		// ポストデータ
		bodyReader := strings.NewReader("")

		// リクエスト生成
		req := httptest.NewRequest("GET", trueCase.url, bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req

		indexMovieUsecase.EXPECT().Search(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.IndexMovieSearchDTO{})) {
				fmt.Println(reflect.TypeOf(data))
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.IndexMovieSearchDTO).KeyWord != trueCase.keyWord {
				t.Fatal("keyWord Not Match,")
			}
			if data.(*usecase.IndexMovieSearchDTO).Page != queryService.IndexMovieQueryServicePage(trueCase.page) && data.(*usecase.IndexMovieSearchDTO).Page != 1 {
				t.Fatal("Page Not Match,")
			}
			if !(data.(*usecase.IndexMovieSearchDTO).Order == "asc" || data.(*usecase.IndexMovieSearchDTO).Order == "desc") {
				t.Fatal("Order Not Match,")
			}
			return nil
		})

		indexMovieHandler.IndexMovie(ginContext)

	}
}
