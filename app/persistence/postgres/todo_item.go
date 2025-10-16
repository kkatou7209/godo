package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/kkatou7209/godo/app/domain/entity"
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/kkatou7209/godo/app/port/out/dto"
)

type TodoItemRepository struct {
	connectionString string
}

func NewTodoItemRepository(connectionString string) *TodoItemRepository {
	return &TodoItemRepository{connectionString}
}

func (r *TodoItemRepository) Create(todo *dto.CreateTodoCommand) error {

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, r.connectionString)

	if err != nil {
		return err
	}

	defer conn.Close(ctx)

	tran, err := conn.Begin(ctx)

	if err != nil {
		return err
	}

	defer func ()  {
		if err != nil {
			_ = tran.Rollback(ctx)
			return
		}
		err = tran.Commit(ctx)
	}()

	_, err = tran.Exec(ctx, `
		INSERT INTO todo_items (
			id, title, description, is_done, user_id
		) 
		VALUES ($1, $2, $3, $4)`,
		uuid.NewString(),
		todo.Title.Value(),
		todo.Description.Value(),
		false,
		todo.UserId.Value(),
	)
	
	return err
}

func (r *TodoItemRepository) Get(todoId value.TodoItemId) (*entity.TodoItem, error) {

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, r.connectionString)

	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	rows, err := conn.Query(ctx, `
		SELECT id, title, description, is_done, user_id
		FROM todo_items
		WHERE id = $1
	`, todoId.Value())

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		id string
		title string
		description string
		isDone bool
		userId string
	)

	if rows.Next() {
		rows.Scan(&id, &title, &description, &isDone, &userId)
	} else {
		return nil, nil
	}

	return entity.NewTodoItem(
		value.NewTodoItemId(id),
		value.NewTodoItemTitle(title),
		value.NewTodoItemDescription(description),
		isDone,
		value.NewUserId(userId),
	), nil
}

func (r *TodoItemRepository) List(userId value.UserId) ([]*entity.TodoItem, error) {

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, r.connectionString)

	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	rows, err := conn.Query(ctx, `
		SELECT id, title, description, is_done
		FROM todo_items
		WHERE user_id = $1
	`,
	userId.Value())

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		id string
		title string
		description string
		isDone bool
	)

	todos := make([]*entity.TodoItem, 0)

	if rows.Next() {

		err = rows.Scan(&id, &title, &description, &isDone, &userId)

		if err != nil {
			return nil, err
		}

		todos = append(todos, entity.NewTodoItem(
			value.NewTodoItemId(id),
			value.NewTodoItemTitle(title),
			value.NewTodoItemDescription(description),
			isDone,
			userId,
		))
	} else {
		return nil, nil
	}

	return todos, nil
}

func (r *TodoItemRepository) Update(todo *entity.TodoItem) error {

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, r.connectionString)

	if err != nil {
		return err
	}

	defer conn.Close(ctx)

	tran, err := conn.Begin(ctx)

	if err != nil {
		return err
	}

	defer func ()  {
		if err != nil {
			_ = tran.Rollback(ctx)
			return
		}
		err = tran.Commit(ctx)
	}()

	_, err = tran.Exec(ctx, `
		UPDATE todo_items
		SET title = $2, description = $3, is_done = $3
		WHERE id = $1`,
		todo.Id().Value(),
		todo.Title().Value(),
		todo.Description().Value(),
		todo.IsDone(),
	)
	
	return err
}

func (r *TodoItemRepository) Delete(todoId value.TodoItemId) error {

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, r.connectionString)

	if err != nil {
		return err
	}

	defer conn.Close(ctx)

	tran, err := conn.Begin(ctx)

	if err != nil {
		return err
	}

	defer func ()  {
		if err != nil {
			_ = tran.Rollback(ctx)
			return
		}
		err = tran.Commit(ctx)
	}()

	_, err = tran.Exec(ctx, `
		DELETE FROM todo_items
		WHERE id = $1
	`, todoId.Value())
	
	return err
}

func (r *TodoItemRepository) NextTodoItemId() value.TodoItemId {
	return value.NewTodoItemId(uuid.NewString())
}