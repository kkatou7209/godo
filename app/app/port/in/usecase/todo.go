package usecase

import "github.com/kkatou7209/godo/app/port/in/dto"

type AddTodoUsecase interface {
	// Add new todo item.
	Add(todo *dto.AddTodoCommand) error
}

type GetTodoUsecase interface {
	// Get todo item.
	Get(todoId string) (*dto.TodoItemDto, error)
}

type ListTodoUsecase interface {
	// List todo items.
	List(userId string) ([]*dto.TodoItemDto, error)
}

type UpdateTodoUsecase interface {
	// Update todo.
	Update(*dto.UpdateTodoCommand) error
}

type CompleteTodoUsecase interface {
	// Complete todo item.
	Complete(userId string, todoId string) error
}

type UncompleteTodoUsecase interface {
	// Uncomplete todo item.
	Uncomplete(userId string, todoId string) error
}

type DeleteTodoUsecase interface {
	// Delete todo item.
	Delete(userId string, todoId string) error
}