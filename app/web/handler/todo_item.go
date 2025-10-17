package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/kkatou7209/godo/app"
	"github.com/kkatou7209/godo/app/port/in/dto"
	"github.com/kkatou7209/godo/app/validation"
	"github.com/labstack/echo/v4"
)

func ListTodoItems(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

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
	}
}

func AddTodoItem(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

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
	}
}

func UpdateTodoItem(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

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
	}
}

func CompleteTodoItem(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

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
	}
}

func UncompleteTodoItem(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

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
	}
}

func DeleteTodoItem(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

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
	}
}