package handler_test

import (
	"log"
	"testing"

	ap "github.com/kkatou7209/godo/app"
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/kkatou7209/godo/app/port/in/dto"
	"github.com/kkatou7209/godo/password"
	"github.com/kkatou7209/godo/persistence/mock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	app *ap.Application
	e *echo.Echo = echo.New()
	userId value.UserId
)

var _ = BeforeSuite(func() {

	todoRepository := mock.NewMockTodoItemRepository()
	userRepository := mock.NewMockUserRepository()

	app = ap.New().
		SetCreateTodoPersistence(todoRepository).
		SetListTodoPersistence(todoRepository).
		SetGetTodoPersistence(todoRepository).
		SetUpdateTodoPersistence(todoRepository).
		SetDeleteTodoPersistence(todoRepository).
		SetCreateUserPersistence(userRepository).
		SetGetUserPersistence(userRepository).
		SetUpdateUserPersistence(userRepository).
		SetPasswordHasher(password.NewBycryptPasswordHasher())

	if err := app.AddUserUsecase().Add(&dto.AddUserCommand{
		UserName: "handler-test-user",
		Email: "handler-test@example.com",
		Password: "handler-test-pass",
	}); err != nil {
		log.Fatalln(err)
	}

	user, err := userRepository.GetByEmail(value.NewEmail("handler-test@example.com"))

	if err != nil {
		log.Fatalln(err)
	}

	userId = user.Id()
})

func TestServer(t *testing.T) {

	RegisterFailHandler(Fail)

	RunSpecs(t, "server test")
}