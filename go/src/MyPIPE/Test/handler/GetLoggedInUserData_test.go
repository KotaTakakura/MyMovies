package handler

import (
	mock_queryService "MyPIPE/Test/mock/queryService"
	mock_usecase "MyPIPE/Test/mock/usecase"
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
	"MyPIPE/handler"
	"MyPIPE/usecase"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"reflect"
	//"fmt"
	//"strings"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestGetLoggedInUserData(t *testing.T){

	trueCases := []struct {
		userId uint64
	}{
		{userId: 10},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	getLoggedInUserDataQueryService := mock_queryService.NewMockGetLoggedInUserDataQueryService(ctrl)
	getLoggedInUserDataUsecase := mock_usecase.NewMockIGetLoggedInUserData(ctrl)
	getLoggedInUserDataHandler := handler.NewGetLoggedInUserData(getLoggedInUserDataQueryService,getLoggedInUserDataUsecase)

	for _,trueCase := range trueCases{
		// リクエスト生成
		req := httptest.NewRequest("GET", "/", nil)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Contextセット
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD",jwt.MapClaims{
			"id":float64(trueCase.userId),
		})

		timeTest := time.Now()
		getLoggedInUserDataDTO := &(queryService.GetLoggedInUserDataDTO{
			ID:               trueCase.userId,
			Name:             "TestName",
			Email:            "TestEmail",
			Birthday:         timeTest,
			ProfileImageName: "TestProfileImage.jpg",
			CreatedAt:        timeTest,
			AvatarName:       "TesetAvatarName",
		})

		getLoggedInUserDataUsecase.EXPECT().Find(gomock.Any()).DoAndReturn(func(data interface{})*queryService.GetLoggedInUserDataDTO{
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.GetLoggedInUserDataDTO{})){
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.GetLoggedInUserDataDTO).UserID != model.UserID(trueCase.userId){
				t.Fatal("UserID Not Match,")
			}
			return getLoggedInUserDataDTO
		})

		assert.Equal(t,http.StatusOK,ginContext.Writer.Status())

		getLoggedInUserDataHandler.GetLoggedInUserData(ginContext)
	}
}
