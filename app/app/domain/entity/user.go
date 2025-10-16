package entity

import "github.com/kkatou7209/godo/app/domain/value"

// Application User
type User struct {
	// AUser ID.
	id value.UserId
	// Name of user
	userName value.UserName
	// Email of usee.
	email value.Email
	// Password of user.
	password value.Password
}

// Create new user.
func NewUser(id value.UserId, userName value.UserName, email value.Email, password value.Password) *User {
	return &User{ id, userName, email, password }
}

// Get user ID.
func (u *User) Id() value.UserId {
	return u.id
}

// Get user name.
func (u *User) UserName() value.UserName {
	return u.userName
}

// Get email.
func (u *User) Email() value.Email {
	return u.email
}

// Get password.
func (u *User) Password() value.Password {
	return u.password
}

// Rename user.
func (u *User) Rename(name string) {
	userName := value.NewUserName(name)
	u.userName = userName
}

// Change email.
func (u *User) ChangeEmail(email string) {
	newEmail := value.NewEmail(email)
	u.email = newEmail	
}

func (u *User) ChangePassword(password string) {
	newPassword := value.NewPassword(password)
	u.password = newPassword
}

// Check if other is same user.
func (u *User) Is(other *User) bool {
	return u.id == other.id && u.email == other.email
}