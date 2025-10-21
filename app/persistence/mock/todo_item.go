package mock

import (
	"sync"

	"github.com/google/uuid"
	"github.com/kkatou7209/godo/app/domain/entity"
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/kkatou7209/godo/app/port/out/dto"
)

type MockTodoItemRepository struct {
	todos map[value.TodoItemId]*entity.TodoItem
	mu sync.Mutex
}

func NewMockTodoItemRepository() *MockTodoItemRepository {
	return &MockTodoItemRepository{
		todos: make(map[value.TodoItemId]*entity.TodoItem),
		mu: sync.Mutex{},
	}
}

func (r *MockTodoItemRepository) Create(todo *dto.CreateTodoCommand) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	t := entity.NewTodoItem(
		value.NewTodoItemId(uuid.NewString()),
		todo.Title,
		todo.Description,
		false,
		todo.UserId,
	)

	r.todos[t.Id()] = t

	return nil
}

func (r *MockTodoItemRepository) Get(todoId value.TodoItemId) (*entity.TodoItem, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	return r.todos[todoId], nil
}

func (r *MockTodoItemRepository) List(userId value.UserId) ([]*entity.TodoItem, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	ts := make([]*entity.TodoItem, 0)

	for _, t := range r.todos {
		ts = append(ts, t)
	}

	return ts, nil
}

func (r *MockTodoItemRepository) Update(todo *entity.TodoItem) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.todos[todo.Id()] = todo

	return nil
}

func (r *MockTodoItemRepository) Delete(todoId value.TodoItemId) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.todos, todoId)

	return nil
}