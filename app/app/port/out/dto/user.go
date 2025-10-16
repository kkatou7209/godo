package dto

import "github.com/kkatou7209/godo/app/domain/value"

type CreateUserCommand struct {
	UserName value.UserName
	Email value.Email
	Password value.Password
}