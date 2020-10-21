package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UploadMovieFile struct{
	MovieRepository	repository.MovieRepository
	ThumbnailUploadRepository	repository.ThumbnailUploadRepository
	MovieUploadRepository	repository.FileUpload
	PostMovieUsecase	usecase.IPostMovie
}

func NewUploadMovieFile(
	movieRepository	repository.MovieRepository,
	thumbnailUploadRepository	repository.ThumbnailUploadRepository,
	movieUploadRepository	repository.FileUpload,
	postMovieUsecase	usecase.IPostMovie,
	)*UploadMovieFile{
	return &UploadMovieFile{
		MovieRepository:           movieRepository,
		ThumbnailUploadRepository: thumbnailUploadRepository,
		MovieUploadRepository:     movieUploadRepository,
		PostMovieUsecase:          postMovieUsecase,
	}
}

func (uploadMovieFile UploadMovieFile)UploadMovieFile(c *gin.Context){

	//バリデーションエラー格納
	validationErrors := make(map[string]string)
	//ユーザーID取得
	iuserId := uint64(jwt.ExtractClaims(c)["id"].(float64))
	userId,userIdErr := model.NewUserID(iuserId)
	//ユーザーID valid
	if userIdErr != nil{
		validationErrors["user_id"] = userIdErr.Error()
	}


	//バリデーションチェック
	if len(validationErrors) != 0{
		validationErrors,_ := json.Marshal(validationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Validation Error.",
			"messages": string(validationErrors),
		})
		c.Abort()
		return
	}

	//動画ファイル取得&バリデーション
	file,header, fileErr :=  c.Request.FormFile("uploadMovie")
	if fileErr != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "MovieFileUpload Error.",
			"messages": fileErr.Error(),
		})
		c.Abort()
		return
	}

	//動画サムネイル取得&バリデーション
	thumbnail,thumbnailHeader, thumbnailErr :=  c.Request.FormFile("uploadThumbnail")
	if thumbnailErr != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "MovieThumbnailUpload Error.",
			"messages": thumbnailErr.Error(),
		})
		c.Abort()
		return
	}

	//動画投稿用のDTO
	postMovieDTO := usecase.NewPostMovieDTO(file,*header,thumbnail,*thumbnailHeader,userId)

	newMovieModel,err := uploadMovieFile.PostMovieUsecase.PostMovie(postMovieDTO)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "PostMovie Error.",
			"messages": err.Error(),
		})
		c.Abort()
		return
	}

	jsonNewMovieModel,_ := json.Marshal(newMovieModel)
	c.JSON(http.StatusOK, string(jsonNewMovieModel))
}
