package handler

import (
	mock_repository "MyPIPE/Test/mock/repository"
	mock_usecase "MyPIPE/Test/mock/usecase"
	"MyPIPE/domain/model"
	"MyPIPE/handler"
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
	"testing"
)

func TestUploadMovieFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	movieUploadRepository := mock_repository.NewMockFileUpload(ctrl)
	postMovieUsecase := mock_usecase.NewMockIPostMovie(ctrl)
	uploadMovieFileHandler := handler.NewUploadMovieFile(movieRepository, thumbnailUploadRepository, movieUploadRepository, postMovieUsecase)

	// サムネイル
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	thumbnailFile, thumbnailFileErr := os.Open("./../TestThumbnail.jpg")
	if thumbnailFileErr != nil {
		fmt.Println(1)
		fmt.Println(thumbnailFileErr)
		return
	}

	movieFile, movieFileErr := os.Open("./../TestMovie.mp4")
	if movieFileErr != nil {
		fmt.Println(2)
		fmt.Println(movieFileErr)
		return
	}

	defer thumbnailFile.Close()
	defer movieFile.Close()

	thumbnailFileForm, thumbnailFileFormErr := w.CreateFormFile("uploadThumbnail", "TestThumbnail.jpg")
	if thumbnailFileFormErr != nil {
		fmt.Println(3)
		fmt.Println(thumbnailFileFormErr)
		return
	}

	_, thumbnailCopyErr := io.Copy(thumbnailFileForm, thumbnailFile)
	if thumbnailCopyErr != nil {
		fmt.Println(5)
		fmt.Println(thumbnailCopyErr)
		return
	}

	movieFileForm, movieFileFormErr := w.CreateFormFile("uploadMovie", "TestMovie.mp4")
	if movieFileFormErr != nil {
		fmt.Println(4)
		fmt.Println(movieFileFormErr)
		return
	}
	_, movieCopyErr := io.Copy(movieFileForm, movieFile)
	if movieCopyErr != nil {
		fmt.Println(6)
		fmt.Println(movieCopyErr)
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

	postMovieUsecase.EXPECT().PostMovie(gomock.Any()).Return(&model.Movie{
		ID:            20,
		StoreName:     "PostedmovieStoreName",
		DisplayName:   "PostedmovieDisplayName",
		Description:   "PostedmovieDescription",
		ThumbnailName: "PostedmovieThumbnailName",
		UserID:        10,
		Public:        0,
		Status:        0,
	}, nil)

	uploadMovieFileHandler.UploadMovieFile(ginContext)

	if recorder.Code != http.StatusOK {
		t.Fatal("Usecase Error")
	}
}

func TestUploadMovieFile_Thumbnail_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	movieUploadRepository := mock_repository.NewMockFileUpload(ctrl)
	postMovieUsecase := mock_usecase.NewMockIPostMovie(ctrl)
	uploadMovieFileHandler := handler.NewUploadMovieFile(movieRepository, thumbnailUploadRepository, movieUploadRepository, postMovieUsecase)

	// サムネイル
	var noThumbnailbuf bytes.Buffer
	noThumbnailWriter := multipart.NewWriter(&noThumbnailbuf)

	movieFileNoThumbnanil, movieFileNoThumbnanilErr := os.Open("./../TestMovie.mp4")
	if movieFileNoThumbnanilErr != nil {
		fmt.Println(2)
		fmt.Println(movieFileNoThumbnanilErr)
		return
	}

	defer movieFileNoThumbnanil.Close()

	movieNoThumbnanilFileForm, movieNoThumbnanilFileFormErr := noThumbnailWriter.CreateFormFile("uploadMovie", "TestMovie.mp4")
	if movieNoThumbnanilFileFormErr != nil {
		fmt.Println(4)
		fmt.Println(movieNoThumbnanilFileFormErr)
		return
	}
	_, movieNoThumbnanilCopyErr := io.Copy(movieNoThumbnanilFileForm, movieFileNoThumbnanil)
	if movieNoThumbnanilCopyErr != nil {
		fmt.Println(6)
		fmt.Println(movieNoThumbnanilCopyErr)
		return
	}

	noThumbnailWriter.Close()

	// サムネイル
	var tooLargetThumbnailbuf bytes.Buffer
	tooLargetThumbnailWriter := multipart.NewWriter(&tooLargetThumbnailbuf)
	tooLargetThumbnailFile, tooLargetThumbnailFileErr := os.Open("./../TestThumbnailTooLarge.png")
	if tooLargetThumbnailFileErr != nil {
		fmt.Println(1)
		fmt.Println(tooLargetThumbnailFileErr)
		return
	}

	movieFile, movieFileErr := os.Open("./../TestMovie.mp4")
	if movieFileErr != nil {
		fmt.Println(2)
		fmt.Println(movieFileErr)
		return
	}

	defer tooLargetThumbnailFile.Close()
	defer movieFile.Close()

	thumbnailTooLargeFileForm, thumbnailTooLargeFileFormErr := tooLargetThumbnailWriter.CreateFormFile("uploadThumbnail", "TestThumbnailTooLarge.png")
	if thumbnailTooLargeFileFormErr != nil {
		fmt.Println(3)
		fmt.Println(thumbnailTooLargeFileFormErr)
		return
	}

	_, thumbnailCopyErr := io.Copy(thumbnailTooLargeFileForm, tooLargetThumbnailFile)
	if thumbnailCopyErr != nil {
		fmt.Println(5)
		fmt.Println(thumbnailCopyErr)
		return
	}

	movieFileToolargeThumbnailForm, movieFileToolargeThumbnailFormErr := tooLargetThumbnailWriter.CreateFormFile("uploadMovie", "TestMovie.mp4")
	if movieFileToolargeThumbnailFormErr != nil {
		fmt.Println(4)
		fmt.Println(movieFileToolargeThumbnailFormErr)
		return
	}
	_, movieCopyErr := io.Copy(movieFileToolargeThumbnailForm, movieFile)
	if movieCopyErr != nil {
		fmt.Println(6)
		fmt.Println(movieCopyErr)
		return
	}

	tooLargetThumbnailWriter.Close()

	// サムネイル
	var invalidThumbnailbuf bytes.Buffer
	invalidThumbnailWriter := multipart.NewWriter(&invalidThumbnailbuf)
	invalidThumbnailFile, invalidThumbnailFileErr := os.Open("./../TestThumbnailTooLarge.png")
	if invalidThumbnailFileErr != nil {
		fmt.Println(1)
		fmt.Println(invalidThumbnailFileErr)
		return
	}

	movieInvalidThumbnailFile, movieInvalidThumbnailFileErr := os.Open("./../TestMovie.mp4")
	if movieInvalidThumbnailFileErr != nil {
		fmt.Println(2)
		fmt.Println(movieInvalidThumbnailFileErr)
		return
	}

	defer tooLargetThumbnailFile.Close()
	defer movieFile.Close()

	invalidThumbnailFileForm, invalidThumbnailFileFormErr := invalidThumbnailWriter.CreateFormFile("uploadThumbnail", "TestThumbnailTooLarge.mp4")
	if invalidThumbnailFileFormErr != nil {
		fmt.Println(3)
		fmt.Println(invalidThumbnailFileFormErr)
		return
	}

	_, invalidThumbnailCopyErr := io.Copy(invalidThumbnailFileForm, invalidThumbnailFile)
	if invalidThumbnailCopyErr != nil {
		fmt.Println(5)
		fmt.Println(invalidThumbnailCopyErr)
		return
	}

	movieFileInvalidThumbnailForm, movieInvalidThumbnailFormErr := invalidThumbnailWriter.CreateFormFile("uploadMovie", "TestMovie.mp4")
	if movieInvalidThumbnailFormErr != nil {
		fmt.Println(4)
		fmt.Println(movieInvalidThumbnailFormErr)
		return
	}
	_, movieInvalidThumbnailCopyErr := io.Copy(movieFileInvalidThumbnailForm, movieInvalidThumbnailFile)
	if movieInvalidThumbnailCopyErr != nil {
		fmt.Println(6)
		fmt.Println(movieInvalidThumbnailCopyErr)
		return
	}

	invalidThumbnailWriter.Close()

	cases := []struct {
		buffer bytes.Buffer
		writer *multipart.Writer
	}{
		{
			buffer: noThumbnailbuf,
			writer: noThumbnailWriter,
		},
		{
			buffer: tooLargetThumbnailbuf,
			writer: tooLargetThumbnailWriter,
		},
		{
			buffer: invalidThumbnailbuf,
			writer: invalidThumbnailWriter,
		},
	}

	for _, Case := range cases {
		// リクエスト生成
		req := httptest.NewRequest("POST", "/", &Case.buffer)

		// Content-Type 設定
		req.Header.Set("Content-Type", Case.writer.FormDataContentType())

		// Contextセット
		recorder := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(recorder)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(10),
		})

		uploadMovieFileHandler.UploadMovieFile(ginContext)

		if recorder.Code != http.StatusBadRequest {
			t.Fatal("Usecase Error")
		}
	}
}

func TestUploadMovieFile_Movie_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	movieUploadRepository := mock_repository.NewMockFileUpload(ctrl)
	postMovieUsecase := mock_usecase.NewMockIPostMovie(ctrl)
	uploadMovieFileHandler := handler.NewUploadMovieFile(movieRepository, thumbnailUploadRepository, movieUploadRepository, postMovieUsecase)

	var noMoviebuf bytes.Buffer
	noMovieWriter := multipart.NewWriter(&noMoviebuf)
	thumbnailNoMovieFile, thumbnailNoMovieFileErr := os.Open("./../TestThumbnail.jpg")
	if thumbnailNoMovieFileErr != nil {
		fmt.Println(1)
		fmt.Println(thumbnailNoMovieFileErr)
		return
	}

	defer thumbnailNoMovieFile.Close()

	thumbnailFileForm, thumbnailFileFormErr := noMovieWriter.CreateFormFile("uploadThumbnail", "TestThumbnail.jpg")
	if thumbnailFileFormErr != nil {
		fmt.Println(3)
		fmt.Println(thumbnailFileFormErr)
		return
	}

	_, thumbnailCopyErr := io.Copy(thumbnailFileForm, thumbnailNoMovieFile)
	if thumbnailCopyErr != nil {
		fmt.Println(5)
		fmt.Println(thumbnailCopyErr)
		return
	}

	noMovieWriter.Close()

	var invalidMoviebuf bytes.Buffer
	invalidMovieWriter := multipart.NewWriter(&invalidMoviebuf)
	thumbnailInvalidMovieFile, thumbnailInvalidMovieFileErr := os.Open("./../TestThumbnail.jpg")
	if thumbnailInvalidMovieFileErr != nil {
		fmt.Println(1)
		fmt.Println(thumbnailInvalidMovieFileErr)
		return
	}

	invalidMovieFile, invalidMovieFileErr := os.Open("./../TestMovie.mp4")
	if invalidMovieFileErr != nil {
		fmt.Println(2)
		fmt.Println(invalidMovieFileErr)
		return
	}

	defer thumbnailInvalidMovieFile.Close()
	defer invalidMovieFile.Close()

	thumbnailInvalidMovieFileForm, thumbnailInvalidMovieFileFormErr := invalidMovieWriter.CreateFormFile("uploadThumbnail", "TestThumbnail.jpg")
	if thumbnailInvalidMovieFileFormErr != nil {
		fmt.Println(3)
		fmt.Println(thumbnailInvalidMovieFileFormErr)
		return
	}

	_, thumbnailInvalidMovieCopyErr := io.Copy(thumbnailInvalidMovieFileForm, thumbnailInvalidMovieFile)
	if thumbnailInvalidMovieCopyErr != nil {
		fmt.Println(5)
		fmt.Println(thumbnailInvalidMovieCopyErr)
		return
	}

	invalidMovieFileForm, invalidMFileFormErr := invalidMovieWriter.CreateFormFile("uploadMovie", "TestMovie.jpg")
	if invalidMFileFormErr != nil {
		fmt.Println(4)
		fmt.Println(invalidMFileFormErr)
		return
	}
	_, invalidMovieCopyErr := io.Copy(invalidMovieFileForm, invalidMovieFile)
	if invalidMovieCopyErr != nil {
		fmt.Println(6)
		fmt.Println(invalidMovieCopyErr)
		return
	}
	invalidMovieWriter.Close()

	cases := []struct {
		buffer bytes.Buffer
		writer *multipart.Writer
	}{
		{
			buffer: noMoviebuf,
			writer: noMovieWriter,
		},
		{
			buffer: invalidMoviebuf,
			writer: invalidMovieWriter,
		},
	}

	for _, Case := range cases {
		// リクエスト生成
		req := httptest.NewRequest("POST", "/", &Case.buffer)

		// Content-Type 設定
		req.Header.Set("Content-Type", Case.writer.FormDataContentType())

		// Contextセット
		recorder := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(recorder)
		ginContext.Request = req
		ginContext.Set("JWT_PAYLOAD", jwt.MapClaims{
			"id": float64(10),
		})

		uploadMovieFileHandler.UploadMovieFile(ginContext)

		if recorder.Code != http.StatusBadRequest {
			t.Fatal("Usecase Error")
		}
	}
}

func TestUploadMovieFile_Usecase_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	movieUploadRepository := mock_repository.NewMockFileUpload(ctrl)
	postMovieUsecase := mock_usecase.NewMockIPostMovie(ctrl)
	uploadMovieFileHandler := handler.NewUploadMovieFile(movieRepository, thumbnailUploadRepository, movieUploadRepository, postMovieUsecase)

	// サムネイル
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	thumbnailFile, thumbnailFileErr := os.Open("./../TestThumbnail.jpg")
	if thumbnailFileErr != nil {
		fmt.Println(1)
		fmt.Println(thumbnailFileErr)
		return
	}

	movieFile, movieFileErr := os.Open("./../TestMovie.mp4")
	if movieFileErr != nil {
		fmt.Println(2)
		fmt.Println(movieFileErr)
		return
	}

	defer thumbnailFile.Close()
	defer movieFile.Close()

	thumbnailFileForm, thumbnailFileFormErr := w.CreateFormFile("uploadThumbnail", "TestThumbnail.jpg")
	if thumbnailFileFormErr != nil {
		fmt.Println(3)
		fmt.Println(thumbnailFileFormErr)
		return
	}

	_, thumbnailCopyErr := io.Copy(thumbnailFileForm, thumbnailFile)
	if thumbnailCopyErr != nil {
		fmt.Println(5)
		fmt.Println(thumbnailCopyErr)
		return
	}

	movieFileForm, movieFileFormErr := w.CreateFormFile("uploadMovie", "TestMovie.mp4")
	if movieFileFormErr != nil {
		fmt.Println(4)
		fmt.Println(movieFileFormErr)
		return
	}
	_, movieCopyErr := io.Copy(movieFileForm, movieFile)
	if movieCopyErr != nil {
		fmt.Println(6)
		fmt.Println(movieCopyErr)
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

	postMovieUsecase.EXPECT().PostMovie(gomock.Any()).Return(nil, errors.New("ERROR"))

	uploadMovieFileHandler.UploadMovieFile(ginContext)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatal("Usecase Error")
	}
}
