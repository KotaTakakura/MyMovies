package handler

import (
	"MyPIPE/domain/factory"
	infra "MyPIPE/infra"
	support "MyPIPE/infra/UploadMovieFile"
	"MyPIPE/usecase"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadMovieFile(c *gin.Context){
	newMovie := factory.NewMovieModelFactory()
	uploadUsecase := usecase.NewPostMovie(support.NewUploadToAmazonS3(),infra.NewMoviePersistence(),*newMovie)

	iuserId := uint64(jwt.ExtractClaims(c)["id"].(float64))
	file,header, _ :=  c.Request.FormFile("uploadImage")

	postMovieDTO := usecase.PostMovieDTO{
		File: file,
		FileHeader: *header,
		UserID:	iuserId,
		DisplayName: c.PostForm("display_name"),
	}

	err := uploadUsecase.PostMovie(file,postMovieDTO)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Validation Error.",
			"messages": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"result": "OK",
		"messages": "OK",
	})
}
