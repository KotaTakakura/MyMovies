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

type UpdateMovieStatus struct {
	UpdateMovieStatusUsecase usecase.IUpdateMovie
}

func NewUpdateMovieStatus(updateMovieStatusUsecase usecase.IUpdateMovie) *UpdateMovieStatus {
	return &UpdateMovieStatus{
		UpdateMovieStatusUsecase: updateMovieStatusUsecase,
	}
}

func (u UpdateMovieStatus) UpdateMovieStatus(c *gin.Context) {
	requestBodyBuf := make([]byte, 2048)
	n, _ := c.Request.Body.Read(requestBodyBuf)
	request := string(requestBodyBuf[0:n])

	var updateMovieStatusJson UpdateMovieStatusJson
	err := json.Unmarshal([]byte(request), &updateMovieStatusJson)
	if err != nil {
		return
	}

	if updateMovieStatusJson.Type == "SubscriptionConfirmation" {
		sess := session.Must(session.NewSession())
		svc := sns.New(sess)
		_, err := svc.ConfirmSubscription(&sns.ConfirmSubscriptionInput{
			Token:    &updateMovieStatusJson.Token,
			TopicArn: &updateMovieStatusJson.TopicArn,
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

	movieIdChangeToUint64err := updateMovieStatusJson.ChangeMovieIdStringToUint64()
	if movieIdChangeToUint64err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Change Status Failed.",
			"messages": movieIdChangeToUint64err.Error(),
		})
		c.Abort()
		return
	}

	movieId, movieIdErr := model.NewMovieID(updateMovieStatusJson.MovieID)
	if movieIdErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Change Status Failed.",
			"messages": movieIdErr.Error(),
		})
		c.Abort()
		return
	}

	updateStatusDTO := usecase.NewUpdateStatusDTO(movieId)
	result := u.UpdateMovieStatusUsecase.UpdateStatus(updateStatusDTO)
	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Change Status Failed.",
			"messages": result.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK!",
	})
}

type UpdateMovieStatusJson struct {
	MovieID       uint64 `json:"movie_id"`
	MovieIDString string `json:"Message"`
	Type          string `json:"Type"`
	TopicArn      string `json:"TopicArn"`
	Token         string `json:"Token"`
}

func (u *UpdateMovieStatusJson) ChangeMovieIdStringToUint64() error {
	var err error
	u.MovieID, err = strconv.ParseUint(u.MovieIDString, 10, 64)
	if err != nil {
		return err
	}
	return nil
}

func (u UpdateMovieStatus) UpdateMovieStatusError(c *gin.Context) {
	requestBodyBuf := make([]byte, 2048)
	n, _ := c.Request.Body.Read(requestBodyBuf)
	request := string(requestBodyBuf[0:n])

	var updateMovieStatusErrorJson UpdateMovieStatusErrorJson
	err := json.Unmarshal([]byte(request), &updateMovieStatusErrorJson)
	if err != nil {
		return
	}

	if updateMovieStatusErrorJson.Type == "SubscriptionConfirmation" {
		sess := session.Must(session.NewSession())
		svc := sns.New(sess)
		_, err := svc.ConfirmSubscription(&sns.ConfirmSubscriptionInput{
			Token:    &updateMovieStatusErrorJson.Token,
			TopicArn: &updateMovieStatusErrorJson.TopicArn,
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

	movieIdChangeToUint64err := updateMovieStatusErrorJson.ChangeMovieIdStringToUint64()
	if movieIdChangeToUint64err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Change Status Failed.",
			"messages": movieIdChangeToUint64err.Error(),
		})
		c.Abort()
		return
	}

	movieId, movieIdErr := model.NewMovieID(updateMovieStatusErrorJson.MovieID)
	if movieIdErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Change Status Failed.",
			"messages": movieIdErr.Error(),
		})
		c.Abort()
		return
	}

	updateStatusErrorDTO := usecase.NewUpdateStatusErrorDTO(movieId)
	result := u.UpdateMovieStatusUsecase.UpdateStatusError(updateStatusErrorDTO)
	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Change Status Failed.",
			"messages": result.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK!",
	})
}

type UpdateMovieStatusErrorJson struct {
	MovieID       uint64 `json:"movie_id"`
	MovieIDString string `json:"Message"`
	Type          string `json:"Type"`
	TopicArn      string `json:"TopicArn"`
	Token         string `json:"Token"`
}

func (u *UpdateMovieStatusErrorJson) ChangeMovieIdStringToUint64() error {
	var err error
	u.MovieID, err = strconv.ParseUint(u.MovieIDString, 10, 64)
	if err != nil {
		return err
	}
	return nil
}
