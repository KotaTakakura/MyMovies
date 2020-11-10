package infra

import "MyPIPE/domain/model"

type MovieStatusPersistence struct{}

func NewMovieStatusPersistence()*MovieStatusPersistence{
	return &MovieStatusPersistence{}
}

func (m MovieStatusPersistence)Find(movieId model.MovieID)(*model.MovieStatusModel,error){
	db := ConnectGorm()
	defer db.Close()

	var movieStatusModel model.MovieStatusModel
	result := db.Table("movies").Where("id = ?",movieId).Take(&movieStatusModel)
	if result.Error != nil{
		return nil,result.Error
	}
	if result.RowsAffected == 0{
		return nil,nil
	}

	return &movieStatusModel,nil
}

func (m MovieStatusPersistence)Save(movieStatusModel *model.MovieStatusModel)error{
	db := ConnectGorm()
	defer db.Close()

	result := db.Table("movies").Where("id = ?",movieStatusModel.MovieID).Save(movieStatusModel)
	if result.Error != nil{
		return result.Error
	}

	return nil
}