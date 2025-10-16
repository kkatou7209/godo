package dto

import "github.com/kkatou7209/godo/app/domain/value"

type TodoItemDto struct {
	Id 			value.TodoItemId
	Title 		value.TodoItemTitle
	Description value.TodoItemDescription
	IsDone 		bool
}

type CreateTodoCommand struct {
	UserId 		value.UserId
	Title 		value.TodoItemTitle
	Description value.TodoItemDescription
}