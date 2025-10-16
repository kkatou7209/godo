package value

import (
	"regexp"
	"strings"
)

type Email struct {
	value string
}

func NewEmail(value string) Email {
	value = strings.TrimSpace(value)
	if value == "" {
		panic("empty string cannot be empty")
	}
	if matched, _ := regexp.MatchString("^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$", value); !matched {
		panic("invalid email")
	}
	return Email{value}
}

// Get value of email.
func (e Email) Value() string {
	return e.value
}