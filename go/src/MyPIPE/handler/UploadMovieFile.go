package handler

import (
	"MyPIPE/domain/model"
	infra "MyPIPE/infra"
	support "MyPIPE/infra/UploadMovieFile"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"MyPIPE/infra/factory"
)

func UploadMovieFile(c *gin.Context){
	validationErrors := make(map[string]string)
	iuserId := uint64(jwt.ExtractClaims(c)["id"].(float64))
	userId,userIdErr := model.NewUserID(iuserId)
	if userIdErr != nil{
		validationErrors["user_id"] = userIdErr.Error()
	}

	displayName,displayNameErr := model.NewMovieDisplayName(c.PostForm("display_name"))
	if displayNameErr != nil{
		validationErrors["display_name"] = displayNameErr.Error()
	}

	if len(validationErrors) != 0{
		validationErrors,_ := json.Marshal(validationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Validation Error.",
			"messages": string(validationErrors),
		})
		c.Abort()
		return
	}

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
		UserID:	userId,
		DisplayName: displayName,
	}

	newMovie := factory.NewMovieModelFactory()
	uploadUsecase := usecase.NewPostMovie(support.NewUploadToAmazonS3(),infra.NewMoviePersistence(),*newMovie)
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
