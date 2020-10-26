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

type ChangeUserProfileImage struct {
	UserRepository                repository.UserRepository
	UserProfileImageRepository    repository.UserProfileImageRepository
	ChangeUserProfileImageUsecase usecase.IChangeUserProfilieImage
}

func NewChangeUserProfileImage(userRepo repository.UserRepository, userProfileImageRepo repository.UserProfileImageRepository, changeUserProfileImageUsecase usecase.IChangeUserProfilieImage) *ChangeUserProfileImage {
	return &ChangeUserProfileImage{
		UserRepository:                userRepo,
		UserProfileImageRepository:    userProfileImageRepo,
		ChangeUserProfileImageUsecase: changeUserProfileImageUsecase,
	}
}

func (changeUserProfileImage ChangeUserProfileImage) ChangeUserProfileImage(c *gin.Context) {
	userIdUint := uint64(jwt.ExtractClaims(c)["id"].(float64))
	validationErrors := make(map[string]string)

	userId, userIdErr := model.NewUserID(userIdUint)
	if userIdErr != nil {
		validationErrors["user_id"] = userIdErr.Error()
	}

	var profileImage *model.UserProfileImage
	var profileImageErr error
	imageFile, imageHeader, imageFileErr := c.Request.FormFile("profileImage")
	if imageFileErr != nil {
		validationErrors["profile_image"] = imageFileErr.Error()
	}else{
		profileImage,profileImageErr = model.NewUserProfileImage(*imageHeader,imageFile)
		if profileImageErr != nil{
			validationErrors["profile_image"] = profileImageErr.Error()
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

	changeUserProfileImageDTO := usecase.NewChangeUserProfileImageDTO(userId, profileImage)
	changeProfileImageErr := changeUserProfileImage.ChangeUserProfileImageUsecase.ChangeUserProfileImage(changeUserProfileImageDTO)
	if changeProfileImageErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Profile Image Set Failed.",
			"messages": changeProfileImageErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Posted!",
	})
}
