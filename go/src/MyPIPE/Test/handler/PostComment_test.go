package handler

import (
	mock_repository "MyPIPE/Test/mock/repository"
	mock_usecase "MyPIPE/Test/mock/usecase"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	jwt "github.com/appleboy/gin-jwt/v2"
	"MyPIPE/handler"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"strings"
	"testing"
	"fmt"
	"reflect"
)

func TestPostComment(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepository := mock_repository.NewMockCommentRepository(ctrl)
	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	postCommentUsecase := mock_usecase.NewMockIPostComment(ctrl)
	postCommentHandler := handler.NewPostComment(commentRepository,movieRepository,postCommentUsecase)

	trueCases := []struct {
		commentBody string
		userId uint64
		movieId uint64
	}{
		{commentBody: "TestCommentBody",userId: 10,movieId: 100},
	}

	falseCases := []struct {
		commentBody string
		userId uint64
		movieId uint64
	}{
		{commentBody: "",userId: 10,movieId: 100},
	}

	for _,trueCase := range trueCases{
		// ポストデータ
		bodyReader := strings.NewReader(`{"comment_body":"` + trueCase.commentBody + `","user_id":` + fmt.Sprint(trueCase.userId) +  `,"movie_id":` + fmt.Sprint(trueCase.movieId) + `}`)

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

		postCommentUsecase.EXPECT().PostComment(gomock.Any()).DoAndReturn(func(data interface{})error{
			if reflect.TypeOf(data) != reflect.TypeOf(&(usecase.PostCommentDTO{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*usecase.PostCommentDTO).UserID != model.UserID(trueCase.userId) {
				t.Fatal("UserID Not Match,")
			}
			if data.(*usecase.PostCommentDTO).MovieID != model.MovieID(trueCase.movieId) {
				t.Fatal("MovieID Not Match,")
			}
			if data.(*usecase.PostCommentDTO).Body != model.CommentBody(trueCase.commentBody) {
				t.Fatal("CommentBody Not Match,")
			}
			return nil
		})

		postCommentHandler.PostComment(ginContext)
	}

	for _,falseCase := range falseCases{
		// ポストデータ
		bodyReader := strings.NewReader(`{"comment_body":"` + falseCase.commentBody + `","user_id":` + fmt.Sprint(falseCase.userId) +  `,"movie_id":` + fmt.Sprint(falseCase.movieId) + `}`)

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

		postCommentHandler.PostComment(ginContext)
	}
}
