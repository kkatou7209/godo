package handler

import (
	"net/http"

	"github.com/kkatou7209/godo/app"
	"github.com/kkatou7209/godo/app/port/in/dto"
	"github.com/kkatou7209/godo/app/validation"
	"github.com/kkatou7209/godo/web/data"
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
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage("invalid request").
					WithErrors("errors", err.Error()),
			)
		}

		userDto := &dto.AddUserCommand{
			UserName: user.Username,
			Email: user.Email,
			Password: user.Password,
		}
	
		if err := app.AddUserUsecase().Add(userDto); err != nil {
			
			if e, ok := err.(*validation.ValidationError); ok {
				return c.JSON(
					http.StatusBadRequest,
					data.NewPayload[any](data.StatusFail, nil).
						WithMessage("validation error").
						WithErrors("couse", e.Error()),
				)
			}
	
			return c.JSON(
				http.StatusInternalServerError,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage("unexpected error occured").
					WithErrors("couse", err.Error()),
			)
		}
	
		return c.JSON(
			http.StatusCreated,
			data.NewPayload[any](data.StatusSuccess, nil).
				WithMessage("user created successdully"),
		)
	}
}

func Login(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error  {
		
		cred := new(struct {
			Email    string `json:"email"    validate:"required,email"`
			Password string `json:"password" validate:"required"`
		})

		if err := c.Bind(&cred); err != nil {

			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage("invalid request").
					WithErrors("errors", err.Error()),
			)
		}

		credDto := &dto.LoginCommand{
			Email: cred.Email,
			Password: cred.Password,
		}

		_, err := app.LoginUsecase().Login(credDto)

		if err != nil {

			if _, ok := err.(*validation.ValidationError); ok {
				return c.JSON(
					http.StatusBadRequest,
					data.NewPayload[any](data.StatusFail, nil).
					WithMessage("invalid email or password").
						WithErrors("errors", err.Error()),
				)
			}

			return c.JSON(
				http.StatusInternalServerError,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage("unexpected error occured").
					WithErrors("couse", err.Error()),
			)
		}

		c.SetCookie(&http.Cookie{
			SameSite: http.SameSiteLaxMode,
			HttpOnly: true,
			Name: "x-api-token",
			Value: "test-token",
		})

		return c.JSON(
			http.StatusOK,
			data.NewPayload[any](data.StatusSuccess, nil).
				WithMessage("user loged in successdully"),
		)
	}
}