package value

import (
	"strings"
)

type UserName struct {
	value string
}

func NewUserName(value string) UserName {
	value = strings.TrimSpace(value)
	if value == "" { panic("empty string cennot be set") }
	return UserName{value}
}

// Get value of user name.
func (u UserName) Value() string {
	return u.value
}