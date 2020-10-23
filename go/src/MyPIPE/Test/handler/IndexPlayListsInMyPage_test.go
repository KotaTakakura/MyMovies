package handler

import (
	mock_queryService "MyPIPE/Test/mock/queryService"
	mock_usecase "MyPIPE/Test/mock/usecase"
	"MyPIPE/domain/model"
	"MyPIPE/handler"
	"MyPIPE/usecase"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestIndexPlayListsInMyPage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	indexPlayListsInMyPageQueryService := mock_queryService.NewMockIndexPlayListsInMyPageQueryService(ctrl)
	indexPlayListsInMyPageUsecase := mock_usecase.NewMockIIndexPlayListsInMyPage(ctrl)
	indexPlayListsInMyPageHandler := handler.NewIndexPlayListsInMyPage(indexPlayListsInMyPageQueryService, indexPlayListsInMyPageUsecase)

	trueCases := []struct {
		userId uint64
	}{
		{userId: 10},
	}

	for _, trueCase := range trueCases {
		// ポストデータ
		bodyReader := strings.NewReader("")

		// リクエスト生成
		req := httptest.NewRequest("GET", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(trueCase.userId),
		})

		indexPlayListsInMyPageUsecase.EXPECT().All(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.IndexPlayListsInMyPageDTO{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.IndexPlayListsInMyPageDTO).UserID != model.UserID(trueCase.userId) {
				t.Fatal("UserID Not Match,")
			}
			return nil
		})

		indexPlayListsInMyPageHandler.IndexPlayListsInMyPage(ginContext)
	}
}
