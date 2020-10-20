package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"mime/multipart"
)

type ChangeUserProfileImage struct{
	UserRepository	repository.UserRepository
	UserProfileImageRepository repository.UserProfileImageRepository
}

func NewChangeUserProfileImage(userRepo repository.UserRepository,userProfileImageRepo repository.UserProfileImageRepository)*ChangeUserProfileImage{
	return &ChangeUserProfileImage{
		UserRepository: userRepo,
		UserProfileImageRepository: userProfileImageRepo,
	}
}

func (c ChangeUserProfileImage)ChangeUserProfileImage(changeUserProfileImageDTO *ChangeUserProfileImageDTO)error{
	user,findUserErr := c.UserRepository.FindById(changeUserProfileImageDTO.UserID)
	if findUserErr != nil{
		return findUserErr
	}

	setProfileImageErr := user.SetProfileImage(changeUserProfileImageDTO.ProfileImageHeader)
	if setProfileImageErr != nil {
		return setProfileImageErr
	}

	uploadErr := c.UserProfileImageRepository.Upload(changeUserProfileImageDTO.ProfileImage,user)
	if uploadErr != nil{
		return uploadErr
	}

	updateUserErr := c.UserRepository.UpdateUser(user)
	if updateUserErr != nil {
		return updateUserErr
	}

	return nil
}

type ChangeUserProfileImageDTO struct{
	UserID model.UserID
	ProfileImage multipart.File
	ProfileImageHeader *multipart.FileHeader
}

func NewChangeUserProfileImageDTO(userId model.UserID, profileImage multipart.File,profileImageHeader *multipart.FileHeader)*ChangeUserProfileImageDTO{
	return &ChangeUserProfileImageDTO{
		UserID:          userId,
		ProfileImage:       profileImage,
		ProfileImageHeader: profileImageHeader,
	}
}