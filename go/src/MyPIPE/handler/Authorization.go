package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"MyPIPE/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Authorization struct{
	UserRepository	repository.UserRepository
	UserTemporaryRegistrationUsecase	usecase.IUserTemporaryRegistration
	UserRegisterUsecase 	usecase.IUserRegister
}

func NewAuthorization(userRep repository.UserRepository,userTemporaryRegistrationUsecase usecase.IUserTemporaryRegistration,userRegisterUsecase usecase.IUserRegister)*Authorization{
	return &Authorization{
		UserRepository: userRep,
		UserTemporaryRegistrationUsecase: userTemporaryRegistrationUsecase,
		UserRegisterUsecase:	userRegisterUsecase,
	}
}

func (authorization Authorization)TemporaryRegisterUser(c *gin.Context) {

	var newUserInfo TemporaryRegisterUserJson
	var newUser model.User
	validationError := map[string]error{}
	validationErrorMessages := map[string]string{}

	c.Bind(&newUserInfo)
	newUser.Email, validationError["user_email"] = model.NewUserEmail(newUserInfo.Email)

	if validationError["user_email"] != nil {
		validationErrorMessages["user_email"] = validationError["user_email"].Error()
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Validation Error",
			"message": validationErrorMessages,
		})
		c.Abort()
		return
	}

	err := authorization.UserTemporaryRegistrationUsecase.TemporaryRegister(&newUser)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
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
	Email	string	`json:"user_email"`
}

func (authorization Authorization)RegisterUser(c *gin.Context) {

	var newUserInfo RegisterUserJson
	c.Bind(&newUserInfo)

	validationErrors := map[string]error{}
	errorMessages := map[string]string{}
	validationErrorFlag := false

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

	registerUserError := authorization.UserRegisterUsecase.RegisterUser(&newUser)

	if registerUserError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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
	Name string	`json:"user_name"`
	Password string	`json:"user_password"`
	Birthday	string	`json:"user_birthday"`
}