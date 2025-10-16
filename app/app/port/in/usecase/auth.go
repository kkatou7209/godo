package usecase

import (
	"github.com/kkatou7209/godo/app/port/in/dto"
)

type LoginUsecase interface {
	// Login user.
	Login(user *dto.LoginCommand) (*dto.UserDto, error)
}