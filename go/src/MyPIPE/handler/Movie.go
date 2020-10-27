package handler

import (
	"MyPIPE/domain/model"
	queryService_uploadMovies "MyPIPE/domain/queryService/UploadedMovies"
	"MyPIPE/domain/repository"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Movie struct {
	UploadMovieQueryService   queryService_uploadMovies.UploadedMovies
	UploadMovieUsecase        usecase.IUploadedMovies
	MovieRepository           repository.MovieRepository
	UpdateMovieUsecase        usecase.IUpdateMovie
	ThumbnailUploadRepository repository.ThumbnailUploadRepository
	ChangeThumbnailUsecase    usecase.IChangeThumbnail
}

func NewMovie(uploadMovieQueryService queryService_uploadMovies.UploadedMovies, uploadMovieUsecase usecase.IUploadedMovies, movieRepository repository.MovieRepository, updateMovieUsecase usecase.IUpdateMovie, thumbnailUploadRepository repository.ThumbnailUploadRepository, changeThumbnailUsecase usecase.IChangeThumbnail) *Movie {
	return &Movie{
		UploadMovieQueryService:   uploadMovieQueryService,
		UploadMovieUsecase:        uploadMovieUsecase,
		MovieRepository:           movieRepository,
		UpdateMovieUsecase:        updateMovieUsecase,
		ThumbnailUploadRepository: thumbnailUploadRepository,
		ChangeThumbnailUsecase:    changeThumbnailUsecase,
	}
}

func (movie Movie) GetUploadedMovies(c *gin.Context) {
	userId := jwt.ExtractClaims(c)["id"]
	iuserId := uint64(userId.(float64))

	userIdModel, err := model.NewUserID(iuserId)
	if err != nil {
		return
	}

	result := movie.UploadMovieUsecase.Get(userIdModel)
	jsonResult, jsonMarshalErr := json.Marshal(result)
	if jsonMarshalErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Internal Server Error.",
			"messages": jsonMarshalErr.Error(),
		})
		c.Abort()
		return
	} else {
		c.JSON(http.StatusOK, string(jsonResult))
	}
}

func (movie Movie) UpdateMovie(c *gin.Context) {
	userId := jwt.ExtractClaims(c)["id"]
	iuserId := uint64(userId.(float64))

	userIdModel, err := model.NewUserID(iuserId)

	var updateMovieDTO UpdateMovieDTO
	c.Bind(&updateMovieDTO)

	validationErrors := make(map[string]string)

	if err != nil {
		validationErrors["user_id"] = err.Error()
	}

	movieId, movieIdErr := model.NewMovieID(updateMovieDTO.MovieID)
	if movieIdErr != nil {
		validationErrors["movie_id"] = movieIdErr.Error()
	}

	displayName, displayNameErr := model.NewMovieDisplayName(updateMovieDTO.DisplayName)
	if displayNameErr != nil {
		validationErrors["display_name"] = displayNameErr.Error()
	}

	description, descriptionErr := model.NewMovieDescription(updateMovieDTO.Description)
	if descriptionErr != nil {
		validationErrors["description"] = descriptionErr.Error()
	}

	public, publicErr := model.NewMoviePublic(updateMovieDTO.Public)
	if publicErr != nil {
		validationErrors["public"] = publicErr.Error()
	}

	status, statusErr := model.NewMovieStatus(updateMovieDTO.Status)
	if statusErr != nil {
		validationErrors["status"] = statusErr.Error()
	}

	if len(validationErrors) != 0 {
		validationErrors, _ := json.Marshal(validationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Validation Error.",
			"messages": string(validationErrors),
		})
		c.Abort()
		return
	}

	updateDTO := usecase.UpdateDTO{
		UserID:      userIdModel,
		MovieID:     movieId,
		DisplayName: displayName,
		Description: description,
		Public:      public,
		Status:      status,
	}

	result, updateMovieUsecaseErr := movie.UpdateMovieUsecase.Update(&updateDTO)
	if updateMovieUsecaseErr != nil {
		jsonUpdateMovieUsecaseErr, _ := json.Marshal(updateMovieUsecaseErr.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Server Error.",
			"messages": string(jsonUpdateMovieUsecaseErr),
		})
		c.Abort()
		return
	}

	updatedData, _ := json.Marshal(result)
	c.JSON(http.StatusOK, gin.H{
		"result":      "Success.",
		"messages":    "OK",
		"updatedData": string(updatedData),
	})
}

type UpdateMovieDTO struct {
	UserID      uint64
	MovieID     uint64 `json:"movie_id"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Public      uint   `json:"public"`
	Status      uint   `json:"status"`
}

func (movie Movie) ChangeThumbnail(c *gin.Context) {

	//バリデーションエラー格納
	validationErrors := make(map[string]string)
	//ユーザーID取得
	iuserId := uint64(jwt.ExtractClaims(c)["id"].(float64))
	userId, userIdErr := model.NewUserID(iuserId)
	//ユーザーID valid
	if userIdErr != nil {
		validationErrors["user_id"] = userIdErr.Error()
	}

	requestMovieId := c.PostForm("movie_id")
	imovieId, _ := strconv.ParseUint(requestMovieId, 10, 64)
	movieId, movieIdErr := model.NewMovieID(imovieId)
	if movieIdErr != nil {
		validationErrors["movie_id"] = movieIdErr.Error()
	}

	thumbnailFile, thumbnailFileHeader, thumbnailFileErr := c.Request.FormFile("uploadThumbnail")
	var thumbnail *model.MovieThumbnail
	var thumbnailErr error
	if thumbnailFileErr != nil {
		validationErrors["thumbnail"] = thumbnailFileErr.Error()
	} else {
		thumbnail, thumbnailErr = model.NewMovieThumbnail(thumbnailFile, *thumbnailFileHeader)
		if thumbnailErr != nil {
			validationErrors["thumbnail"] = thumbnailErr.Error()
		}
	}
	if len(validationErrors) != 0 {
		validationErrors, _ := json.Marshal(validationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Validation Error.",
			"messages": string(validationErrors),
		})
		c.Abort()
		return
	}

	changeThumbnailDTO := usecase.NewChangeThumbnailDTO(userId, movieId, *thumbnail)
	changeThumbnailUsecaseErr := movie.ChangeThumbnailUsecase.ChangeThumbnail(changeThumbnailDTO)
	if changeThumbnailUsecaseErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Server Error.",
			"messages": changeThumbnailUsecaseErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":   "OK",
		"messages": "OK",
	})

}
