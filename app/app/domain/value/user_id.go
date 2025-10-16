package value

//lint:file-ignore ST1001 test
import (
	"strings"
)

// ID of user
type UserId struct {
	value string
}

func NewUserId(value string) UserId {
	value = strings.TrimSpace(value)
	if value == "" {
		panic("empty cannot be set")
	}
	return UserId{value}
}

// Get value of user ID.
func (u UserId) Value() string {
	return u.value
}