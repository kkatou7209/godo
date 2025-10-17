package handler

import (
	"net/http"

	"github.com/kkatou7209/godo/app"
	"github.com/kkatou7209/godo/app/port/in/dto"
	"github.com/kkatou7209/godo/app/validation"
	"github.com/labstack/echo/v4"
)

func SignUp(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

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
	}
}

func Login(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error  {
		
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
	}
}