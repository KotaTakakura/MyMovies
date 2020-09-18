package handler

import (
	infra "MyPIPE/infra"
	support "MyPIPE/infra/UploadMovieFile"
	"MyPIPE/usecase"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"MyPIPE/infra/factory"
)

func UploadMovieFile(c *gin.Context){
	newMovie := factory.NewMovieModelFactory()
	uploadUsecase := usecase.NewPostMovie(support.NewUploadToAmazonS3(),infra.NewMoviePersistence(),*newMovie)

	iuserId := uint64(jwt.ExtractClaims(c)["id"].(float64))
	file,header, fileErr :=  c.Request.FormFile("uploadMovie")
	if fileErr != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Validation Error.",
			"messages": fileErr.Error(),
		})
		c.Abort()
		return
	}

	postMovieDTO := usecase.PostMovieDTO{
		File: file,
		FileHeader: *header,
		UserID:	iuserId,
		DisplayName: c.PostForm("display_name"),
	}

	err := uploadUsecase.PostMovie(postMovieDTO)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Validation Error.",
			"messages": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "OK",
		"messages": "OK",
	})
}
