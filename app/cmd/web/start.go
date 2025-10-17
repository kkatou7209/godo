package main

import (
	"os"

	"github.com/kkatou7209/godo/app"
	"github.com/kkatou7209/godo/password"
	"github.com/kkatou7209/godo/persistence/postgres"
	"github.com/kkatou7209/godo/web"
)

func main() {

	app := app.New()

	connextionString := os.Getenv("DATABASE_URL")

	todoRepository := postgres.NewTodoItemRepository(connextionString)

	userRepository := postgres.NewUserRepository(connextionString)

	app.
		SetCreateTodoPersistence(todoRepository).
		SetListTodoPersistence(todoRepository).
		SetGetTodoPersistence(todoRepository).
		SetUpdateTodoPersistence(todoRepository).
		SetDeleteTodoPersistence(todoRepository).
		SetCreateUserPersistence(userRepository).
		SetGetUserPersistence(userRepository).
		SetUpdateUserPersistence(userRepository).
		SetPasswordHasher(password.NewBycryptPasswordHasher())

	s := web.NewServer(app, 8080, "localhost")

	s.Run()
}