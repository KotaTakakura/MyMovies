package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func ChangeUserProfileImage(c *gin.Context){
	userIdUint := uint64(jwt.ExtractClaims(c)["id"].(float64))
	validationErrors := make(map[string]string)

	userId,userIdErr := model.NewUserID(userIdUint)
	if userIdErr != nil{
		validationErrors["user_id"] = userIdErr.Error()
	}

	imageFile,imageHeader,imageFileErr := c.Request.FormFile("profileImage")
	if imageFileErr != nil{
		validationErrors["profile_image"] = imageFileErr.Error()
		fmt.Println(imageFileErr.Error())
	}


	if len(validationErrors) != 0{
		validationErrors,_ :=  json.Marshal(validationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Validation Error.",
			"messages": string(validationErrors),
		})
		c.Abort()
		return
	}

	changeUserProfileImageDTO := usecase.NewChangeUserProfileImageDTO(userId,imageFile,imageHeader)
	userRepository := infra.NewUserPersistence()
	userProfileImageRepository := infra.NewUserProfileImagePersistence()
	changeUserProfileImageUsecase := usecase.NewChangeUserProfileImage(userRepository,userProfileImageRepository)
	changeProfileImageErr := changeUserProfileImageUsecase.ChangeUserProfileImage(changeUserProfileImageDTO)
	if changeProfileImageErr != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Profile Image Set Failed.",
			"messages": changeProfileImageErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Posted!",
	})

}
