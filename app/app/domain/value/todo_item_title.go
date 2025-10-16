package value

import "strings"

// Title of ToDo item.
type TodoItemTitle struct {
	value string
}

// Create new ToDo item title.
func NewTodoItemTitle(value string) TodoItemTitle {
	value = strings.TrimSpace(value)
	if value == "" {
		panic("empty cannot be set")
	}
	return TodoItemTitle{value}
}

// Get value of todo item title.
func (t TodoItemTitle) Value() string {
	return t.value
}