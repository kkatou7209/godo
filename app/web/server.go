package web

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/kkatou7209/godo/app"
	"github.com/kkatou7209/godo/app/port/in/dto"
	"github.com/kkatou7209/godo/app/validation"
	"github.com/kkatou7209/godo/password"
	"github.com/kkatou7209/godo/persistence/postgres"
	"github.com/labstack/echo/v4"
)

func Run() {

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

	e := echo.New()

	e.POST("/auth/signup", func(c echo.Context) error {

		user := new(struct {
			Username string `json:"username"  validate:"required"`
			Email    string `json:"email"     validate:"required,email"`
			Password string `json:"password"  validate:"required"`
		})

		if err := c.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		userDto := &dto.AddUserCommand{
			UserName: user.Username,
			Email: user.Email,
			Password: user.Password,
		}

		if err := app.AddUserUsecase().Add(userDto); err != nil {
			
			if e, ok := err.(*validation.ValidationError); ok {
				return echo.NewHTTPError(http.StatusBadRequest, e.Error())
			}

			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusCreated)
	})

	e.POST("/auth/login", func(c echo.Context) error {

		cred := new(struct {
			Email    string `json:"email"    validate:"required,email"`
			Password string `json:"password" validate:"required"`
		})

		if err := c.Bind(&cred); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		credDto := &dto.LoginCommand{
			Email: cred.Email,
			Password: cred.Password,
		}

		_, err := app.LoginUsecase().Login(credDto)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		c.SetCookie(&http.Cookie{
			SameSite: http.SameSiteLaxMode,
			HttpOnly: true,
			Name: "token",
			Value: "test-token",
		})

		return c.NoContent(http.StatusOK)
	});

	e.PUT("/user/:userId", func(c echo.Context) error {

		userId := c.Param("userId")

		user, err := app.GetUserUsecase().Get(userId)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if user == nil {
			return c.NoContent(http.StatusNotFound)
		}

		userJson := struct {
			Id       string `json:"id"`
			Username string `json:"username"`
			Email    string `json:"emal"`
		} {
			Id:       user.Id,
			Username: user.UserName,
			Email:    user.Email,
		}

		return c.JSON(http.StatusOK, struct {
			User any
		} {
			User: userJson,
		})
	})

	e.PATCH("/user/:userId/password", func(c echo.Context) error {

		userId := c.Param("userId")

		passwords := new(struct {
			NewPassword string `json:"newPassword" validate:"required"`
			OldPassword string `json:"oldPassword" validate:"required"`
		})

		if err := c.Bind(&passwords); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := app.ChangeUserPasswordUsecase().ChangePassword(userId, passwords.NewPassword, passwords.OldPassword); err != nil {

			if e, ok := err.(*validation.ValidationError); ok {
				return echo.NewHTTPError(http.StatusBadRequest, e.Error())
			}

			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusOK)
	})
	
	e.GET("/user/:userId/todo-items", func(c echo.Context) error {

		userId := c.Param("userId")

		todos, err := app.ListTodoUsecase().List(userId)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		type TodoJson struct {
			Id          string `json:"id"`
			Title 		string `json:"title"`
			Description string `json:"description"`
			IsDone 		bool   `json:"isDone"`
		}

		type TodosJson struct {
			Todos []*TodoJson `json:"todos"`
		}

		if len(todos) == 0 {
			return c.JSON(http.StatusOK, new(TodosJson))
		}

		todoJsons := make([]*TodoJson, len(todos))

		for i, todo := range todos {
			todoJsons[i].Id = todo.Id
			todoJsons[i].Title = todo.Title
			todoJsons[i].Description = todo.Description
			todoJsons[i].IsDone = todo.IsDone
		}

		return c.JSON(http.StatusOK, &TodosJson{
			Todos: todoJsons,
		})
	})

	e.POST("/user/:userId/todo-item", func(c echo.Context) error {

		userId := c.Param("userId")

		if strings.TrimSpace(userId) == "" {
			return echo.NewHTTPError(http.StatusBadRequest, errors.New("userId not provided"))
		}

		todo := new(struct {
			Title string       `json:"title"       validate:"required"`
			Description string `json:"description"`
		})

		if err := c.Bind(&todo); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		todoDto := &dto.AddTodoCommand{
			UserId: 	 userId,
			Title: 		 todo.Title,
			Description: todo.Description,
		}

		if err := app.AddTodoUsecase().Add(todoDto); err != nil {

			if e, ok := err.(*validation.ValidationError); ok {
				return echo.NewHTTPError(http.StatusBadRequest, e.Error())
			}

			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusCreated)
	})

	e.PUT("/user/:userId/todo-item/:todoItemId", func(c echo.Context) error {

		userId := c.Param("userId")

		if strings.TrimSpace(userId) == "" {
			return echo.NewHTTPError(http.StatusBadRequest, errors.New("userId not provided"))
		}

		todoItemId := c.Param("todoItemId")

		if strings.TrimSpace(todoItemId) == "" {
			return echo.NewHTTPError(http.StatusBadRequest, errors.New("todoItemId not provided"))
		}

		todo := new(struct {
			Title 		string `json:"title"       validate:"required"`
			Description string `json:"description" validate:"required"`
			IsDone 		bool   `json:"isDone"      validate:"required"`
		})

		if err := c.Bind(&todo); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if strings.TrimSpace(todo.Title) == "" {
			return echo.NewHTTPError(http.StatusBadRequest, errors.New("title cannot be empty"))
		}

		todoDto := &dto.TodoItemDto{
			Id: todoItemId,
			Title: todo.Title,
			Description: todo.Description,
			IsDone: todo.IsDone,
			UserId: userId,
		}

		if err := app.UpdateTodoUsecase().Update(todoDto); err != nil {

			if e, ok := err.(*validation.ValidationError); ok {
				return echo.NewHTTPError(http.StatusBadRequest, e.Error())
			}

			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusAccepted)
	})

	e.PATCH("/user/:userId/todo-item/:todoItemId/complete", func(c echo.Context) error {

		userId := c.Param("userId")

		if strings.TrimSpace(userId) == "" {
			return echo.NewHTTPError(http.StatusBadRequest, errors.New("userId not provided"))
		}

		todoItemId := c.Param("todoItemId")

		if strings.TrimSpace(todoItemId) == "" {
			return echo.NewHTTPError(http.StatusBadRequest, errors.New("todoItemId not provided"))
		}

		if err := app.CompleteTodoUsecase().Complete(userId, todoItemId); err != nil {
		
			if e, ok := err.(*validation.ValidationError); ok {
				return echo.NewHTTPError(http.StatusBadRequest, e.Error())
			}

			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusAccepted)
	})

	e.PATCH("/user/:userId/todo-item/:todoItemId/uncomplete", func(c echo.Context) error {

		userId := c.Param("userId")

		if strings.TrimSpace(userId) == "" {
			return echo.NewHTTPError(http.StatusBadRequest, errors.New("userId not provided"))
		}

		todoItemId := c.Param("todoItemId")

		if strings.TrimSpace(todoItemId) == "" {
			return echo.NewHTTPError(http.StatusBadRequest, errors.New("todoItemId not provided"))
		}

		if err := app.UncompleteTodoUsecase().Uncomplete(userId, todoItemId); err != nil {

			if e, ok := err.(*validation.ValidationError); ok {
				return echo.NewHTTPError(http.StatusBadRequest, e.Error())
			}

			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusAccepted)
	})

	e.DELETE("/user/:userId/todo-item/:todoItemId", func(c echo.Context) error {

		userId := c.Param("userId")

		if strings.TrimSpace(userId) == "" {
			return echo.NewHTTPError(http.StatusBadRequest, errors.New("userId not provided"))
		}

		todoItemId := c.Param("todoItemId")

		if strings.TrimSpace(todoItemId) == "" {
			return echo.NewHTTPError(http.StatusBadRequest, errors.New("todoItemId not provided"))
		}

		if err := app.DeleteTodoUsecase().Delete(userId, todoItemId); err != nil {

			if e, ok := err.(*validation.ValidationError); ok {
				return echo.NewHTTPError(http.StatusBadRequest, e.Error())
			}
			
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusAccepted)
	})

	if err := e.Start(":1323"); err != nil {
		e.Logger.Fatal(err)
	}
}