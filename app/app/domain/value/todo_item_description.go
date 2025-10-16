package value

import "strings"

// Description of ToDo item.
type TodoItemDescription struct {
	value string
}

// Create new ToDo item description.
func NewTodoItemDescription(value string) TodoItemDescription {
	value = strings.TrimSpace(value)
	if value == "" {
		panic("empty cannot be set")
	}
	return TodoItemDescription{value}
}

// Get value of todo item description.
func (t TodoItemDescription) Value() string {
	return t.value
}