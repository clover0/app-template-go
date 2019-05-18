package handler

import (
	"auth465/core"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type UserCreateForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUserHandler(service core.UserService) func(c echo.Context) (err error) {
	return func(c echo.Context) (err error) {
		form := new(UserCreateForm)
		if err := c.Bind(form); err != nil {
			return err
		}

		user := new(core.User)
		user.Email = form.Email
		user.Password = form.Password
		res, err := service.CheckDuplicateEmail(user)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusInternalServerError, "error")
		}
		if !res {
			return c.JSON(http.StatusBadRequest, "bad request")
		}

		if err := service.Register(user); err != nil {
			log.Error(err)
			return c.JSON(http.StatusInternalServerError, "error")
		}

		return c.JSON(http.StatusOK, 1)
	}
}

func UpdateUserHandler(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, 1)
}
