package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gin-gonic/gin"
	"net/http"
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
	fmt.Println(request)

	var updateMovieStatusJson UpdateMovieStatusJson
	err := json.Unmarshal([]byte(request), &updateMovieStatusJson)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(updateMovieStatusJson)

	if updateMovieStatusJson.Type == "SubscriptionConfirmation"{
		sess := session.Must(session.NewSession())
		svc := sns.New(sess)
		_,err := svc.ConfirmSubscription(&sns.ConfirmSubscriptionInput{
			Token:                     &updateMovieStatusJson.Token,
			TopicArn:                  &updateMovieStatusJson.TopicArn,
		})
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"result":   "Failed.",
				"messages": err.Error(),
			})
			c.Abort()
			return
		}
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
	MovieID uint64 `json:"movie_id"`
	Type string `json:"Type"`
	TopicArn string `json:"TopicArn"`
	Token string `json:"Token"`
}
