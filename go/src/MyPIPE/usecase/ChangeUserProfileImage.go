package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type IChangeUserProfilieImage interface {
	ChangeUserProfileImage(changeUserProfileImageDTO *ChangeUserProfileImageDTO) error
}

type ChangeUserProfileImage struct {
	UserRepository             repository.UserRepository
	UserProfileImageRepository repository.UserProfileImageRepository
}

func NewChangeUserProfileImage(userRepo repository.UserRepository, userProfileImageRepo repository.UserProfileImageRepository) *ChangeUserProfileImage {
	return &ChangeUserProfileImage{
		UserRepository:             userRepo,
		UserProfileImageRepository: userProfileImageRepo,
	}
}

func (c ChangeUserProfileImage) ChangeUserProfileImage(changeUserProfileImageDTO *ChangeUserProfileImageDTO) error {
	user, findUserErr := c.UserRepository.FindById(changeUserProfileImageDTO.UserID)
	if findUserErr != nil {
		return findUserErr
	}

	setProfileImageErr := user.SetProfileImage(changeUserProfileImageDTO.ProfileImage)
	if setProfileImageErr != nil {
		return setProfileImageErr
	}

	uploadErr := c.UserProfileImageRepository.Upload(changeUserProfileImageDTO.ProfileImage.File, user)
	if uploadErr != nil {
		return uploadErr
	}

	updateUserErr := c.UserRepository.UpdateUser(user)
	if updateUserErr != nil {
		return updateUserErr
	}

	return nil
}

type ChangeUserProfileImageDTO struct {
	UserID             model.UserID
	ProfileImage       model.UserProfileImage
}

func NewChangeUserProfileImageDTO(userId model.UserID,profileImage *model.UserProfileImage) *ChangeUserProfileImageDTO {
	return &ChangeUserProfileImageDTO{
		UserID:             userId,
		ProfileImage:       *profileImage,
	}
}
