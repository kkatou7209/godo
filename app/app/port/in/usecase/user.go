package usecase

import "github.com/kkatou7209/godo/app/port/in/dto"

type AddUserUsecase interface {
	// Add user.
	Add(user *dto.AddUserCommand) error
}

type GetUserUsecase interface {
	// Get user.
	Get(userId string) (*dto.UserDto, error)
}

type ChangeUserInfoUsecase interface {
	// Change user info.
	ChangeInfo(user *dto.UserDto) error
}

type ChangeUserPasswordUsecase interface {
	// Change user password.
	ChangePassword(userId string, password string, oldPassword string) error
}