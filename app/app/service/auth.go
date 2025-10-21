package service

import (
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/kkatou7209/godo/app/port/in/dto"
	inDto "github.com/kkatou7209/godo/app/port/in/dto"
	"github.com/kkatou7209/godo/app/port/out/password"
	"github.com/kkatou7209/godo/app/port/out/persistence"
	"github.com/kkatou7209/godo/app/validation"
)

// LoginUsecase implementation.
type LoginService struct {
	getUserPersistence persistence.GetUserPersistence
	passwordHasher password.PasswordHasher
}

func NewLoginService(
	getUserPersistence persistence.GetUserPersistence,
	passwordHasher password.PasswordHasher,
) *LoginService {
	return &LoginService{
		getUserPersistence,
		passwordHasher,
	}
}

func (s *LoginService) Login(credential *inDto.LoginCommand) (*dto.UserDto, error) {

	user, err := s.getUserPersistence.GetByEmail(value.NewEmail(credential.Email))

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, validation.ErrUserNotFound
	}

	if !s.passwordHasher.Verify(credential.Password, user.Password().Value()) {
		return nil, validation.ErrInvalidPassword
	}

	return &dto.UserDto{
		Id: user.Id().Value(),
		UserName: user.UserName().Value(),
		Email: user.Email().Value(),
	}, nil
}