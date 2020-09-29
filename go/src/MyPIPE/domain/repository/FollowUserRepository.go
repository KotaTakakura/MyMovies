package repository

import "MyPIPE/domain/model"

type FollowUserRepository interface {
	FindByUserIdAndFollowId(userId model.UserID,followId model.UserID)*model.FollowUser
	Save(user *model.FollowUser)error
}
