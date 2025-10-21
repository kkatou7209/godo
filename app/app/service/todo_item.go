package service

import (
	"errors"
	"strings"

	"github.com/kkatou7209/godo/app/domain/value"
	inDto "github.com/kkatou7209/godo/app/port/in/dto"
	"github.com/kkatou7209/godo/app/port/out/dto"
	"github.com/kkatou7209/godo/app/port/out/persistence"
	"github.com/kkatou7209/godo/app/validation"
)

// AddTodoUsecase implementation.
type AddTodoService struct {
	createTodoPersistence persistence.CreateTodoPersistence
	
}

func NewAddTodoService(createTodoPersistence persistence.CreateTodoPersistence) *AddTodoService {
	return &AddTodoService{createTodoPersistence}
}

func (s *AddTodoService) Add(todo *inDto.AddTodoCommand) error {
	return s.createTodoPersistence.Create(&dto.CreateTodoCommand{
		UserId: 	 value.NewUserId(todo.UserId),
		Title: 		 value.NewTodoItemTitle(todo.Title),
		Description: value.NewTodoItemDescription(todo.Description),
	})
}

// GetTodoUsecase implementation.
type GetTodoService struct {
	getTodoPersistence persistence.GetTodoPersistence
}

func NewGetTodoService(getTodoPersistence persistence.GetTodoPersistence) *GetTodoService {
	return &GetTodoService{getTodoPersistence}
}

func (s *GetTodoService) Get(todoId string) (*inDto.TodoItemDto, error) {
	
	todo, err := s.getTodoPersistence.Get(value.NewTodoItemId(todoId))
	
	if err != nil {
		return nil, err
	}
	
	return &inDto.TodoItemDto{
		Id: todo.Id().Value(),
		Title: todo.Title().Value(),
		Description: todo.Description().Value(),
		IsDone: todo.IsDone(),
	}, nil
}

// ListTodoUsecase implementation.
type ListTodoService struct {
	listTodoPersistence persistence.ListTodoPersistence
}

func NewListTodoService(listTodoPersistence persistence.ListTodoPersistence) *ListTodoService {
	return &ListTodoService{listTodoPersistence}
}

func (s *ListTodoService) List(userId string) ([]*inDto.TodoItemDto, error) {
	
	todos, err := s.listTodoPersistence.List(value.NewUserId(userId))

	if err != nil {
		return nil, err
	}

	dtoTodos := make([]*inDto.TodoItemDto, len(todos))
	
	for i, todo := range todos {
		dtoTodos[i] = &inDto.TodoItemDto{
			Id: todo.Id().Value(),
			Title: todo.Title().Value(),
			Description: todo.Description().Value(),
			IsDone: todo.IsDone(),
		}
	}

	return dtoTodos, nil
}

type UpdateTodoService struct {
	updateTodoPersistence persistence.UpdateTodoPersistence
	getTodoPersistence    persistence.GetTodoPersistence
}

func NewUpdateTodoService(updateTodoPersistence persistence.UpdateTodoPersistence, getTodoPersistence persistence.GetTodoPersistence) *UpdateTodoService {
	return &UpdateTodoService{updateTodoPersistence, getTodoPersistence}
}

func (s *UpdateTodoService) Update(todoDto *inDto.UpdateTodoCommand) error {

	todo, err := s.getTodoPersistence.Get(value.NewTodoItemId(todoDto.Id))

	if err != nil {
		return err
	}

	if todo == nil {
		return validation.ErrTodoNotDound
	}

	if strings.TrimSpace(todoDto.Title) == "" {
		return errors.New("title cannot be empty")
	}

	
	if todo.UserId() != value.NewUserId(todoDto.UserId) {
		return validation.ErrInvalidUser
	}

	todo.ChangeDescription(todoDto.Description)
	todo.ChangeTitle(todoDto.Title)

	return s.updateTodoPersistence.Update(todo)
}

// CompleteTodoUsecase implementation.
type CompleteTodoService struct {
	completeTodoPersistence persistence.UpdateTodoPersistence
	getTodoPersistence persistence.GetTodoPersistence
}

func (s *CompleteTodoService) Complete(userId string, todoId string) error {

	todo, err := s.getTodoPersistence.Get(value.NewTodoItemId(todoId))
	
	if err != nil {
		return err
	}

	if todo.UserId() != value.NewUserId(userId) {
		return errors.New("invalid user")
	}
	
	todo.Complete()
	
	return s.completeTodoPersistence.Update(todo)
}

func NewCompleteTodoService(completeTodoPersistence persistence.UpdateTodoPersistence, getTodoPersistence persistence.GetTodoPersistence) *CompleteTodoService {
	return &CompleteTodoService{completeTodoPersistence, getTodoPersistence}
}

// UncompleteTodoUsecase implementation.
type UncompleteTodoService struct {
	uncompleteTodoPersistence persistence.UpdateTodoPersistence
	getTodoPersistence persistence.GetTodoPersistence
}

func NewUncompleteTodoService(uncompleteTodoPersistence persistence.UpdateTodoPersistence, getTodoPersistence persistence.GetTodoPersistence) *UncompleteTodoService {
	return &UncompleteTodoService{uncompleteTodoPersistence, getTodoPersistence}
}

func (s *UncompleteTodoService) Uncomplete(userId string, todoId string) error {
	
	todo, err := s.getTodoPersistence.Get(value.NewTodoItemId(todoId))
	
	if err != nil {
		return err
	}

	if todo.UserId() != value.NewUserId(userId) {
		return errors.New("invalid user")
	}
	
	todo.Uncomplete()
	
	return s.uncompleteTodoPersistence.Update(todo)
}

// DeleteTodoUsecase implementation.
type DeleteTodoService struct {
	deleteTodoPersistence persistence.DeleteTodoPersistence
	getTodoPersistence persistence.GetTodoPersistence
}

func NewDeleteTodoService(deleteTodoPersistence persistence.DeleteTodoPersistence, getTodoPersistence persistence.GetTodoPersistence) *DeleteTodoService {
	return &DeleteTodoService{deleteTodoPersistence, getTodoPersistence}
}

func (s *DeleteTodoService) Delete(userId string, todoId string) error {

	todo, err := s.getTodoPersistence.Get(value.NewTodoItemId(todoId))
	
	if err != nil {
		return err
	}

	if todo.UserId() != value.NewUserId(userId) {
		return errors.New("invalid user")
	}

	return s.deleteTodoPersistence.Delete(todo.Id())
}