package handler

import (
	"net/http"
	"strings"

	"github.com/kkatou7209/godo/app"
	"github.com/kkatou7209/godo/app/port/in/dto"
	"github.com/kkatou7209/godo/app/validation"
	"github.com/kkatou7209/godo/web/data"
	"github.com/labstack/echo/v4"
)

type TodoData struct {
	Id          string `json:"id"`
	Title 		string `json:"title"`
	Description string `json:"description"`
	IsDone 		bool   `json:"isDone"`
}

func ListTodoItems(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

		userId := c.Param("userId")

		if strings.TrimSpace(userId) == "" {
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage("user ID not provided").
					WithErrors("userId", "userId cannot be empty"),
			)
		}

		todos, err := app.ListTodoUsecase().List(userId)

		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage("unexpected error").
					WithErrors("couse", err.Error()),
			)
		}

		if len(todos) == 0 {
			return c.JSON(
				http.StatusOK,
				data.NewPayload[any](data.StatusSuccess, make([]TodoData, 0)).
					WithMessage("no todos found"),
			)
		}

		todoJsons := make([]TodoData, len(todos))

		for i, todo := range todos {
			todoJsons[i] = TodoData{
				Id: todo.Id,
				Title: todo.Title,
				Description: todo.Description,
				IsDone: todo.IsDone,
			}
		}

		return c.JSON(
			http.StatusOK,
			data.NewPayload(data.StatusSuccess, todoJsons).
				WithMessage("get todo items successfully"),
		)
	}
}

func AddTodoItem(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

		userId := c.Param("userId")

		if strings.TrimSpace(userId) == "" {
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithErrors("userId", "empty cannot be set"),
			)
		}

		todo := new(struct {
			Title string       `json:"title"       validate:"required"`
			Description string `json:"description"`
		})

		if err := c.Bind(&todo); err != nil {
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage(err.Error()),
			)
		}

		todoDto := &dto.AddTodoCommand{
			UserId: 	 userId,
			Title: 		 todo.Title,
			Description: todo.Description,
		}

		if err := app.AddTodoUsecase().Add(todoDto); err != nil {

			if e, ok := err.(*validation.ValidationError); ok {
				return c.JSON(
					http.StatusBadRequest,
					data.NewPayload[any](data.StatusFail, nil).
						WithMessage(e.Error()),
				)
			}

			return c.JSON(
				http.StatusInternalServerError,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage("unexpected error").
					WithErrors("couse", err.Error()),
			)
		}

		return c.JSON(
			http.StatusCreated,
			data.NewPayload[any](data.StatusSuccess, nil).
				WithMessage("todo item created"),
		)
	}
}

func UpdateTodoItem(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

		userId := c.Param("userId")

		if strings.TrimSpace(userId) == "" {
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithErrors("userId", "empty cannot be set"),
			)
		}

		todoItemId := c.Param("todoItemId")

		if strings.TrimSpace(todoItemId) == "" {
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithErrors("todoItemId", "empty cannot be set"),
			)
		}

		todo := new(struct {
			Title 		string `json:"title"       validate:"required"`
			Description string `json:"description" validate:"required"`
		})

		if err := c.Bind(&todo); err != nil {
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage(err.Error()),
			)
		}

		if strings.TrimSpace(todo.Title) == "" {
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage("title is empty").
					WithErrors("title", "empty cannot be set"),
			)
		}

		todoDto := &dto.UpdateTodoCommand{
			Id: todoItemId,
			Title: todo.Title,
			Description: todo.Description,
			UserId: userId,
		}

		if err := app.UpdateTodoUsecase().Update(todoDto); err != nil {

			if e, ok := err.(*validation.ValidationError); ok {
				return c.JSON(
					http.StatusBadRequest,
					data.NewPayload[any](data.StatusFail, nil).
						WithMessage(e.Error()),
				)
			}

			return c.JSON(
				http.StatusInternalServerError,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage("unexpected error").
					WithErrors("couse", err.Error()),
			)
		}

		return c.JSON(
			http.StatusOK,
			data.NewPayload[any](data.StatusSuccess, nil).
				WithMessage("todo item udated"),
		)
	}
}

func CompleteTodoItem(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

		userId := c.Param("userId")

		if strings.TrimSpace(userId) == "" {
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithErrors("userId", "empty cannot be set"),
			)
		}

		todoItemId := c.Param("todoItemId")

		if strings.TrimSpace(todoItemId) == "" {
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithErrors("todoItemId", "empty cannot be set"),
			)
		}

		if err := app.CompleteTodoUsecase().Complete(userId, todoItemId); err != nil {
		
			if e, ok := err.(*validation.ValidationError); ok {
				return c.JSON(
					http.StatusBadRequest,
					data.NewPayload[any](data.StatusFail, nil).
						WithMessage(e.Error()),
				)
			}

			return c.JSON(
				http.StatusInternalServerError,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage("unexpected error").
					WithErrors("couse", err.Error()),
			)
		}

		return c.JSON(
			http.StatusOK,
			data.NewPayload[any](data.StatusSuccess, nil).
				WithMessage("todo item completed"),
		)
	}
}

func UncompleteTodoItem(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

		userId := c.Param("userId")

		if strings.TrimSpace(userId) == "" {
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithErrors("userId", "empty cannot be set"),
			)
		}

		todoItemId := c.Param("todoItemId")

		if strings.TrimSpace(todoItemId) == "" {
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithErrors("todoItemId", "empty cannot be set"),
			)
		}

		if err := app.UncompleteTodoUsecase().Uncomplete(userId, todoItemId); err != nil {

			if e, ok := err.(*validation.ValidationError); ok {
				return c.JSON(
					http.StatusBadRequest,
					data.NewPayload[any](data.StatusFail, nil).
						WithMessage(e.Error()),
				)
			}

			return c.JSON(
				http.StatusInternalServerError,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage("unexpected error").
					WithErrors("couse", err.Error()),
			)
		}

		return c.JSON(
			http.StatusOK,
			data.NewPayload[any](data.StatusSuccess, nil).
				WithMessage("todo item uncompleted"),
		)
	}
}

func DeleteTodoItem(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

		userId := c.Param("userId")

		if strings.TrimSpace(userId) == "" {
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithErrors("userId", "empty cannot be set"),
			)
		}

		todoItemId := c.Param("todoItemId")

		if strings.TrimSpace(todoItemId) == "" {
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithErrors("todoItemId", "empty cannot be set"),
			)
		}

		if err := app.DeleteTodoUsecase().Delete(userId, todoItemId); err != nil {

			if e, ok := err.(*validation.ValidationError); ok {
				return c.JSON(
					http.StatusBadRequest,
					data.NewPayload[any](data.StatusFail, nil).
						WithMessage(e.Error()),
				)
			}

			return c.JSON(
				http.StatusInternalServerError,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage("unexpected error").
					WithErrors("couse", err.Error()),
			)
		}

		return c.JSON(
			http.StatusOK,
			data.NewPayload[any](data.StatusSuccess, nil).
				WithMessage("todo item deleted"),
		)
	}
}