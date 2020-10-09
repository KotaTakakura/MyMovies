package test

import (
	mock_factory "MyPIPE/Test/mock/factory"
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestPostMovie(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postMovieDTO := usecase.PostMovieDTO{
		UserID:	model.UserID(10),
	}

	MovieRepository := mock_repository.NewMockMovieRepository(ctrl)
	//mockgen -source domain/repository/MovieRepository.go -destination Test/mock/repository/mock_movieRepository.go

	FileRepository := mock_repository.NewMockFileUpload(ctrl)
	//mockgen -source domain/repository/FileUpload.go  -destination Test/mock/repository/mock_fileUpload.go

	ThumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	//mockgen -source domain/repository/ThumbnailUploadRepository.go  -destination Test/mock/repository/mock_thumbnailUpload.go

	MovieFactory := mock_factory.NewMockIMovieModelFactory(ctrl)
	//mockgen -source domain/factory/IMovieModel.go -destination Test/mock/factory/mock_movieFactory.go

	MovieFactory.EXPECT().
		CreateMovieModel(postMovieDTO.UserID,postMovieDTO.FileHeader,postMovieDTO.ThumbnailHeader).
		Return(&model.Movie{
			UserID: postMovieDTO.UserID,
	},nil)

	MovieRepository.EXPECT().Save(gomock.Any()).DoAndReturn(func(movie model.Movie)(*model.Movie,error){
		if movie.UserID != postMovieDTO.UserID{
			t.Error("Invalid UserID Name.")
			return nil,nil
		}
			return &model.Movie{},nil
	})

	FileRepository.EXPECT().Upload(postMovieDTO.File,postMovieDTO.FileHeader,model.Movie{}.ID).Return(nil)
	ThumbnailUploadRepository.EXPECT().Upload(postMovieDTO.Thumbnail,model.Movie{}).Return(nil)

	postMovieUsecase := usecase.NewPostMovie(FileRepository,ThumbnailUploadRepository,MovieRepository,MovieFactory)
	_,result := postMovieUsecase.PostMovie(&postMovieDTO)
	if result != nil{
		t.Error("PostMovie Usecase Error.")
	}
}
