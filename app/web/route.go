package web

import (
	"github.com/kkatou7209/godo/app"
	"github.com/kkatou7209/godo/web/handler"
	"github.com/labstack/echo/v4"
)

func MapRoutes(e *echo.Echo, app *app.Application) {

	e.POST("/auth/signup", handler.SignUp(app))

	e.POST("/auth/login", handler.Login(app));

	e.GET("/user/:userId", handler.GetUserById(app))

	e.PUT("/user/:userId", handler.UpdateUser(app))

	e.PATCH("/user/:userId/password", handler.ChangeUserPassword(app))
	
	e.GET("/user/:userId/todo-items", handler.ListTodoItems(app))

	e.POST("/user/:userId/todo-item", handler.AddTodoItem(app))

	e.PUT("/user/:userId/todo-item/:todoItemId", handler.UpdateTodoItem(app))

	e.PATCH("/user/:userId/todo-item/:todoItemId/complete", handler.CompleteTodoItem(app))

	e.PATCH("/user/:userId/todo-item/:todoItemId/uncomplete", handler.UncompleteTodoItem(app))

	e.DELETE("/user/:userId/todo-item/:todoItemId", handler.DeleteTodoItem(app))
}