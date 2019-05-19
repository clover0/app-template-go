package handler

import (
	"auth465/core"
	"golang.org/x/crypto/bcrypt"

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
		//cookie, err := c.Cookie("ESESSION")
		//if err != nil {
		//	log.Error(err)
		//	return err
		//}
		//v, err := session.Get(cookie.Value).Result()
		//if err == redis.Nil {
		//	log.Print("Key does not exists")
		//} else if err != nil {
		//	log.Error(err)
		//	return err
		//}
		//log.Info("value: ", v)

		form := new(UserCreateForm)
		if err := c.Bind(form); err != nil {
			return err
		}

		user := new(core.User)
		user.Email = form.Email
		password, err := bcrypt.GenerateFromPassword([]byte(form.Password), 10)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusInternalServerError, "error")
		}
		user.Password = string(password)
		ok, err := service.CheckDuplicateEmail(user)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusInternalServerError, "error")
		}
		if !ok {
			return c.JSON(http.StatusBadRequest, "bad request")
		}

		if err := service.Register(user); err != nil {
			log.Error(err)
			return c.JSON(http.StatusInternalServerError, "error")
		}

		return c.JSON(http.StatusCreated, 1)
	}
}

func UpdateUserHandler(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, 1)
}
