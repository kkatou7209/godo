package web

import (
	"fmt"

	"github.com/kkatou7209/godo/app"
	"github.com/kkatou7209/godo/web/handler"
	"github.com/labstack/echo/v4"
)

type Server struct {
	app 	   *app.Application
	port       int
	host       string
}

func NewServer(app *app.Application, port int, host string) *Server {
	return &Server{
		app,
		port,
		host,
	}
}

func (s *Server) Run() {

	e := echo.New()

	e.POST("/auth/signup", handler.SignUp(s.app))

	e.POST("/auth/login", handler.Login(s.app));

	e.PUT("/user/:userId", handler.UpdateUser(s.app))

	e.PATCH("/user/:userId/password", handler.ChangeUserPassword(s.app))
	
	e.GET("/user/:userId/todo-items", handler.ListTodoItems(s.app))

	e.POST("/user/:userId/todo-item", handler.AddTodoItem(s.app))

	e.PUT("/user/:userId/todo-item/:todoItemId", handler.UpdateTodoItem(s.app))

	e.PATCH("/user/:userId/todo-item/:todoItemId/complete", handler.CompleteTodoItem(s.app))

	e.PATCH("/user/:userId/todo-item/:todoItemId/uncomplete", handler.UncompleteTodoItem(s.app))

	e.DELETE("/user/:userId/todo-item/:todoItemId", handler.DeleteTodoItem(s.app))

	if err := e.Start(fmt.Sprintf("%s:%d", s.host, s.port)); err != nil {
		e.Logger.Fatal(err)
	}
}