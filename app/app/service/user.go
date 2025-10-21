package service

import (
	"github.com/kkatou7209/godo/app/domain/entity"
	"github.com/kkatou7209/godo/app/domain/value"
	inDto "github.com/kkatou7209/godo/app/port/in/dto"
	outDto "github.com/kkatou7209/godo/app/port/out/dto"
	"github.com/kkatou7209/godo/app/port/out/password"
	"github.com/kkatou7209/godo/app/port/out/persistence"
	"github.com/kkatou7209/godo/app/validation"
)

// AddUserUsecase implementation.
type AddUserService struct {
	createUserPersistence persistence.CreateUserPersistence
	getUserPersistence persistence.GetUserPersistence
	passwordHasher password.PasswordHasher
}

func NewAddUserService(createUserPersistence persistence.CreateUserPersistence, getUserPersistence persistence.GetUserPersistence, passwordHasher password.PasswordHasher) *AddUserService {
	return &AddUserService{createUserPersistence, getUserPersistence, passwordHasher}
}

func (s *AddUserService) Add(user *inDto.AddUserCommand) error {
	
	sameEmailUser, err := s.getUserPersistence.GetByEmail(value.NewEmail(user.Email))
	
	if err != nil {
		return err
	}
	
	if sameEmailUser != nil {
		return validation.ErrEmailAlreadyExists
	}
	
	password, err := s.passwordHasher.Hash(user.Password)
	
	if err != nil {
		return err
	}
	
	return s.createUserPersistence.Create(&outDto.CreateUserCommand{
		UserName: value.NewUserName(user.UserName),
		Email: value.NewEmail(user.Email),
		Password: value.NewPassword(password),
	})
}

// GetUserUsecase implementation.
type GetUserService struct {
	getUserPersistence persistence.GetUserPersistence
}

func NewGetUserService(getUserPersistence persistence.GetUserPersistence) *GetUserService {
	return &GetUserService{getUserPersistence}
}

func (s *GetUserService) Get(userId string) (*inDto.UserDto, error) {
	
	user, err := s.getUserPersistence.GetById(value.NewUserId(userId))
	
	if err != nil {
		return nil, err
	}
	
	return &inDto.UserDto{
		Id: user.Id().Value(),
		UserName: user.UserName().Value(),
		Email: user.Email().Value(),
	}, nil
}

// ChangeUserInfoUsecase implementation.
type ChangeUserInfoService struct {
	updateUserPersistence persistence.UpdateUserPersistence
	getUserPersistence persistence.GetUserPersistence
}

func NewChangeUserInfoService(updateUserPersistence persistence.UpdateUserPersistence, getUserPersistence persistence.GetUserPersistence) *ChangeUserInfoService {
	return &ChangeUserInfoService{updateUserPersistence, getUserPersistence}
}

func (s *ChangeUserInfoService) ChangeInfo(user *inDto.UserDto) error {
	
	sameEmailUser, err := s.getUserPersistence.GetByEmail(value.NewEmail(user.Email))
	
	if err != nil {
		return err
	}
	
	if sameEmailUser != nil && sameEmailUser.Id() != value.NewUserId(user.Id) {
		return validation.ErrEmailAlreadyExists
	}
	
	currentUser, err := s.getUserPersistence.GetById(value.NewUserId(user.Id))
	
	if err != nil {
		return err
	}

	if currentUser == nil {
		return validation.ErrUserNotFound
	}
	
	return s.updateUserPersistence.Update(entity.NewUser(
		value.NewUserId(user.Id),
		value.NewUserName(user.UserName),
		value.NewEmail(user.Email),
		currentUser.Password(),
	))
}

// ChangeUserPasswordUsecase implementation.
type ChangeUserPasswordService struct {
	updateUserPersistence persistence.UpdateUserPersistence
	getUserPersistence persistence.GetUserPersistence
	passwordHasher password.PasswordHasher
}

func NewChangeUserPasswordService(updateUserPersistence persistence.UpdateUserPersistence, getUserPersistence persistence.GetUserPersistence, passwordHasher password.PasswordHasher) *ChangeUserPasswordService {
	return &ChangeUserPasswordService{updateUserPersistence, getUserPersistence, passwordHasher}
}

func (s *ChangeUserPasswordService) ChangePassword(userId string, password string, oldPassword string) error {
	
	currentUser, err := s.getUserPersistence.GetById(value.NewUserId(userId))
	
	if err != nil {
		return err
	}

	if currentUser == nil {
		return validation.ErrUserNotFound
	}

	if !s.passwordHasher.Verify(oldPassword, currentUser.Password().Value()) {
		return validation.ErrInvalidPassword
	}

	hashedPassword, err := s.passwordHasher.Hash(password)

	if err != nil {
		return err
	}
	
	currentUser.ChangePassword(hashedPassword)

	return s.updateUserPersistence.Update(currentUser)
}
