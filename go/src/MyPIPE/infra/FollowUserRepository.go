package infra

import (
	"MyPIPE/domain/model"
)

type FollowUserPersistence struct{}

func NewFollowUserPersistence()*FollowUserPersistence{
	return &FollowUserPersistence{}
}

func (f FollowUserPersistence)FindByUserIdAndFollowId(userId model.UserID,followId model.UserID)*model.FollowUser{
	db := ConnectGorm()
	defer db.Close()
	var followUser model.FollowUser
	findResult := db.Where("user_id = ? and follow_id = ?",userId,followId).Find(&followUser)
	if findResult.RowsAffected == 0{
		return nil
	}
	return &followUser
}

func (f FollowUserPersistence)Save(user *model.FollowUser)error{
	db := ConnectGorm()
	defer db.Close()
	var followUser model.FollowUser
	result := db.Where("user_id = ? and follow_id = ?",user.UserID,user.FollowID).Find(&followUser)

	if result.RowsAffected == 0{
		createResult := db.Create(&user)
		if createResult.Error != nil{
			return createResult.Error
		}
	}
	return nil
}