package persistence

import (
	"github.com/kkatou7209/godo/app/domain/entity"
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/kkatou7209/godo/app/port/out/dto"
)

type GetUserPersistence interface { 
	// Get user by ID.
	GetById(userId value.UserId) (*entity.User, error)
	// Get user by email.
	GetByEmail(email value.Email) (*entity.User, error)
}

type UpdateUserPersistence interface {
	// Update user.
	Update(user *entity.User) error
}

type CreateUserPersistence interface {
	// Create new user.
	Create(user *dto.CreateUserCommand) error
}