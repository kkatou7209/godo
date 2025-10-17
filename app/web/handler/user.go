package handler

import (
	"net/http"

	"github.com/kkatou7209/godo/app"
	"github.com/kkatou7209/godo/app/validation"
	"github.com/labstack/echo/v4"
)

func UpdateUser(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

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
	}
}

func ChangeUserPassword(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

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
	}
}