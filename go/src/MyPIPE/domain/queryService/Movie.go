package queryService

import "MyPIPE/domain/model"

type Movie interface {
	UploadedMovies(userId model.UserID)
}
