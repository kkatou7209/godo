package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kkatou7209/godo/app"
	"github.com/kkatou7209/godo/password"
	"github.com/kkatou7209/godo/persistence/postgres"
	"github.com/kkatou7209/godo/web"
	"github.com/labstack/echo/v4"
	"github.com/urfave/cli/v2"
)

func main() {

	cli := &cli.App{
		Name: "godo",
		Usage: "Serve GoDo API",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "host",
				Aliases: []string{"H"},
				Value: "localhost",
				Usage: "Specify the host name.",
			},
			&cli.StringFlag{
				Name: "port",
				Aliases: []string{"p"},
				Value: "8000",
				Usage: "Specify the port number.",
			},
			&cli.StringFlag{
				Name: "connection",
				Aliases: []string{"c"},
				Value: os.Getenv("GODO_DATABASE_URL"),
				Usage: "Specify the database connection string.",
			},
		},
		Action: func(c *cli.Context) error {

			host := c.String("host")
			port := c.String("port")
			conn := c.String("connection")

			app := app.New()

			todoRepository := postgres.NewTodoItemRepository(conn)

			userRepository := postgres.NewUserRepository(conn)

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
			e.HideBanner = true

			web.MapRoutes(e, app)

			return e.Start(fmt.Sprintf("%s:%s", host, port))
		},
	}

	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}