package handler

import (
	"auth465/core"
	"fmt"

	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

const cookieSessionName = "ESESSION"

type sessionCreateForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateSessionHandler(service core.SessionService) func(c echo.Context) (err error) {
	return func(c echo.Context) (err error) {

		// bind form from request
		form := new(sessionCreateForm)
		if err := c.Bind(form); err != nil {
			return err
		}

		// check database
		user, err := service.FindUserByEmail(form.Email)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusInternalServerError, "error")
		}
		// exists user?
		if user != nil {
			if service.ComparePassword(user, form.Password) == nil {
				sessionId, err := service.CreateSession(user.ID)
				if err != nil {
					log.Error(err)
					panic(err)
				}
				cookie := new(http.Cookie)
				cookie.Name = cookieSessionName
				cookie.Value = sessionId
				cookie.Expires = time.Now().Add(14 * 24 * time.Hour)
				cookie.Path = "/"
				c.SetCookie(cookie)
				return c.JSON(http.StatusCreated, 1)
			} else { // exists user but password no match
				return c.JSON(http.StatusBadRequest, "password no match")
			}
		} else { // not exists
			return c.JSON(http.StatusBadRequest, "email no match")
		}
	}
}

func DeleteSessionHandler(session *redis.Client) func(c echo.Context) (err error) {
	return func(c echo.Context) (err error) {
		cookie, err := c.Cookie(cookieSessionName)
		if err == http.ErrNoCookie {
			return c.JSON(http.StatusOK, "no cookie")
		} else if err != nil {
			return c.JSON(http.StatusInternalServerError, "can not read cookie")
		}
		session.Del(cookie.Value)
		return c.JSON(http.StatusOK, "ok")
	}
}

func ShowCurrentSessionHandler(service core.SessionService, session *redis.Client) func(c echo.Context) (err error) {
	return func(c echo.Context) (err error) {
		cookie, err := c.Cookie(cookieSessionName)
		if err == http.ErrNoCookie {
			return c.JSON(http.StatusOK, "no sign in")
		} else if err != nil {
			return c.JSON(http.StatusInternalServerError, "can not read cookie")
		}
		userId, err := session.Get(cookie.Value).Result()
		if err == redis.Nil {
			return c.JSON(http.StatusOK, "no sign in")
		} else if err != nil {
			log.Error(err)
			panic(err)
		}
		log.Info(userId)
		id, err := strconv.ParseUint(userId, 10, 32)
		if err != nil {
			log.Error(err)
			panic(err)
		}
		user, err := service.FindUserById(uint32(id))
		log.Info(id)
		log.Info(uint32(id))
		if user == nil {
			return c.JSON(http.StatusInternalServerError, "error")
		}

		return c.JSON(http.StatusOK, fmt.Sprintf("already sign in! Your email is %s", user.Email))
	}
}

