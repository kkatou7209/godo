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

type UserData struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GetUserById(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

		userId := c.Param("userId")

		if strings.TrimSpace(userId) == "" {
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage("invalid request").
					WithErrors("userId", "userId cannot be empty"),
			)
		}

		user, err := app.GetUserUsecase().Get(userId)

		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage("unexpected error").
					WithErrors("coause", err.Error()),
			)
		}

		if user == nil {
			return c.JSON(
				http.StatusOK,
				data.NewPayload[any](data.StatusSuccess, nil).
					WithMessage("no user found"),
			)
		}

		userJson := &UserData{
			Id:       user.Id,
			Username: user.UserName,
			Email:    user.Email,
		}

		return c.JSON(http.StatusOK, 
			data.NewPayload(data.StatusSuccess, userJson).
				WithMessage("user found"),
		)
	}
}

func UpdateUser(app *app.Application) (func(c echo.Context) error) {

	return func(c echo.Context) error {

		userId := c.Param("userId")

		if strings.TrimSpace(userId) == "" {
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithErrors("userId", "empty cannot be set"),
			)
		}

		userInfo := new(struct{
			Username string `json:"username"`
			Email    string `json:"email"`
		})

		err := c.Bind(&userInfo)

		if err != nil {
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithErrors("couse", err.Error()),
			)
		}

		err = app.ChangeUserInfoUsecase().ChangeInfo(&dto.UserDto{
			Id: userId,
			UserName: userInfo.Username,
			Email: userInfo.Email,
		})

		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				data.NewPayload[any](data.StatusFail, nil).
					WithErrors("couse", err.Error()),
			)
		}

		return c.JSON(
			http.StatusOK,
			data.NewPayload[any](data.StatusSuccess, nil).
				WithMessage("user updated successfully"),
		)
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
			return c.JSON(
				http.StatusBadRequest,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage(err.Error()),
			)
		}

		if err := app.ChangeUserPasswordUsecase().ChangePassword(userId, passwords.NewPassword, passwords.OldPassword); err != nil {

			if e, ok := err.(*validation.ValidationError); ok {
				return c.JSON(
					http.StatusBadRequest,
					data.NewPayload[any](data.StatusFail, nil).
						WithMessage("invalid password").
						WithErrors("password", e.Error()),
				)
			}

			return c.JSON(
				http.StatusInternalServerError,
				data.NewPayload[any](data.StatusFail, nil).
					WithMessage(err.Error()),
			)
		}

		return c.JSON(
			http.StatusOK,
			data.NewPayload[any](data.StatusSuccess, nil).
				WithMessage("password successfully changed"),
		)
	}
}