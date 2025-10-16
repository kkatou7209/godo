package persistence

import (
	"github.com/kkatou7209/godo/app/domain/entity"
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/kkatou7209/godo/app/port/out/dto"
)

type CreateTodoPersistence interface {
	// Create new todo item.
	Create(todo *dto.CreateTodoCommand) error
}

type ListTodoPersistence interface {
	// List todo items.
	List(userId value.UserId) ([]*entity.TodoItem, error)
}

type UpdateTodoPersistence interface {
	// Update todo item.
	Update(todo *entity.TodoItem) error
}

type GetTodoPersistence interface {
	// Get todo item.
	Get(todoId value.TodoItemId) (*entity.TodoItem, error)
}

type DeleteTodoPersistence interface {
	// Delete todo item.
	Delete(todoId value.TodoItemId) error
}