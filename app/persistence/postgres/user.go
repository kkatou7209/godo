package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/kkatou7209/godo/app/domain/entity"
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/kkatou7209/godo/app/port/out/dto"
)

type UserRepository struct {
	connectionString string
}

func NewUserRepository(connectionString string) *UserRepository {
	return &UserRepository{connectionString}
}

func (r *UserRepository) Create(user *dto.CreateUserCommand) error {

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
		INSERT INTO users (
			id, username, email, password
		)
		VALUES (
			$1, $2, $3, $4
		)`,
		uuid.NewString(),
		user.UserName.Value(),
		user.Email.Value(),
		user.Password.Value(),
	)

	return err
}

func (r *UserRepository) GetById(userId value.UserId) (*entity.User, error) {

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, r.connectionString)

	if err != nil { return nil, err }

	rows, err := conn.Query(ctx, `
		SELECT id, username, email, password
		FROM users
		WHERE id = $1
	`, userId.Value())

	if err != nil { 
		return nil, err
	}

	var (
		id string
		username string
		email string
		password string
	)

	if rows.Next() {
		rows.Scan(&id, &username, &email, &password)
		return entity.NewUser(
			value.NewUserId(id),
			value.NewUserName(username),
			value.NewEmail(email),
			value.NewPassword(password),
		), nil
	}

	return nil, nil
}

func (r *UserRepository) GetByEmail(email value.Email) (*entity.User, error) {

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, r.connectionString)

	if err != nil { return nil, err }

	defer conn.Close(ctx)

	rows, err := conn.Query(ctx, `
		SELECT id, username, password
		FROM users
		WHERE email = $1
	`, email.Value())

	if err != nil { 
		return nil, err
	}

	var (
		id string
		username string
		password string
	)

	if rows.Next() {
		rows.Scan(&id, &username, &password)
		return entity.NewUser(
			value.NewUserId(id),
			value.NewUserName(username),
			email,
			value.NewPassword(password),
		), nil
	}

	return nil, nil
}

func (r *UserRepository) Update(user *entity.User) error {

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

	_, err = tran.Exec(context.Background(), `
		UPDATE users
		SET username = $1, email = $2, password = $3
		`,
		user.UserName().Value(),
		user.Email().Value(),
		user.Password().Value(),
	)

	return err
}