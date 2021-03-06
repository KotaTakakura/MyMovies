package main

import (
	"MyPIPE/domain/model"
	"MyPIPE/handler"
	"MyPIPE/infra"
	uploadMovieRepository_infra "MyPIPE/infra/UploadMovieFile"
	support "MyPIPE/infra/UploadThumbnail"
	uploadThumbnailRepository_infra "MyPIPE/infra/UploadThumbnail"
	"MyPIPE/infra/factory"
	queryService_infra "MyPIPE/infra/queryService"
	"MyPIPE/usecase"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
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
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims[identityKey],
		"text":   "Hello World.",
	})
}

func main() {

	userRepository := infra.NewUserPersistence()
	userProfileImageRepository := infra.NewUserProfileImagePersistence()
	movieEvaluationRepository := infra.NewMovieEvaluatePersistence()
	commentRepository := infra.NewCommentPersistence()
	movieRepository := infra.NewMoviePersistence()
	movieStatusRepository := infra.NewMovieStatusPersistence()
	playListRepository := infra.NewPlayListPersistence()
	playListMovieRepository := infra.NewPlayListMoviePersistence()
	thumbnailUploadRepository := uploadThumbnailRepository_infra.NewUploadThumbnailToAmazonS3()
	movieUploadRepository := uploadMovieRepository_infra.NewUploadToAmazonS3()
	temporaryRegisterMailRepository := infra.NewTemporaryRegisterMailRepository()

	updateMovieUsecase := usecase.NewUpdateMovie(movieRepository,movieStatusRepository)

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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://www.frommymovies.com"},
		AllowMethods:     []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin", "Content-type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/login", authMiddleware.LoginHandler)
	router.GET("/refresh_token", authMiddleware.RefreshHandler)

	userTemporaryRegistrationUsecase := usecase.NewUserTemporaryRegistration(userRepository, temporaryRegisterMailRepository)
	userRegisterUsecase := usecase.NewUserRegister(userRepository)
	authorizationHandler := handler.NewAuthorization(userRepository, userTemporaryRegistrationUsecase, userRegisterUsecase)
	router.POST("/new", authorizationHandler.TemporaryRegisterUser)
	router.POST("/register", authorizationHandler.RegisterUser)

	checkUserAlreadyLikedMovieUsecase := usecase.NewCheckUserAlreadyLikedMovie(movieEvaluationRepository)
	checkUserAlreadyLikedMovieHandler := handler.NewCheckUserAlreadyLikedMovie(movieEvaluationRepository, checkUserAlreadyLikedMovieUsecase)
	router.GET("/evaluated", checkUserAlreadyLikedMovieHandler.CheckUserAlreadyLikedMovie)

	api := router.Group("/api/v1")

	resetPasswordEmailRepository := infra.NewResetPasswordEmail()
	setPasswordRememberTokenUsecase := usecase.NewSetPasswordRememberToken(userRepository, resetPasswordEmailRepository)
	setPasswordRememberTokenHandler := handler.NewSetPasswordRememberToken(setPasswordRememberTokenUsecase)
	api.POST("/remember", setPasswordRememberTokenHandler.SetPasswordRememberToken)

	resetPasswordusecase := usecase.NewResetPassword(userRepository)
	resetPasswordHandler := handler.NewResetPassword(resetPasswordusecase)
	api.POST("/reset", resetPasswordHandler.ResetPassword)

	commentQueryService := queryService_infra.NewCommentQueryService()
	getCommentsUsecase := usecase.NewGetMovieAndComments(commentQueryService)
	getMovieAndCommentsHandler := handler.NewGetMovieAndComments(commentQueryService, getCommentsUsecase)
	api.GET("/movie-and-comments", getMovieAndCommentsHandler.GetMovieAndComments)

	indexMovieQueryService := queryService_infra.NewIndexMovie()
	indexMovieUsecase := usecase.NewIndexMovie(indexMovieQueryService)
	indexMovieHandler := handler.NewIndexMovie(indexMovieQueryService, indexMovieUsecase)
	api.GET("/index-movies", indexMovieHandler.IndexMovie)

	updateMovieStatus := handler.NewUpdateMovieStatus(updateMovieUsecase)
	api.POST("/movie-status", updateMovieStatus.UpdateMovieStatus)
	api.POST("/movie-status-error", updateMovieStatus.UpdateMovieStatusError)

	updateMovieThumbnailStatusHandler := handler.NewUpdateMovieThumbnailStatus(updateMovieUsecase)
	api.POST("/movie-thumbnail-status", updateMovieThumbnailStatusHandler.UpdateMovieThumbnailStatus)

	changeEmailUsecase := usecase.NewChangeEmail(userRepository)
	changeEmailHandler := handler.NewChangeEmail(changeEmailUsecase)
	api.PUT("/email", changeEmailHandler.ChangeEmail)

	auth := router.Group("/auth/api/v1")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		changeEmailMailRepository := infra.NewChangeEmailMail()
		setChangeEmailTokenUsecase := usecase.NewSetEmailChangeToken(userRepository, changeEmailMailRepository)
		setChangeEmailTokenHandler := handler.NewSetEmailChangeToken(setChangeEmailTokenUsecase)
		auth.POST("/email", setChangeEmailTokenHandler.SetEmailChangeToken)

		getLoggedInUserDataQueryService := queryService_infra.NewGetLoggedInUserData()
		getLoggedInUserDataUsecase := usecase.NewGetLoggedInUserData(getLoggedInUserDataQueryService)
		getLoggedInUserDataHandler := handler.NewGetLoggedInUserData(getLoggedInUserDataQueryService, getLoggedInUserDataUsecase)
		auth.GET("/user", getLoggedInUserDataHandler.GetLoggedInUserData)

		changeUserNameUsecase := usecase.NewChangeUserName(userRepository)
		changeUserNameHandler := handler.NewChangeUserName(userRepository, changeUserNameUsecase)
		auth.PUT("/user-name", changeUserNameHandler.ChangeUserName)

		changePasswordUsecase := usecase.NewChangePassword(userRepository)
		changePasswordHandler := handler.NewChangePassword(userRepository, *changePasswordUsecase)
		auth.PUT("/password", changePasswordHandler.ChangePassword)

		changeUserProfileImageUsecase := usecase.NewChangeUserProfileImage(userRepository, userProfileImageRepository)
		changeUserProfileImageHandler := handler.NewChangeUserProfileImage(userRepository, userProfileImageRepository, changeUserProfileImageUsecase)
		auth.PUT("/profile-image", changeUserProfileImageHandler.ChangeUserProfileImage)

		postCommentUsecase := usecase.NewPostComment(commentRepository, movieRepository)
		postCommentHandler := handler.NewPostComment(commentRepository, movieRepository, postCommentUsecase)
		auth.POST("/comments", postCommentHandler.PostComment)

		deleteCommentUsecase := usecase.NewDeleteComment(commentRepository)
		deleteCommentHandler := handler.NewDeleteComment(deleteCommentUsecase)
		auth.DELETE("/comments", deleteCommentHandler.DeleteComment)

		movieFactory := factory.NewMovieModelFactory()
		postMovieUsecase := usecase.NewPostMovie(movieUploadRepository, thumbnailUploadRepository, movieRepository, movieFactory)
		uploadMovieHandler := handler.NewUploadMovieFile(movieRepository, thumbnailUploadRepository, movieUploadRepository, postMovieUsecase)
		auth.POST("/movie", uploadMovieHandler.UploadMovieFile)

		uploadedMoviesQueryService := queryService_infra.NewUploadedMovies()
		uploadedMoviesUsecase := usecase.NewUploadedMovies(uploadedMoviesQueryService)
		thumbnailUploadRepository := support.NewUploadThumbnailToAmazonS3()
		changeThumbnailUsecase := usecase.NewChangeThumbnail(movieRepository, thumbnailUploadRepository)
		movieHandler := handler.NewMovie(
			uploadedMoviesQueryService,
			uploadedMoviesUsecase,
			movieRepository,
			updateMovieUsecase,
			thumbnailUploadRepository,
			changeThumbnailUsecase,
		)
		auth.PUT("/movie", movieHandler.UpdateMovie)
		auth.PUT("/thumbnail", movieHandler.ChangeThumbnail)
		auth.GET("/movies", movieHandler.GetUploadedMovies)

		evaluateMovieUsecase := usecase.NewEvaluateUsecase(movieRepository, movieEvaluationRepository)
		evaluateMovieHandler := handler.NewEvaluateMovie(movieRepository, movieEvaluationRepository, evaluateMovieUsecase)
		auth.POST("/evaluates", evaluateMovieHandler.EvaluateMovie)

		createPlayListUsecase := usecase.NewCreatePlayList(userRepository, playListRepository)
		createPlayListHandler := handler.NewCreatePlayList(userRepository, playListRepository, createPlayListUsecase)
		auth.POST("/play-lists", createPlayListHandler.CreatePlayList)

		indexPlayListsInMyPageQueryService := queryService_infra.NewIndexPlayListsInMyPage()
		indexPlayListsInMyPageUsecase := usecase.NewIndexPlayListsInMyPage(indexPlayListsInMyPageQueryService)
		indexPlayListsInMyPageHandler := handler.NewIndexPlayListsInMyPage(indexPlayListsInMyPageQueryService, indexPlayListsInMyPageUsecase)
		auth.GET("/play-lists", indexPlayListsInMyPageHandler.IndexPlayListsInMyPage)

		updatePlayListUsecase := usecase.NewUpdatePlayList(playListRepository, movieRepository)
		updatePlayListHandler := handler.NewUpdatePlayListHandler(updatePlayListUsecase)
		auth.PUT("/play-lists", updatePlayListHandler.Update)

		deletePlayListUsecase := usecase.NewDeletePlayList(playListRepository)
		deletePlayListHandler := handler.NewDeletePlayList(playListRepository, deletePlayListUsecase)
		auth.DELETE("/play-lists", deletePlayListHandler.DeletePlayList)

		playListMovieFactory := factory.NewPlayListMovieFactory()
		addPlayListItemUsecase := usecase.NewAddPlayListItem(playListRepository, playListMovieRepository, playListMovieFactory)
		deletePlayListMovieUsecase := usecase.NewDeletePlayListMovie(playListRepository, playListMovieRepository)
		addPlayListMovieHandler := handler.NewPlayListItem(playListRepository, playListMovieRepository, playListMovieFactory, addPlayListItemUsecase, deletePlayListMovieUsecase)
		auth.POST("/play-list-items", addPlayListMovieHandler.AddPlayListMovie)
		auth.DELETE("/play-list-items", addPlayListMovieHandler.DeletePlayListMovie)

		changeOrderOfPlayListMoviesUsecase := usecase.NewChangeOrderOfPlayListMovies(playListMovieRepository)
		changeOrderOfPlayListMoviesHandler := handler.NewChangeOrderOfPlayListMovies(playListMovieRepository, changeOrderOfPlayListMoviesUsecase)
		auth.PUT("/play-list-items", changeOrderOfPlayListMoviesHandler.ChangeOrderOfPlayListMovies)

		auth.POST("/follows", handler.FollowUser)

		indexPlaylistMoviesQueryService := queryService_infra.NewIndexPlayListMovieInMyPage()
		indexPlaylistMoviesUsecase := usecase.NewIndexPlayListItemInMyPage(indexPlaylistMoviesQueryService)
		indexPlayListMoviesHandler := handler.NewIndexPlaylistMovies(indexPlaylistMoviesQueryService, indexPlaylistMoviesUsecase)
		auth.GET("/play-list-items/:play_list_id", indexPlayListMoviesHandler.IndexPlaylistMovies)

		indexPlayListInMovieListPageQueryService := queryService_infra.NewIndexPlayListInMovieListPage()
		indexPlayListInMovieListPageUsecase := usecase.NewIndexPlayListInMovieListPage(indexPlayListInMovieListPageQueryService)
		indexPlayListInMovieListPageHandler := handler.NewIndexPlayListInMovieListPage(indexPlayListInMovieListPageQueryService, indexPlayListInMovieListPageUsecase)
		auth.GET("play-lists/:movie_id", indexPlayListInMovieListPageHandler.IndexPlayListInMovieListPage)

		deleteUserUsecase := usecase.NewDeleteUser(userRepository)
		deleteUserHandler := handler.NewDeleteUser(deleteUserUsecase)
		auth.DELETE("/user", deleteUserHandler.DeleteUser)

		deleteMovieUsecase := usecase.NewDeleteMovie(movieRepository)
		deleteMovieHandler := handler.NewDeleteMovie(deleteMovieUsecase)
		auth.DELETE("/movie", deleteMovieHandler.DeleteMovie)
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Healthy."})
	})

	router.Run()
}

func authMiddlewareByJWT() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       os.Getenv("Realm"),
		Key:         []byte(os.Getenv("JWT_SECRET_KEY")),
		Timeout:     time.Hour * 24 * 30,
		MaxRefresh:  time.Hour * 24 * 30,
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
