package handler

import (
	mock_queryService "MyPIPE/Test/mock/queryService"
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
	"strings"
	"testing"
	"time"
)

func TestMovieGetUploadedMovies(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uploadMoviesQueryService := mock_queryService.NewMockUploadedMovies(ctrl)
	uploadMoviesUsecase := mock_usecase.NewMockIUploadedMovies(ctrl)
	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	updateMovieUsecase := mock_usecase.NewMockIUpdateMovie(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	changeThumbnailUsecase := mock_usecase.NewMockIChangeThumbnail(ctrl)
	movieHandler := handler.NewMovie(uploadMoviesQueryService, uploadMoviesUsecase, movieRepository, updateMovieUsecase, thumbnailUploadRepository, changeThumbnailUsecase)

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

		uploadMoviesUsecase.EXPECT().Get(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(model.UserID(0)) {
				t.Fatal("Type Not Match.")
			}
			if data.(model.UserID) != model.UserID(trueCase.userId) {
				t.Fatal("UserID Not Match,")
			}
			return nil
		})

		movieHandler.GetUploadedMovies(ginContext)
	}
}

func TestMovieUpdateMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uploadMoviesQueryService := mock_queryService.NewMockUploadedMovies(ctrl)
	uploadMoviesUsecase := mock_usecase.NewMockIUploadedMovies(ctrl)
	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	updateMovieUsecase := mock_usecase.NewMockIUpdateMovie(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	changeThumbnailUsecase := mock_usecase.NewMockIChangeThumbnail(ctrl)
	movieHandler := handler.NewMovie(uploadMoviesQueryService, uploadMoviesUsecase, movieRepository, updateMovieUsecase, thumbnailUploadRepository, changeThumbnailUsecase)

	trueCases := []struct {
		userId      uint64
		movieId     uint64
		displayName string
		description string
		public      uint
		status      uint
	}{
		{userId: 10, movieId: 100, description: "TestDescription", displayName: "TestDisplayName", public: 1, status: 2},
	}

	for _, trueCase := range trueCases {
		// ポストデータ
		bodyReader := strings.NewReader(`{
			"user_id":` + fmt.Sprint(trueCase.userId) + `,
			"movie_id":` + fmt.Sprint(trueCase.movieId) + `,
			"display_name":"` + fmt.Sprint(trueCase.displayName) + `",
			"description":"` + fmt.Sprint(trueCase.description) + `",
			"public":` + fmt.Sprint(trueCase.public) + `,
			"status":` + fmt.Sprint(trueCase.status) + `
}`)

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

		updateMovieUsecase.EXPECT().Update(gomock.Any()).DoAndReturn(func(data interface{}) (model.Movie, error) {
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.UpdateDTO{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.UpdateDTO).UserID != model.UserID(trueCase.userId) {
				t.Fatal("MovieID Not Match,")
			}
			if data.(*usecase.UpdateDTO).MovieID != model.MovieID(trueCase.movieId) {
				t.Fatal("MovieID Not Match,")
			}
			if data.(*usecase.UpdateDTO).DisplayName != model.MovieDisplayName(trueCase.displayName) {
				t.Fatal("DisplayName Not Match,")
			}
			if data.(*usecase.UpdateDTO).Description != model.MovieDescription(trueCase.description) {
				t.Fatal("Description Not Match,")
			}
			if data.(*usecase.UpdateDTO).Public != model.MoviePublic(trueCase.public) {
				t.Fatal("Public Not Match,")
			}
			if data.(*usecase.UpdateDTO).Status != model.MovieStatus(trueCase.status) {
				t.Fatal("Status Not Match,")
			}

			return model.Movie{
				ID:            model.MovieID(trueCase.movieId),
				StoreName:     "StoreName",
				DisplayName:   model.MovieDisplayName(trueCase.displayName),
				Description:   model.MovieDescription(trueCase.description),
				ThumbnailName: "Thumbnailname",
				UserID:        model.UserID(trueCase.userId),
				Public:        model.MoviePublic(trueCase.public),
				Status:        model.MovieStatus(trueCase.status),
				CreatedAt:     time.Time{},
				UpdatedAt:     time.Time{},
			}, nil
		})

		movieHandler.UpdateMovie(ginContext)

	}

	falseCases := []struct {
		userId      uint64
		movieId     uint64
		displayName string
		description string
		public      uint
		status      uint
	}{
		{userId: 10, movieId: 100, description: "TestDescription", displayName: "TestDisplayName", public: 999, status: 0},
		{userId: 10, movieId: 100, description: "TestDescription", displayName: "TestDisplayName", public: 0, status: 999},
	}

	for _, falseCase := range falseCases {
		// ポストデータ
		bodyReader := strings.NewReader(`{
			"user_id":` + fmt.Sprint(falseCase.userId) + `,
			"movie_id":` + fmt.Sprint(falseCase.movieId) + `,
			"display_name":"` + fmt.Sprint(falseCase.displayName) + `",
			"description":"` + fmt.Sprint(falseCase.description) + `",
			"public":` + fmt.Sprint(falseCase.public) + `,
			"status":` + fmt.Sprint(falseCase.status) + `
}`)

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

		movieHandler.UpdateMovie(ginContext)
	}

	for _, falseCase := range falseCases {
		// ポストデータ
		//movie_idが無効
		bodyReader := strings.NewReader(`{
			"user_id":` + fmt.Sprint(falseCase.userId) + `,
			"movie_id":"INVALIDMOVIEID",
			"display_name":"` + fmt.Sprint(falseCase.displayName) + `",
			"description":"` + fmt.Sprint(falseCase.description) + `",
			"public":` + fmt.Sprint(falseCase.public) + `,
			"status":` + fmt.Sprint(falseCase.status) + `
}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(falseCase.userId),
		})

		movieHandler.UpdateMovie(ginContext)

		if w.Code != http.StatusBadRequest {
			t.Fatal("Error.")
		}
	}
}

func TestMovieUpdateMovie_UsecaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uploadMoviesQueryService := mock_queryService.NewMockUploadedMovies(ctrl)
	uploadMoviesUsecase := mock_usecase.NewMockIUploadedMovies(ctrl)
	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	updateMovieUsecase := mock_usecase.NewMockIUpdateMovie(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	changeThumbnailUsecase := mock_usecase.NewMockIChangeThumbnail(ctrl)
	movieHandler := handler.NewMovie(uploadMoviesQueryService, uploadMoviesUsecase, movieRepository, updateMovieUsecase, thumbnailUploadRepository, changeThumbnailUsecase)

	cases := []struct {
		userId      uint64
		movieId     uint64
		displayName string
		description string
		public      uint
		status      uint
	}{
		{userId: 10, movieId: 100, description: "TestDescription", displayName: "TestDisplayName", public: 1, status: 2},
	}

	for _, Case := range cases {
		bodyReader := strings.NewReader(`{
			"user_id":` + fmt.Sprint(Case.userId) + `,
			"movie_id":` + fmt.Sprint(Case.movieId) + `,
			"display_name":"` + fmt.Sprint(Case.displayName) + `",
			"description":"` + fmt.Sprint(Case.description) + `",
			"public":` + fmt.Sprint(Case.public) + `,
			"status":` + fmt.Sprint(Case.status) + `
}`)

		// リクエスト生成
		req := httptest.NewRequest("POST", "/", bodyReader)

		// Content-Type 設定
		req.Header.Set("Content-Type", "application/json")

		// Contextセット
		w := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(w)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(Case.userId),
		})

		updateMovieUsecase.EXPECT().Update(gomock.Any()).Return(nil, errors.New("Usecase Error"))

		movieHandler.UpdateMovie(ginContext)

		if w.Code != http.StatusInternalServerError {
			t.Fatal("Error.")
		}
	}
}

func TestMovie_ChangeThumbnail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uploadMoviesQueryService := mock_queryService.NewMockUploadedMovies(ctrl)
	uploadMoviesUsecase := mock_usecase.NewMockIUploadedMovies(ctrl)
	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	updateMovieUsecase := mock_usecase.NewMockIUpdateMovie(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	changeThumbnailUsecase := mock_usecase.NewMockIChangeThumbnail(ctrl)
	movieHandler := handler.NewMovie(uploadMoviesQueryService, uploadMoviesUsecase, movieRepository, updateMovieUsecase, thumbnailUploadRepository, changeThumbnailUsecase)

	// ポストデータ
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	f, e1 := os.Open("./../TestThumbnail.jpg")
	if e1 != nil {
		fmt.Println(e1)
		return
	}
	defer f.Close()
	fw, e2 := w.CreateFormFile("uploadThumbnail", "TestThumbnail.jpg")
	if e2 != nil {
		fmt.Println(e2)
		return
	}
	_, e3 := io.Copy(fw, f)
	if e3 != nil {
		fmt.Println(e3)
		return
	}

	w.WriteField("movie_id", "10")
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

	changeThumbnailUsecase.EXPECT().ChangeThumbnail(gomock.Any()).DoAndReturn(func(data interface{}) error {
		if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.ChangeThumbnailDTO{})) {
			t.Fatal("Type Not Match.")
		}
		if data.(*usecase.ChangeThumbnailDTO).MovieID != model.MovieID(10) {
			t.Fatal("Movie Not Match,")
		}
		file, fileHeader, _ := ginContext.Request.FormFile("uploadThumbnail")
		movieThumbnail, _ := model.NewMovieThumbnail(file, *fileHeader)
		if reflect.DeepEqual(data.(*usecase.ChangeThumbnailDTO).Thumbnail, *movieThumbnail) {
			t.Fatal("Thumbnail Not Match,")
		}
		return nil
	})

	movieHandler.ChangeThumbnail(ginContext)

	if recorder.Code != http.StatusOK {
		t.Fatal("Usecase Error.")
	}
}

func TestMovie_ChangeThumbnail_MovieID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uploadMoviesQueryService := mock_queryService.NewMockUploadedMovies(ctrl)
	uploadMoviesUsecase := mock_usecase.NewMockIUploadedMovies(ctrl)
	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	updateMovieUsecase := mock_usecase.NewMockIUpdateMovie(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	changeThumbnailUsecase := mock_usecase.NewMockIChangeThumbnail(ctrl)
	movieHandler := handler.NewMovie(uploadMoviesQueryService, uploadMoviesUsecase, movieRepository, updateMovieUsecase, thumbnailUploadRepository, changeThumbnailUsecase)

	// ポストデータ
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	f, e1 := os.Open("./../TestThumbnail.jpg")
	if e1 != nil {
		fmt.Println(e1)
		return
	}
	defer f.Close()
	fw, e2 := w.CreateFormFile("uploadThumbnail", "TestThumbnail.jpg")
	if e2 != nil {
		fmt.Println(e2)
		return
	}
	_, e3 := io.Copy(fw, f)
	if e3 != nil {
		fmt.Println(e3)
		return
	}

	w.WriteField("movie_id", "INVALIDMOVIEID")
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

	movieHandler.ChangeThumbnail(ginContext)

	if recorder.Code != http.StatusBadRequest {
		t.Fatal("Usecase Error.")
	}
}

func TestMovie_ChangeThumbnail_ThumbnailInvalid_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uploadMoviesQueryService := mock_queryService.NewMockUploadedMovies(ctrl)
	uploadMoviesUsecase := mock_usecase.NewMockIUploadedMovies(ctrl)
	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	updateMovieUsecase := mock_usecase.NewMockIUpdateMovie(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	changeThumbnailUsecase := mock_usecase.NewMockIChangeThumbnail(ctrl)
	movieHandler := handler.NewMovie(uploadMoviesQueryService, uploadMoviesUsecase, movieRepository, updateMovieUsecase, thumbnailUploadRepository, changeThumbnailUsecase)

	// ポストデータ
	var tooLargeFileBuffer bytes.Buffer
	tooLargeFile := multipart.NewWriter(&tooLargeFileBuffer)
	f1, e11 := os.Open("./../TestThumbnailTooLarge.png")
	if e11 != nil {
		return
	}
	defer f1.Close()
	fw1, e12 := tooLargeFile.CreateFormFile("uploadThumbnail", "TestThumbnailTooLarge.png")
	if e12 != nil {
		return
	}
	_, e13 := io.Copy(fw1, f1)
	if e13 != nil {
		return
	}
	tooLargeFile.WriteField("movie_id", "10")
	tooLargeFile.Close()

	// ポストデータ
	var notImageFileBuffer bytes.Buffer
	notImageFile := multipart.NewWriter(&notImageFileBuffer)
	f2, e21 := os.Open("./../TestThumbnail.jpg")
	if e21 != nil {
		return
	}
	defer f2.Close()
	fw2, e22 := notImageFile.CreateFormFile("uploadThumbnail", "TestThumbnailNotImage.mp4")
	if e22 != nil {
		return
	}
	_, e23 := io.Copy(fw2, f2)
	if e23 != nil {
		return
	}
	notImageFile.WriteField("movie_id", "10")
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

		movieHandler.ChangeThumbnail(ginContext)

		if recorder.Code != http.StatusBadRequest {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestMovie_ChangeThumbnail_Usecase_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uploadMoviesQueryService := mock_queryService.NewMockUploadedMovies(ctrl)
	uploadMoviesUsecase := mock_usecase.NewMockIUploadedMovies(ctrl)
	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	updateMovieUsecase := mock_usecase.NewMockIUpdateMovie(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	changeThumbnailUsecase := mock_usecase.NewMockIChangeThumbnail(ctrl)
	movieHandler := handler.NewMovie(uploadMoviesQueryService, uploadMoviesUsecase, movieRepository, updateMovieUsecase, thumbnailUploadRepository, changeThumbnailUsecase)

	// ポストデータ
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	f, e1 := os.Open("./../TestThumbnail.jpg")
	if e1 != nil {
		fmt.Println(e1)
		return
	}
	defer f.Close()
	fw, e2 := w.CreateFormFile("uploadThumbnail", "TestThumbnail.jpg")
	if e2 != nil {
		fmt.Println(e2)
		return
	}
	_, e3 := io.Copy(fw, f)
	if e3 != nil {
		fmt.Println(e3)
		return
	}

	w.WriteField("movie_id", "10")
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

	changeThumbnailUsecase.EXPECT().ChangeThumbnail(gomock.Any()).Return(errors.New("ERROR"))

	movieHandler.ChangeThumbnail(ginContext)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatal("Usecase Error.")
	}
}
