package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"MyPIPE/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TemporaryRegisterUser(c *gin.Context) {
	userPersistence := infra.NewUserPersistence()
	userRegistration := usecase.NewUserTemporaryRegistration(userPersistence)

	var newUserInfo TemporaryRegisterUserJson
	var newUser model.User

	c.Bind(&newUserInfo)
	newUser.Email = model.NewUserEmail(newUserInfo.Email)	//TODO::バリデーション
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

	fmt.Println("|||||||||||")
	fmt.Print(newUserInfo)
	fmt.Println("|||||||||||")

	userPersistence := infra.NewUserPersistence()
	userRegistration := usecase.NewUserRegister(userPersistence)


	token := c.Query("token")

	var newUser model.User
	newUser.Name = model.NewUserName(newUserInfo.Name)	//TODO::バリデーション
	newUser.Password = model.NewUserPassword(newUserInfo.Password)	//TODO::バリデーション
	newUser.Token = model.NewUserToken(token)	//TODO::バリデーション（空文字のときを必ずエラーとして処理する）
	newUser.SetBirthday(newUserInfo.Birthday)	//TODO::バリデーション

	userRegistration.RegisterUser(&newUser)	//TODO::エラー処理

	c.JSON(http.StatusOK, gin.H{
		"message": "Registered!",
	})
}

type RegisterUserJson struct{
	Name string	`json:"name"`
	Password string	`json:"password"`
	Birthday	string	`json:"birthday"`
}