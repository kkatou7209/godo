package value

import "strings"

// Password of user.
type Password struct {
	value string
}

func NewPassword(value string) Password {
	value = strings.TrimSpace(value)
	if value == "" {
		panic("empty string cannot be empty")
	}
	return Password{value}
}

// Get value of password.
func (p Password) Value() string {
	return p.value
}