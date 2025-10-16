package value

import (
	"strings"
)

// ID of ToDo item.
type TodoItemId struct {
	value string
}

func NewTodoItemId(value string) TodoItemId {
	value = strings.TrimSpace(value)
	if value == "" {
		panic("empty cannot be set")
	}
	return TodoItemId{value}
}

// Get value of todo item ID.
func (t TodoItemId) Value() string {
	return t.value
}