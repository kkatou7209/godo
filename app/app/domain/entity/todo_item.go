package entity

import "github.com/kkatou7209/godo/app/domain/value"

// ToDo item.
type TodoItem struct {
	// ID of todo item.
	id value.TodoItemId
	// Title of todo item.
	title value.TodoItemTitle
	// Description of todo item.
	description value.TodoItemDescription
	// Is done flag of todo item.
	isDone bool

	userId value.UserId
}

// Create new todo item.
func NewTodoItem(id value.TodoItemId, title value.TodoItemTitle, description value.TodoItemDescription, isDone bool, userId value.UserId) *TodoItem {
	return &TodoItem{id, title, description, isDone, userId}
}

// Get id of todo item.
func (t *TodoItem) Id() value.TodoItemId {
	return t.id
}

// Get title of todo item.
func (t *TodoItem) Title() value.TodoItemTitle {
	return t.title
}

// Get description of todo item.
func (t *TodoItem) Description() value.TodoItemDescription {
	return t.description
}

// Check if todo item is done.
func (t *TodoItem) IsDone() bool {
	return t.isDone
}

func (t *TodoItem) UserId() value.UserId {
	return t.userId
}

// Check if other is same todo item.
func (t *TodoItem) Is(other *TodoItem) bool {
	return t.id == other.id
}

// Complete todo item.
func (t *TodoItem) Complete() {
	t.isDone = true
}

// Uncomplete todo item.
func (t *TodoItem) Uncomplete() {
	t.isDone = false
}

// Change title of todo item.
func (t *TodoItem) ChangeTitle(title string) {
	newTitle := value.NewTodoItemTitle(title)
	t.title = newTitle
}

// Change description of todo item.
func (t *TodoItem) ChangeDescription(description string) {
	newDescription := value.NewTodoItemDescription(description)
	t.description = newDescription
}