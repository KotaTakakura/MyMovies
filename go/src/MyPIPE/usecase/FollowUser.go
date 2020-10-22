package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
)

type FollowUser struct {
	UserRepository       repository.UserRepository
	FollowUserRepository repository.FollowUserRepository
}

func NewFollowUser(u repository.UserRepository, f repository.FollowUserRepository) *FollowUser {
	return &FollowUser{
		UserRepository:       u,
		FollowUserRepository: f,
	}
}

func (f FollowUser) Follow(followDTO FollowDTO) error {
	followUser := f.FollowUserRepository.FindByUserIdAndFollowId(followDTO.UserID, followDTO.FollowID)
	if followUser != nil {
		return errors.New("Already Followed.")
	}

	existUser, _ := f.UserRepository.FindById(followDTO.UserID)
	if existUser == nil {
		return errors.New("No Such User.")
	}

	existFollowUser, _ := f.UserRepository.FindById(followDTO.FollowID)
	if existFollowUser == nil {
		return errors.New("No Such User To Follow.")
	}

	newFollowUser := model.NewFollowUser(followDTO.UserID, followDTO.FollowID)
	saveErr := f.FollowUserRepository.Save(newFollowUser)
	if saveErr != nil {
		return saveErr
	}
	return nil
}

type FollowDTO struct {
	UserID   model.UserID
	FollowID model.UserID
}

func NewFollowDTO(userId model.UserID, followId model.UserID) FollowDTO {
	return FollowDTO{
		UserID:   userId,
		FollowID: followId,
	}
}
