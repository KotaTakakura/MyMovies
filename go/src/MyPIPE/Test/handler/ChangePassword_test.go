package handler

import (
	mock_repository "MyPIPE/Test/mock/repository"
	mock_usecase "MyPIPE/Test/mock/usecase"
	"MyPIPE/domain/model"
	"MyPIPE/handler"
	"MyPIPE/usecase"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestChangePassword(t *testing.T){

	trueCases := []struct {
		id uint64
		password  string
	}{
		{id: 10, password: "myFirstPassword"},
		{id: 20, password: "mySecondPasword"},
	}

	falseCases := []struct {
		id uint64
		password  string
	}{
		{id: 20, password: ""},
		{id: 30, password: "myFirstPasswordmyFirstPasswordmyFirstPasswordmyFirstPasswordmyFirstPasswordmyFirst" +
			"PasswordmyFirstPasswordmyFirstPasswordmyFirstPasswordmyFirstPasswordmyFirstPasswordmyFirstPasswordmyFirs" +
			"tPasswordmyFirstPasswordmyFirstPasswordmyFirstPassword"},
		{id: 20, password: "ああああああああ"},
	}


	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepository := mock_repository.NewMockUserRepository(ctrl)
	changePasswordUsecase := mock_usecase.NewMockIChangePassword(ctrl)
	changePasswordHandler := handler.NewChangePassword(userRepository,changePasswordUsecase)

	//正常系
	for _,trueCase := range trueCases{
		changePasswordUsecase.EXPECT().ChangePassword(gomock.Any()).DoAndReturn(func (data interface{})error{
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.ChangePasswordDTO{})){
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.ChangePasswordDTO).UserID != model.UserID(trueCase.id){
				t.Fatal("UserID Not Match,")
			}

			err := bcrypt.CompareHashAndPassword(([]byte)(data.(*usecase.ChangePasswordDTO).Password), ([]byte)(trueCase.password))
			if err != nil{
				t.Fatal("Password Not Match,")
			}
			return nil
		})

		// ポストデータ
		bodyReader := strings.NewReader(`{"password": "`+ trueCase.password +`"}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")


		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD",jwt.MapClaims{
			"id":float64(trueCase.id),
		})

		changePasswordHandler.ChangePassword(ginContext)
	}

	//異常系
	for _,falseCase := range falseCases{
		// ポストデータ
		bodyReader := strings.NewReader(`{"password": "`+ falseCase.password +`"}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")


		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD",jwt.MapClaims{
			"id":float64(falseCase.id),
		})

		changePasswordHandler.ChangePassword(ginContext)
	}
}