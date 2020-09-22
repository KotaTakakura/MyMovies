package infra

import (
	"MyPIPE/domain/model"
	"github.com/jinzhu/gorm"
)

type PlayListPersistence struct{
	databaseAccessor *gorm.DB
}

func NewPlayListPersistence()*PlayListPersistence{
	return &PlayListPersistence{
		databaseAccessor: ConnectGorm(),
	}
}

func (p PlayListPersistence) FindByID(playListID model.PlayListID) (*model.PlayList, error) {
	panic("implement me")
}

func (p PlayListPersistence) FindByName(playListName model.PlayListName) ([]model.PlayList, error) {
	panic("implement me")
}

func (p PlayListPersistence) FindByUserID(playListUserID model.UserID) ([]model.PlayList, error) {
	panic("implement me")
}

func (p PlayListPersistence) FindByUserIDAndName(playListUserID model.UserID, playListName model.PlayListName) ([]model.PlayList, error) {
	var playLists []model.PlayList
	result := p.databaseAccessor.Where("user_id = ? and name = ?",playListUserID,playListName).Find(&playLists)
	if result.Error != nil{
		return nil,result.Error
	}
	return playLists,nil
}

func (p *PlayListPersistence) Save(playList *model.PlayList) error {
	if playList.ID == 0{
		createResult := p.databaseAccessor.Create(&playList)
		if createResult.Error != nil{
			return createResult.Error
		}
		return nil
	}
	saveResult := p.databaseAccessor.Save(playList)
	if saveResult.Error != nil{
		return saveResult.Error
	}
	return nil
}
