package infra

import (
	"MyPIPE/domain/model"
	"github.com/jinzhu/gorm"
)

type FollowUserPersistence struct{
	DatabaseAccessor *gorm.DB
}

func NewFollowUserPersistence()*FollowUserPersistence{
	return &FollowUserPersistence{
		DatabaseAccessor: ConnectGorm(),
	}
}

func (f FollowUserPersistence)FindByUserIdAndFollowId(userId model.UserID,followId model.UserID)*model.FollowUser{
	var followUser model.FollowUser
	findResult := f.DatabaseAccessor.Where("user_id = ? and follow_id = ?",userId,followId).Find(&followUser)
	if findResult.RowsAffected == 0{
		return nil
	}
	return &followUser
}

func (f FollowUserPersistence)Save(user *model.FollowUser)error{
	var followUser model.FollowUser
	result := f.DatabaseAccessor.Where("user_id = ? and follow_id = ?",user.UserID,user.FollowID).Find(&followUser)

	if result.RowsAffected == 0{
		createResult := f.DatabaseAccessor.Create(&user)
		if createResult.Error != nil{
			return createResult.Error
		}
	}
	return nil
}