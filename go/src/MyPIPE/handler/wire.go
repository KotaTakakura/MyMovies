package handler

import(
	"MyPIPE/Controllers"
	"MyPIPE/Services"
	Iservice "MyPIPE/Interfaces/Services"
	"MyPIPE/domain/repository"
	"MyPIPE/infra"
	"MyPIPE/usecase"
	"github.com/google/wire"
)

func initVideoFileController() *Controllers.VideoFileController{
	wire.Build(
		Controllers.NewVideoFileController,
		Services.NewUploadToS3,
		wire.Bind(new(Iservice.IUploadToS3), new(*Services.UploadToS3)),
	)
	return &Controllers.VideoFileController{}
}

func InitRandomUser() *usecase.RandomUser{
	wire.Build(
		usecase.NewRandomUser,
		infra.NewUserPersistence,
		wire.Bind(new(repository.UserRepository), new(*infra.UserPersistence)),
	)
	return &usecase.RandomUser{}
}