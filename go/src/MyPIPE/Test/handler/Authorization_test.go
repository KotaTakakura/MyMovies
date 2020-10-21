package handler

import (
	mock_repository "MyPIPE/Test/mock/repository"
	mock_usecase "MyPIPE/Test/mock/usecase"
	"MyPIPE/domain/model"
	"MyPIPE/handler"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestTemporaryRegistration(t *testing.T){

	trueCases := []struct {
		email string
	}{
		{email: "trueUsecase1@mail.com"},
		{email: "trueUsecase2@amaammama.jp"},
	}

	falseCases := []struct {
		email string
	}{
		{email: ""},
		{email: "trueUsecaseamaammama.jp"},
		{email: "trueUsecaseamaammama.jp"},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepository := mock_repository.NewMockUserRepository(ctrl)
	userTemporaryRegistrationUsecase := mock_usecase.NewMockIUserTemporaryRegistration(ctrl)
	userRegisterUsecase := mock_usecase.NewMockIUserRegister(ctrl)
	authorizationHandler := handler.NewAuthorization(userRepository,userTemporaryRegistrationUsecase,userRegisterUsecase)

	for _,value := range trueCases{
		// ポストデータ
		bodyReader := strings.NewReader(`{"user_email": "`+ value.email + `"}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")


		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req

		userTemporaryRegistrationUsecase.EXPECT().TemporaryRegister(gomock.Any()).DoAndReturn(func(data interface{})error{
			if reflect.TypeOf(data) != reflect.TypeOf(&(model.User{})){
				t.Fatal("Type Not Match.")
			}
			if data.(*model.User).Email != model.UserEmail(value.email){
				t.Fatal("Email Not Match,")
			}
			return  nil
		})

		authorizationHandler.TemporaryRegisterUser(ginContext)
	}

	for _,value := range falseCases{
		// ポストデータ
		bodyReader := strings.NewReader(`{"user_email": "`+ value.email + `"}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req

		authorizationHandler.TemporaryRegisterUser(ginContext)
	}
}

func TestRegisterUser(t *testing.T){

	trueCases := []struct {
		token string
		userName string
		userPassword string
		userBirthday string
	}{
		{
			token: "191983-82723723-2329323-232323233",
			userName: "田中太郎",
			userPassword: "iekapcr92248",
			userBirthday:	"2000-01-11",
		},
	}

	falseCases := []struct {
		token string
		userName string
		userPassword string
		userBirthday string
	}{
		{
			token: "",//false
			userName: "田中太郎",
			userPassword: "iekapcr92248",
			userBirthday:	"2000-01-11",
		},
		{
			token: "191983-82723723-2329323-232323233",
			userName: "",//false
			userPassword: "iekapcr92248",
			userBirthday:	"2000-01-11",
		},
		{
			token: "191983-82723723-2329323-232323233",
			userName: "田中太郎",
			userPassword: "",//false
			userBirthday:	"2000-01-11",
		},
		{
			token: "191983-82723723-2329323-232323233",
			userName: "田中太郎",
			userPassword: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
							"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
							"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",//false
			userBirthday:	"2000-01-11",
		},
		{
			token: "191983-82723723-2329323-232323233",
			userName: "田中太郎",
			userPassword: "iekapcr92248",
			userBirthday:	"",//false
		},
		{
			token: "191983-82723723-2329323-232323233",
			userName: "田中太郎",
			userPassword: "iekapcr92248",
			userBirthday:	"20000101",//false
		},
		{
			token: "191983-82723723-2329323-232323233",
			userName: "田中太郎",
			userPassword: "iekapcr92248",
			userBirthday:	"2000/01/01",//false
		},
		{
			token: "191983-82723723-2329323-232323233",
			userName: "田中太郎",
			userPassword: "iekapcr92248",
			userBirthday:	"2000/13/01",//false
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepository := mock_repository.NewMockUserRepository(ctrl)
	userTemporaryRegistrationUsecase := mock_usecase.NewMockIUserTemporaryRegistration(ctrl)
	userRegisterUsecase := mock_usecase.NewMockIUserRegister(ctrl)
	authorizationHandler := handler.NewAuthorization(userRepository,userTemporaryRegistrationUsecase,userRegisterUsecase)

	for _,value := range trueCases{
		// ポストデータ
		bodyReader := strings.NewReader(`{"user_name": "`+ value.userName + `","user_password": "`+ value.userPassword + `","user_birthday": "`+ value.userBirthday + `"}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/?token=" + value.token, bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")


		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req

		userRegisterUsecase.EXPECT().RegisterUser(gomock.Any()).DoAndReturn(func(data interface{})error{
			if reflect.TypeOf(data) != reflect.TypeOf(&(model.User{})){
				t.Fatal("Type Not Match.")
			}

			if data.(*model.User).Name != model.UserName(value.userName){
				t.Fatal("Name Not Match,")
			}

			err := bcrypt.CompareHashAndPassword(([]byte)(data.(*model.User).Password), ([]byte)(value.userPassword))
			if err != nil{
				t.Fatal("Password Not Match,")
			}

			if data.(*model.User).Token != model.UserToken(value.token){
				t.Fatal("Token Not Match,")
			}
			return  nil
		})

		authorizationHandler.RegisterUser(ginContext)
	}

	for _,value := range falseCases{
		// ポストデータ
		bodyReader := strings.NewReader(`{"user_name": "`+ value.userName + `","user_password": "`+ value.userPassword + `","user_birthday": "`+ value.userBirthday + `"}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/?token=" + value.token, bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")


		// Contextセット
		ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		ginContext.Request = req

		authorizationHandler.RegisterUser(ginContext)
	}
}
