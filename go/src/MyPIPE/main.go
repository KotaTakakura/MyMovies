package main

import (
	"MyPIPE/domain/model"
	"MyPIPE/handler"
	"MyPIPE/infra"
	"MyPIPE/usecase"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"time"
)

func init() {
	const location = "Asia/Tokyo"
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

type login struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"text":     "Hello World.",
	})
}

func main() {
	// the jwt middleware
	authMiddleware, err := authMiddlewareByJWT()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	router := gin.Default()
	router.POST("/login", authMiddleware.LoginHandler)
	router.GET("/refresh_token", authMiddleware.RefreshHandler)
	router.POST("/new", handler.TemporaryRegisterUser)
	router.POST("/register", handler.RegisterUser)

	auth := router.Group("/auth/api/v1")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.POST("/comments", handler.PostComment)
		auth.GET("/hello", helloHandler)
		auth.POST("/movie", handler.UploadMovieFile)
		auth.POST("/evaluates", handler.EvaluateMovie)
		auth.POST("/play-lists",handler.CreatePlayList)
	}

	router.GET("/test", func(c *gin.Context) {
		c.String(200, "fesfes.")
	})

	router.Run()
}

func authMiddlewareByJWT() (*jwt.GinJWTMiddleware, error){
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       os.Getenv("Realm"),
		Key:         []byte(os.Getenv("JWT_SECRET_KEY")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(model.UserID); ok {
				return jwt.MapClaims{
					identityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			//claims := jwt.ExtractClaims(c)
			//return &User{
			//	UserName: claims[identityKey].(string),
			//}
			return true
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			email, _ := model.NewUserEmail(loginVals.Email)
			password := loginVals.Password

			userRepository := infra.NewUserPersistence()
			userExistsUsecaes := usecase.NewUserExists(userRepository)

			userExists, err := userExistsUsecaes.CheckUserExistsForAuth(email, password)

			if userExists != nil && err == nil {
				return userExists.ID, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		//TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}
