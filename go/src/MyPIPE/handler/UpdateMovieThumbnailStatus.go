package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateMovieThumbnailStatus struct {
	UpdateMovieThumbnailStatusUsecase usecase.IUpdateMovie
}

func NewUpdateMovieThumbnailStatus(updateMovieThumbnailStatusUsecase usecase.IUpdateMovie) *UpdateMovieThumbnailStatus {
	return &UpdateMovieThumbnailStatus{
		UpdateMovieThumbnailStatusUsecase: updateMovieThumbnailStatusUsecase,
	}
}

func (u UpdateMovieThumbnailStatus) UpdateMovieThumbnailStatus(c *gin.Context) {
	requestBodyBuf := make([]byte, 2048)
	n, _ := c.Request.Body.Read(requestBodyBuf)
	request := string(requestBodyBuf[0:n])

	var updateMovieThumbnailStatusJson UpdateMovieThumbnailStatusJson
	err := json.Unmarshal([]byte(request), &updateMovieThumbnailStatusJson)
	if err != nil {
		return
	}

	if updateMovieThumbnailStatusJson.Type == "SubscriptionConfirmation" {
		sess := session.Must(session.NewSession())
		svc := sns.New(sess)
		_, err := svc.ConfirmSubscription(&sns.ConfirmSubscriptionInput{
			Token:    &updateMovieThumbnailStatusJson.Token,
			TopicArn: &updateMovieThumbnailStatusJson.TopicArn,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result":   "Failed.",
				"messages": err.Error(),
			})
			c.Abort()
			return
		}
	}

	movieIdChangeToUint64err := updateMovieThumbnailStatusJson.ChangeMovieIdStringToUint64()
	if movieIdChangeToUint64err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Change Status Failed.",
			"messages": movieIdChangeToUint64err.Error(),
		})
		c.Abort()
		return
	}

	movieId, movieIdErr := model.NewMovieID(updateMovieThumbnailStatusJson.MovieID)
	if movieIdErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Change Thumbnail Status Failed.",
			"messages": movieIdErr.Error(),
		})
		c.Abort()
		return
	}

	updateThumbnailStatusDTO := usecase.NewUpdateThumbnailStatusDTO(movieId)
	result := u.UpdateMovieThumbnailStatusUsecase.UpdateThumbnailStatus(updateThumbnailStatusDTO)
	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Change Thumbnail Status Failed.",
			"messages": result.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK!",
	})
}

type UpdateMovieThumbnailStatusJson struct {
	MovieID       uint64 `json:"movie_id"`
	MovieIDString string `json:"Message"`
	Type          string `json:"Type"`
	TopicArn      string `json:"TopicArn"`
	Token         string `json:"Token"`
}

func (u *UpdateMovieThumbnailStatusJson) ChangeMovieIdStringToUint64() error {
	var err error
	u.MovieID, err = strconv.ParseUint(u.MovieIDString, 10, 64)
	if err != nil {
		return err
	}
	return nil
}
