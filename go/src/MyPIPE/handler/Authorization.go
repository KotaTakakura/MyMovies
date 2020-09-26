package handler

import (
	"MyPIPE/domain/model"
	domain_service "MyPIPE/domain/service/User"
	"MyPIPE/infra"
	"MyPIPE/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TemporaryRegisterUser(c *gin.Context) {
	userPersistence := infra.NewUserPersistence()
	userRegistration := usecase.NewUserTemporaryRegistration(userPersistence)

	var newUserInfo TemporaryRegisterUserJson
	var newUser model.User
	validationError := map[string]error{}
	validationErrorMessages := map[string]string{}

	c.Bind(&newUserInfo)
	newUser.Email, validationError["user_email"] = model.NewUserEmail(newUserInfo.Email)

	if validationError["user_email"] != nil {
		validationErrorMessages["user_email"] = validationError["user_email"].Error()
		c.JSON(http.StatusOK, gin.H{
			"result": "Validation Error",
			"message": validationErrorMessages,
		})
		c.Abort()
		return
	}

	err := userRegistration.TemporaryRegister(&newUser)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"message": "Temporary Registered!",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Temporary Registered!",
	})
}

type TemporaryRegisterUserJson struct{
	Email	string	`json:"email"`
}

func RegisterUser(c *gin.Context) {

	var newUserInfo RegisterUserJson
	c.Bind(&newUserInfo)

	validationErrors := map[string]error{}
	errorMessages := map[string]string{}
	validationErrorFlag := false

	userPersistence := infra.NewUserPersistence()
	userService := domain_service.NewUserService(userPersistence)
	userRegistration := usecase.NewUserRegister(userPersistence,userService)


	token := c.Query("token")

	var newUser model.User
	newUser.Name,validationErrors["user_name"] = model.NewUserName(newUserInfo.Name)

	newUser.Password,validationErrors["user_password"] = model.NewUserPassword(newUserInfo.Password)

	newUser.Token,validationErrors["user_token"] = model.NewUserToken(token)

	validationErrors["user_birthday"] = newUser.SetBirthday(newUserInfo.Birthday)

	for errorKey, errorContent := range validationErrors {
		if errorContent != nil{
			validationErrorFlag = true
			errorMessages[errorKey] = errorContent.Error()
		}
	}

	if validationErrorFlag {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Validation Error",
			"messages": errorMessages,
		})
		c.Abort()
		return
	}

	registerUserError := userRegistration.RegisterUser(&newUser)

	if registerUserError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Registration Error",
			"messages": registerUserError.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registered!",
	})
}

type RegisterUserJson struct{
	Name string	`json:"name"`
	Password string	`json:"password"`
	Birthday	string	`json:"birthday"`
}