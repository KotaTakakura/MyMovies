package handler

import (
	mock_repository "MyPIPE/Test/mock/repository"
	mock_usecase "MyPIPE/Test/mock/usecase"
	"MyPIPE/domain/model"
	"MyPIPE/handler"
	"MyPIPE/usecase"
	"bytes"
	"errors"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestChangeUserProfileImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_repository.NewMockUserRepository(ctrl)
	userProfileImageRepository := mock_repository.NewMockUserProfileImageRepository(ctrl)
	changeUserProfileImageUsecase := mock_usecase.NewMockIChangeUserProfilieImage(ctrl)
	changeUserProfileImageHandler := handler.NewChangeUserProfileImage(userRepository, userProfileImageRepository, changeUserProfileImageUsecase)

	// ポストデータ
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	f, e1 := os.Open("./../TestProfileImage.jpg")
	if e1 != nil {
		fmt.Println(e1)
		return
	}
	defer f.Close()
	fw, e2 := w.CreateFormFile("profileImage", "TestProfileImage.jpg")
	if e2 != nil {
		fmt.Println(e2)
		return
	}
	_, e3 := io.Copy(fw, f)
	if e3 != nil {
		fmt.Println(e3)
		return
	}
	w.Close()

	// リクエスト生成
	req := httptest.NewRequest("POST", "/", &buf)

	// Content-Type 設定
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Contextセット
	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	ginContext.Request = req
	ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
		"id": float64(10),
	})

	changeUserProfileImageUsecase.EXPECT().ChangeUserProfileImage(gomock.Any()).DoAndReturn(func(data interface{}) error {
		if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.ChangeUserProfileImageDTO{})) {
			t.Fatal("Type Not Match.")
		}
		if data.(*usecase.ChangeUserProfileImageDTO).UserID != model.UserID(10) {
			t.Fatal("UserID Not Match,")
		}
		file, fileHeader, _ := ginContext.Request.FormFile("profileImage")
		requestProfileImage, _ := model.NewUserProfileImage(*fileHeader, file)
		if reflect.DeepEqual(data.(*usecase.ChangeUserProfileImageDTO).ProfileImage, requestProfileImage) {
			t.Fatal("ProfileImage Not Match,")
		}
		return nil
	})

	changeUserProfileImageHandler.ChangeUserProfileImage(ginContext)
}

func TestChangeUserProfileImage_ProfileImage_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_repository.NewMockUserRepository(ctrl)
	userProfileImageRepository := mock_repository.NewMockUserProfileImageRepository(ctrl)
	changeUserProfileImageUsecase := mock_usecase.NewMockIChangeUserProfilieImage(ctrl)
	changeUserProfileImageHandler := handler.NewChangeUserProfileImage(userRepository, userProfileImageRepository, changeUserProfileImageUsecase)

	// ポストデータ
	var tooLargeFileBuffer bytes.Buffer
	tooLargeFile := multipart.NewWriter(&tooLargeFileBuffer)
	f1, e11 := os.Open("./../TestProfileImageTooLarge.png")
	if e11 != nil {
		return
	}
	defer f1.Close()
	fw1, e12 := tooLargeFile.CreateFormFile("profileImage", "TestProfileImageTooLarge.png")
	if e12 != nil {
		return
	}
	_, e13 := io.Copy(fw1, f1)
	if e13 != nil {
		return
	}
	tooLargeFile.Close()

	// ポストデータ
	var notImageFileBuffer bytes.Buffer
	notImageFile := multipart.NewWriter(&notImageFileBuffer)
	f2, e21 := os.Open("./../TestProfileImageTooLarge.png")
	if e21 != nil {
		return
	}
	defer f2.Close()
	fw2, e22 := tooLargeFile.CreateFormFile("profileImage", "TestProfileImageTooLarge.mp4")
	if e22 != nil {
		return
	}
	_, e23 := io.Copy(fw2, f2)
	if e23 != nil {
		return
	}
	notImageFile.Close()

	cases := []struct {
		fileBuffer bytes.Buffer
		fileWriter *multipart.Writer
	}{
		{
			fileBuffer: tooLargeFileBuffer,
			fileWriter: tooLargeFile,
		},
		{
			fileBuffer: notImageFileBuffer,
			fileWriter: notImageFile,
		},
	}

	for _, Case := range cases {
		// リクエスト生成
		req := httptest.NewRequest("POST", "/", &Case.fileBuffer)

		// Content-Type 設定
		req.Header.Set("Content-Type", Case.fileWriter.FormDataContentType())

		// Contextセット
		recorder := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(recorder)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(10),
		})

		changeUserProfileImageHandler.ChangeUserProfileImage(ginContext)

		if recorder.Code != http.StatusBadRequest {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestChangeUserProfileImage_UsecaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_repository.NewMockUserRepository(ctrl)
	userProfileImageRepository := mock_repository.NewMockUserProfileImageRepository(ctrl)
	changeUserProfileImageUsecase := mock_usecase.NewMockIChangeUserProfilieImage(ctrl)
	changeUserProfileImageHandler := handler.NewChangeUserProfileImage(userRepository, userProfileImageRepository, changeUserProfileImageUsecase)

	// ポストデータ
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	f, e1 := os.Open("./../TestProfileImage.jpg")
	if e1 != nil {
		fmt.Println(e1)
		return
	}
	defer f.Close()
	fw, e2 := w.CreateFormFile("profileImage", "TestProfileImage.jpg")
	if e2 != nil {
		fmt.Println(e2)
		return
	}
	_, e3 := io.Copy(fw, f)
	if e3 != nil {
		fmt.Println(e3)
		return
	}
	w.Close()

	// リクエスト生成
	req := httptest.NewRequest("POST", "/", &buf)

	// Content-Type 設定
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Contextセット
	recorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(recorder)
	ginContext.Request = req
	ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
		"id": float64(10),
	})

	changeUserProfileImageUsecase.EXPECT().ChangeUserProfileImage(gomock.Any()).Return(errors.New("Usecase Error."))

	changeUserProfileImageHandler.ChangeUserProfileImage(ginContext)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatal("Usecase error")
	}
}
