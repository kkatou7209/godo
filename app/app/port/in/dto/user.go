package dto

type UserDto struct {
	Id string
	UserName string
	Email string
}

type AddUserCommand struct {
	UserName string
	Email string
	Password string
}